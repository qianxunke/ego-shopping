package dao

import (
	"errors"
	"fmt"
	"inventory-service/utils/bean_util"
	"inventory-service/utils/time_util"
	productBean "inventory-service/modules/product/bean"
	attributeValueBean "inventory-service/modules/product_attribute_value/bean"
	fullReductionBean "inventory-service/modules/product_full_reduction/bean"
	ladderBean "inventory-service/modules/product_ladder/bean"
	memberBean "inventory-service/modules/product_member_price/bean"
	skuStockBean "inventory-service/modules/product_sku_stock/bean"
	productProto "github.com/qianxunke/ego-shopping/ego-common-protos/go_out/inventory/product"
	"github.com/qianxunke/ego-shopping/ego-plugins/db"
	"inventory-service/utils/uuid"
	"time"
)

func (dao *productDaoIml) FindById(id string) (product *productProto.ProductDetails, err error) {
	product = &productProto.ProductDetails{}
	product.Product = &productProto.Product{}
	DB := db.MasterEngine()
	err = DB.Where("id = ?", id).First(&product.Product).Error
	if err != nil {
		return
	}
	//获取价格梯度
	err = DB.Model(&productProto.ProductLadder{}).Where("product_id = ?", id).Find(&product.ProductLadderList).Error
	if err != nil {
		return
	}
	//获取减满
	err = DB.Model(&productProto.ProductFullReduction{}).Where("product_id = ?", id).Find(&product.ProductFullReductionList).Error
	if err != nil {
		return
	}
	//获取会员价格
	err = DB.Model(&productProto.ProductMemberPrice{}).Where("product_id = ?", id).Find(&product.ProductMemberPriceList).Error
	if err != nil {
		return
	}
	//获取sku列表
	err = DB.Model(&productProto.ProductSkuStock{}).Where("product_id = ?", id).Find(&product.ProductSkuStockList).Error
	if err != nil {
		return
	}
	//获取属性值
	err = DB.Model(&productProto.ProductAttributeValue{}).Where("product_id = ?", id).Find(&product.ProductAttributeValueList).Error
	return
}

func (dao *productDaoIml) Insert(product *productProto.In_CreateProduct) (id string, err error) {
	DB := db.MasterEngine()
	tx := DB.Begin()
	defer func() {
		if re := recover(); re != nil {
			tx.Rollback()
			if err == nil {
				err = errors.New(fmt.Sprintf("%v", re))
			}
		}
	}()
	product.ProductDetail.Product.Id = uuid.GetUuid()
	//先插入商品信息
	product.ProductDetail.Product.CreatedTime = time.Now().Format("2006-01-02 15:04:05")
	product.ProductDetail.Product.UpdateTime = product.ProductDetail.Product.CreatedTime
	standardProduct := &productBean.Product{}
	err = bean_util.ProToStandard(product.ProductDetail.Product, standardProduct)
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Create(&standardProduct).Error
	//创建会员价格
	if product.ProductDetail.ProductMemberPriceList != nil && len(product.ProductDetail.ProductMemberPriceList) > 0 {
		for i := 0; i < len(product.ProductDetail.ProductMemberPriceList); i++ {
			product.ProductDetail.ProductMemberPriceList[i].ProductId = product.ProductDetail.Product.Id
			err = tx.Create(&product.ProductDetail.ProductMemberPriceList[i]).Error
			if err != nil {
				tx.Rollback()
				return
			}
		}
	}
	//创建价格梯度
	if product.ProductDetail.ProductLadderList != nil && len(product.ProductDetail.ProductLadderList) > 0 {
		for i := 0; i < len(product.ProductDetail.ProductLadderList); i++ {
			product.ProductDetail.ProductLadderList[i].ProductId = product.ProductDetail.Product.Id
			err = tx.Create(&product.ProductDetail.ProductLadderList[i]).Error
			if err != nil {
				tx.Rollback()
				return
			}
		}
	}
	//创建满减价格
	if product.ProductDetail.ProductFullReductionList != nil && len(product.ProductDetail.ProductFullReductionList) > 0 {
		for i := 0; i < len(product.ProductDetail.ProductFullReductionList); i++ {
			product.ProductDetail.ProductFullReductionList[i].ProductId = product.ProductDetail.Product.Id
			err = tx.Create(&product.ProductDetail.ProductFullReductionList[i]).Error
			if err != nil {
				tx.Rollback()
				return
			}
		}
	}
	//sku
	if product.ProductDetail.ProductSkuStockList != nil && len(product.ProductDetail.ProductSkuStockList) > 0 {
		for i := 0; i < len(product.ProductDetail.ProductSkuStockList); i++ {
			product.ProductDetail.ProductSkuStockList[i].ProductId = product.ProductDetail.Product.Id
			now := time.Now()
			product.ProductDetail.ProductSkuStockList[i].SkuCode = fmt.Sprint("%s%03d", now.Format("20060102"), i+1)
			err = tx.Create(&product.ProductDetail.ProductSkuStockList[i]).Error
			if err != nil {
				tx.Rollback()
				return
			}
		}
	}
	//商品属性
	if product.ProductDetail.ProductAttributeValueList != nil && len(product.ProductDetail.ProductAttributeValueList) > 0 {
		for i := 0; i < len(product.ProductDetail.ProductAttributeValueList); i++ {
			product.ProductDetail.ProductAttributeValueList[i].ProductId = product.ProductDetail.Product.Id
			err = tx.Create(&product.ProductDetail.ProductAttributeValueList[i]).Error
			if err != nil {
				tx.Rollback()
				return
			}
		}
	}
	tx.Commit()
	return product.ProductDetail.Product.Id, err
}

func (dao *productDaoIml) SimpleQuery(limit int64, pages int64, key string, startTime string, endTime string, order string) (rsp *productProto.Out_GetProducts, err error) {
	DB := db.MasterEngine()
	rsp = &productProto.Out_GetProducts{}
	offset := (pages - 1) * limit
	if len(key) == 0 {
		if len(startTime) > 0 && len(endTime) == 0 {
			err = DB.Model(&productProto.Product{}).Where("created_time > ?", endTime).Order(order).Count(&rsp.Total).Error
			if err == nil && rsp.Total > 0 {
				err = DB.Where("created_time > ? ", startTime).Order(order).Offset(offset).Limit(limit).Find(&rsp.ProductList).Error
			}
		} else if len(startTime) == 0 && len(endTime) > 0 {
			err = DB.Model(&productProto.Product{}).Where("created_time < ? ", endTime).Order(order).Count(&rsp.Total).Error
			if err == nil && rsp.Total > 0 {
				err = DB.Where("created_time < ? ", endTime).Order(order).Offset(offset).Limit(limit).Find(&rsp.ProductList).Error
			}
		} else if len(startTime) > 0 && len(endTime) > 0 {
			err = DB.Model(&productProto.Product{}).Where("created_time  between ? and ?", startTime, endTime).Order(order).Count(&rsp.Total).Error
			if err == nil && rsp.Total > 0 {
				err = DB.Where("created_time  between ? and ?", startTime, endTime).Order(order).Offset(offset).Limit(limit).Find(&rsp.ProductList).Error
			}
		} else {
			//先统计
			err = DB.Model(&productProto.Product{}).Order(order).Count(&rsp.Total).Error
			if err == nil && rsp.Total > 0 {
				err = DB.Order(order).Offset(offset).Limit(limit).Find(&rsp.ProductList).Error
			}
		}
	} else {
		searchKey := "%" + key + "%"
		if len(startTime) > 0 && len(endTime) == 0 {
			err = DB.Model(&productProto.Product{}).Where("(name like ? ) and created_time > ? ", searchKey, startTime).Order(order).Count(&rsp.Total).Error
			if err == nil && rsp.Total > 0 {
				err = DB.Model(&productProto.Product{}).Where("(name like ?) and created_time > ? ", searchKey, startTime).Order(order).Offset(offset).Limit(limit).Find(&rsp.ProductList).Error
			}
		} else if len(startTime) == 0 && len(endTime) > 0 {
			err = DB.Model(&productProto.Product{}).Where("(name like ?) and created_time < ? ", searchKey, endTime).Order(order).Count(&rsp.Total).Error
			if err == nil && rsp.Total > 0 {
				err = DB.Where("(name like ?) and created_time < ? ", searchKey, endTime).Order(order).Offset(offset).Limit(limit).Find(&rsp.ProductList).Error
			}
		} else if len(startTime) > 0 && len(endTime) > 0 {
			err = DB.Model(&productProto.Product{}).Where("(name like ?) and created_time between ? and ?", searchKey, startTime, endTime).Order(order).Count(&rsp.Total).Error
			if err == nil && rsp.Total > 0 {
				err = DB.Where("(name like ?) and created_time between ? and ?", searchKey, startTime, endTime).Order(order).Offset(offset).Limit(limit).Find(&rsp.ProductList).Error
			}
		} else {
			err = DB.Model(&productProto.Product{}).Where("name like ?", searchKey).Order(order).Count(&rsp.Total).Error
			if err == nil && rsp.Total > 0 {
				err = DB.Where("name like ?", searchKey).Order(order).Offset(offset).Limit(limit).Find(&rsp.ProductList).Error
			}
		}
	}
	return
}

func (dao *productDaoIml) Delete(deleteStatus int64, ids []string) (err error) {
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
		err = DB.Model(&productProto.Product{}).Where("id = ?", ids[i]).Update("delete_status", deleteStatus).Error
		if err != nil {
			break
		}
	}
	if err != nil {
		DB.Commit()
	}
	return
}

func (dao *productDaoIml) Update(product *productProto.ProductDetails) (err error) {
	DB := db.MasterEngine()
	tx := DB.Begin()
	defer func() {
		if re := recover(); re != nil {
			tx.Rollback()
			if err == nil {
				err = errors.New(fmt.Sprintf("%v", re))
			}
		}
	}()
	product.Product.CreatedTime, err = time_util.TimeTtringToTimeString(product.Product.CreatedTime)
	if err != nil {
		return
	}
	if len(product.Product.PromotionStartTime) > 0 {
		product.Product.PromotionStartTime, err = time_util.TimeTtringToTimeString(product.Product.PromotionStartTime)
		if err != nil {
			return
		}
	}
	if len(product.Product.PromotionEndTime) > 0 {
		product.Product.PromotionEndTime, err = time_util.TimeTtringToTimeString(product.Product.PromotionEndTime)
		if err != nil {
			return
		}
	}

	//先插入商品信息
	product.Product.UpdateTime = time.Now().Format("2006-01-02 15:04:05")
	err = tx.Model(&productProto.Product{}).Where("id = ?", product.Product.Id).Update(&product.Product).Error
	if err != nil {
		tx.Rollback()
		return
	}
	//创建会员价格
	//先删除旧数据
	DB.Where("product_id = ?", product.Product.Id).Delete(&productProto.ProductMemberPrice{})
	if product.ProductMemberPriceList != nil && len(product.ProductMemberPriceList) > 0 {
		for i := 0; i < len(product.ProductMemberPriceList); i++ {
			product.ProductMemberPriceList[i].ProductId = product.Product.Id
			standStrict := memberBean.ProductMemberPrice{}
			item := product.ProductMemberPriceList[i]
			err = memberBean.ProToStandard(item, &standStrict)
			if err != nil {
				tx.Rollback()
				return
			}
			err = tx.Create(&standStrict).Error
			if err != nil {
				tx.Rollback()
				return
			}
		}
	}
	//创建价格梯度
	DB.Where("product_id = ?", product.Product.Id).Delete(&productProto.ProductLadder{})
	if product.ProductLadderList != nil && len(product.ProductLadderList) > 0 {
		for i := 0; i < len(product.ProductLadderList); i++ {
			product.ProductLadderList[i].ProductId = product.Product.Id
			standStruck := ladderBean.ProductLadder{}
			err = ladderBean.ProToStandard(product.ProductLadderList[i], &standStruck)
			if err != nil {
				tx.Rollback()
				return
			}
			err = tx.Create(&standStruck).Error
			if err != nil {
				tx.Rollback()
				return
			}
		}
	}
	//创建满减价格
	DB.Where("product_id = ?", product.Product.Id).Delete(&productProto.ProductFullReduction{})
	if product.ProductFullReductionList != nil && len(product.ProductFullReductionList) > 0 {
		for i := 0; i < len(product.ProductFullReductionList); i++ {
			product.ProductFullReductionList[i].ProductId = product.Product.Id
			standStruck := fullReductionBean.ProductFullReduction{}
			err = fullReductionBean.ProToStandard(product.ProductFullReductionList[i], &standStruck)
			if err != nil {
				tx.Rollback()
				return
			}
			err = tx.Create(&standStruck).Error
			if err != nil {
				tx.Rollback()
				return
			}
		}
	}
	//sku
	DB.Where("product_id = ?", product.Product.Id).Delete(&productProto.ProductSkuStock{})
	if product.ProductSkuStockList != nil && len(product.ProductSkuStockList) > 0 {
		for i := 0; i < len(product.ProductSkuStockList); i++ {
			product.ProductSkuStockList[i].ProductId = product.Product.Id
			now := time.Now()
			product.ProductSkuStockList[i].SkuCode = fmt.Sprintf("%s%03d", now.Format("20060102"), i+1)
			standStruct := skuStockBean.ProductSkuStock{}
			err = skuStockBean.ProToStandard(product.ProductSkuStockList[i], &standStruct)
			if err != nil {
				tx.Rollback()
				return
			}
			err = tx.Create(&standStruct).Error
			if err != nil {
				tx.Rollback()
				return
			}
		}
	}
	//商品属性
	DB.Where("product_id = ?", product.Product.Id).Delete(&productProto.ProductAttributeValue{})
	if product.ProductAttributeValueList != nil && len(product.ProductAttributeValueList) > 0 {
		for i := 0; i < len(product.ProductAttributeValueList); i++ {
			product.ProductAttributeValueList[i].ProductId = product.Product.Id
			standStruct := attributeValueBean.ProductAttributeValue{}
			err = attributeValueBean.ProToStandard(product.ProductAttributeValueList[i], &standStruct)
			if err != nil {
				tx.Rollback()
				return
			}
			err = tx.Create(&standStruct).Error
			if err != nil {
				tx.Rollback()
				return
			}
		}
	}
	tx.Commit()
	return
}
