package service

import (
	"context"
	"github.com/go-log/log"
	productSkuStockProto "github.com/qianxunke/ego-shopping/ego-common-protos/go_out/inventory/product_sku_stock"
	"github.com/qianxunke/ego-shopping/ego-plugins/lock"
	"github.com/qianxunke/ego-shopping/ego-plugins/lock/etcd"
	productSkuStockDao "inventory-service/modules/product_sku_stock/dao"
	"inventory-service/utils/string_utl"
	"net/http"
	"reflect"
	"time"
)

//获取信息
func (s *service) GetProductSkuStockById(ctx context.Context, req *productSkuStockProto.In_GetProductSkuStockById) (rsp *productSkuStockProto.Out_GetProductSkuStockById, err error) {
	rsp = &productSkuStockProto.Out_GetProductSkuStockById{}
	rsp.Error = &productSkuStockProto.Error{}
	if req.Id <= 0 {
		rsp.Error = &productSkuStockProto.Error{
			Code:    http.StatusBadRequest,
			Message: "请求参数有误！",
		}
		return
	}
	dao, err := productSkuStockDao.GetDao()
	if err != nil {
		rsp.Error = &productSkuStockProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}
	rsp.ProductSkuStock, err = dao.FindById(req.Id)
	if err != nil {
		rsp.Error = &productSkuStockProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}
	rsp.Error = &productSkuStockProto.Error{
		Code:    http.StatusOK,
		Message: "查询成功！",
	}
	return

}

//修改信息
func (s *service) UpdateProductSkuStockInfo(ctx context.Context, req *productSkuStockProto.In_UpdateProductSkuStockInfo) (rsp *productSkuStockProto.Out_UpdateProductSkuStockInfo, err error) {
	rsp = &productSkuStockProto.Out_UpdateProductSkuStockInfo{}
	updateData := map[string]interface{}{}
	elem := reflect.ValueOf(&req.ProductSkuStock).Elem()
	relType := elem.Type()
	for i := 0; i < relType.NumField(); i++ {
		updateData[relType.Field(i).Name] = elem.Field(i).Interface()
	}
	dao, err := productSkuStockDao.GetDao()
	if err != nil {
		rsp.Error = &productSkuStockProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}
	delete(updateData, "id")
	err = dao.Update(req.ProductSkuStock.Id, updateData)
	if err != nil {
		rsp.Error = &productSkuStockProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}
	rsp.Error = &productSkuStockProto.Error{
		Code:    http.StatusOK,
		Message: "修改成功！",
	}
	return
}

//获取列表
func (s *service) GetProductSkuStocks(ctx context.Context, req *productSkuStockProto.In_GetProductSkuStocks) (rsp *productSkuStockProto.Out_GetProductSkuStocks, err error) {
	rsp = &productSkuStockProto.Out_GetProductSkuStocks{}
	dao, err := productSkuStockDao.GetDao()
	if err != nil {
		rsp.Error = &productSkuStockProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}
	rsp, err = dao.SimpleQuery(req.ProductId)
	if err != nil {
		rsp.Error = &productSkuStockProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}
	var message string
	if rsp.Total > 0 && len(rsp.ProductSkuStockList) > 0 {
		message = "查询成功！"
	} else {
		message = "没有数据了！"
	}
	//统计有多少条
	rsp.Error = &productSkuStockProto.Error{
		Code:    http.StatusOK,
		Message: message,
	}
	rsp.Limit = 0
	rsp.Pages = 0
	return

}

//删除列表
func (s *service) DeleteProductSkuStocks(ctx context.Context, req *productSkuStockProto.In_DeleteProductSkuStocks) (rsp *productSkuStockProto.Out_DeleteProductSkuStocks, err error) {
	rsp = &productSkuStockProto.Out_DeleteProductSkuStocks{}
	if len(req.ProductSkuStockList) <= 0 {
		rsp.Error = &productSkuStockProto.Error{
			Code:    http.StatusBadRequest,
			Message: "参数不正确",
		}
		return
	}
	dao, err := productSkuStockDao.GetDao()
	if err != nil {
		log.Logf("ERROR: %v", err)
		rsp.Error = &productSkuStockProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	err = dao.Delete(req.ProductSkuStockList)
	if err != nil {
		rsp.Error = &productSkuStockProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}
	rsp.Error = &productSkuStockProto.Error{
		Code:    http.StatusOK,
		Message: "删除成功!",
	}
	return
}

//新建信息
func (s *service) CreateProductSkuStock(ctx context.Context, req *productSkuStockProto.In_CreateProductSkuStock) (rsp *productSkuStockProto.Out_CreateProductSkuStock, err error) {
	rsp = &productSkuStockProto.Out_CreateProductSkuStock{}
	//查询该等级是否存在
	dao, err := productSkuStockDao.GetDao()
	if err != nil {
		log.Logf("ERROR: %v", err)
		rsp.Error = &productSkuStockProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	err = dao.Insert(req.ProductSkuStock)
	if err != nil {
		rsp.Error = &productSkuStockProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}
	rsp.Error = &productSkuStockProto.Error{
		Code:    http.StatusOK,
		Message: "新增成功!",
	}
	return
}

func (s *service) GetSellSkuStock(ctx context.Context, req *productSkuStockProto.In_GetSellSkuStock) (rsp *productSkuStockProto.Out_GetSellSkuStock, err error) {
	rsp = &productSkuStockProto.Out_GetSellSkuStock{}
	rsp.SkuOutLog = &productSkuStockProto.SkuOutLog{}
	id := req.ProductId + "[" + string_utl.Int64ToString(req.Id) + "]"
	//开始锁库
	etcdLock := etcd.NewLock(lock.Nodes("localhost:2379"))
	err = etcdLock.Acquire(id, lock.TTL(time.Duration(time.Second*5)), lock.Wait(time.Duration(time.Second*5)))
	if err != nil {
		rsp.Error = &productSkuStockProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}
	dao, err := productSkuStockDao.GetDao()
	if err != nil {
		log.Logf("ERROR: %v", err)
		rsp.Error = &productSkuStockProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	rspValue, err := dao.UpdateStock(req.Id, req.ProductId, req.MemberId, req.SellNum)
	if err != nil {
		rsp.Error = &productSkuStockProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	rsp.SkuOutLog = rspValue
	rsp.Error = &productSkuStockProto.Error{
		Code:    http.StatusOK,
		Message: "ok",
	}
	_ = etcdLock.Release(id)
	return
}
