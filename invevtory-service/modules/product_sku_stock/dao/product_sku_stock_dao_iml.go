package dao

import (
	"errors"
	"fmt"
	productSkuStockProto "github.com/qianxunke/ego-shopping/ego-common-protos/go_out/inventory/product_sku_stock"
	"github.com/qianxunke/ego-shopping/ego-plugins/db"
	"github.com/shopspring/decimal"
	"inventory-service/modules/product_sku_stock/bean"
	"inventory-service/utils/string_utl"
	"inventory-service/utils/time_util"
)

func (dao *daoIml) FindById(id int64) (product *productSkuStockProto.ProductSkuStock, err error) {
	product = &productSkuStockProto.ProductSkuStock{}
	DB := db.MasterEngine()
	err = DB.Where("id = ?", id).First(&product).Error
	return
}

func (dao *daoIml) Insert(product *productSkuStockProto.ProductSkuStock) (err error) {
	DB := db.MasterEngine()
	err = DB.Create(&product).Error
	return
}

func (dao *daoIml) SimpleQuery(productId string) (rsp *productSkuStockProto.Out_GetProductSkuStocks, err error) {
	DB := db.MasterEngine()
	rsp = &productSkuStockProto.Out_GetProductSkuStocks{}
	err = DB.Model(&productSkuStockProto.ProductSkuStock{}).Where("product_id = ?", productId).Count(&rsp.Total).Error
	if err == nil && rsp.Total > 0 {
		err = DB.Where("product_id = ?", productId).Find(&rsp.ProductSkuStockList).Error
	}
	return
}

func (dao *daoIml) Delete(ids []int64) (err error) {
	if len(ids) == 0 {
		return
	}
	DB := db.MasterEngine()
	DB.Begin()
	defer func() {
		if err != nil {
			DB.Rollback()
		}
	}()
	for i := 0; i < len(ids); i++ {
		err = DB.Where("id = ?", ids[i]).Delete(&productSkuStockProto.ProductSkuStock{}).Error
		if err != nil {
			break
		}
	}
	if err != nil {
		DB.Commit()
	}
	return
}

func (dao *daoIml) Update(id int64, reqMap map[string]interface{}) (err error) {
	DB := db.MasterEngine()
	err = DB.Model(&productSkuStockProto.ProductSkuStock{}).Where("id = ?", id).Updates(reqMap).Error
	return
}

func (dao *daoIml) UpdateStock(id int64, productId string, memberId int64, stockNum int64) (rsp *productSkuStockProto.SkuOutLog, err error) {

	tx := db.MasterEngine().Begin()
	defer func() {
		if re := recover(); re != nil {
			tx.Rollback()
			if err == nil {
				err = errors.New(fmt.Sprintf("%v", re))
			}
		} else {
			if err != nil {
				tx.Rollback()
			}
		}
	}()
	skuStock := &productSkuStockProto.ProductSkuStock{}

	//查询库存和版本判断是否够不够
	err = tx.Where("id = ? and product_id = ?", id, productId).First(&skuStock).Error
	if err != nil {
		return
	}
	if len(skuStock.ProductId) == 0 {
		err = errors.New("The item stock does not exist！")
		return
	}
	if skuStock.Stock-stockNum < 0 {
		err = errors.New("Insufficient inventory")
		return
	}
	//开始减库存
	upMap := map[string]interface{}{}
	upMap["stock"] = skuStock.Stock - stockNum
	upMap["lock_stock"] = skuStock.LockStock + 1
	r := tx.Where("id = ? and lock_stock = ? ", skuStock.Id, skuStock.LockStock).Updates(&upMap).RowsAffected
	if r == 0 {
		//版本号不对，重试，直到库存不足
		tx.Rollback()
	_:
		dao.UpdateStock(id, productId, memberId, stockNum)
		return
	} else {
		//	totalPrice:=stockNum*skuStock.PromotionPrice
		numDecimal, err := decimal.NewFromString(string_utl.Int64ToString(stockNum))
		if err != nil {
			return nil, err
		}
		unitPrice, err := decimal.NewFromString(string_utl.Float32ToFloat32String(skuStock.PromotionPrice))
		if err != nil {
			return nil, err
		}
		sValue, err := numDecimal.Mul(unitPrice).Value()
		if err != nil {
			return nil, err
		}
		value, err := string_utl.Float64StringToFloat32(sValue.(string))
		if err != nil {
			return nil, err
		}
		//生成销库记录
		sellStockLog := &bean.Sku_out_log{
			ProductId: productId, SellNum: stockNum,
			SkuPrice:    skuStock.PromotionPrice,
			SkuId:       id,
			TotalAmount: value,
			CreateTime:  time_util.GetCurrentTime(time_util.Layout_Standard),
			MemberId:    memberId,
		}
		err = tx.Create(&sellStockLog).Error
		if err != nil {
			return nil, err
		}
		rsp = &productSkuStockProto.SkuOutLog{}
		//获取销库信息
		r := tx.Where("sku_id = ? and member_id = ? and create_id = ?", sellStockLog.SkuId, sellStockLog.MemberId, sellStockLog.CreateTime).First(&rsp).RowsAffected
		if r == 0 {
			err = errors.New("get sell stock out error")
			return nil, err
		}
		return rsp, nil
	}
}
