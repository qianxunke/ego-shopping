package service

import (
	"context"
	productDao "ego-goods-search-service/modules/product_es/dao"
	"github.com/micro/go-micro/util/log"
	productProto "github.com/qianxunke/ego-shopping/ego-common-protos/go_out/inventory/product"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

//获取商品信息
func (s *service) GetProductById(ctx context.Context, req *productProto.In_GetProductById) (rsp *productProto.Out_GetProductById, err error) {
	defer func() {
		if err := recover(); err != nil {
			log.Logf("%v", err)
			rsp.Error = &productProto.Error{
				Message: "ERROE",
				Code:    http.StatusInternalServerError,
			}
		}
	}()
	rsp = &productProto.Out_GetProductById{}
	rsp.Error = &productProto.Error{}
	if len(req.Id) <= 0 {
		rsp.Error = &productProto.Error{
			Code:    http.StatusBadRequest,
			Message: "请求参数有误！",
		}
		return
	}
	dao, err := productDao.GetDao()
	if err != nil {
		rsp.Error = &productProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}
	rsp.Product.Product, err = dao.FindById(req.Id)
	if err != nil {
		rsp.Error = &productProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}
	rsp.Error = &productProto.Error{
		Code:    http.StatusOK,
		Message: "查询成功！",
	}
	return
}

//修改信息
func (s *service) UpdateProductInfo(ctx context.Context, req *productProto.In_UpdateProductInfo) (rsp *productProto.Out_UpdateProductInfo, err error) {
	defer func() {
		if err := recover(); err != nil {
			log.Logf("%v", err)
			rsp.Error = &productProto.Error{
				Message: "ERROE",
				Code:    http.StatusInternalServerError,
			}
		}
	}()
	rsp = &productProto.Out_UpdateProductInfo{}
	updateData := map[string]interface{}{}
	elem := reflect.ValueOf(req.Product).Elem()
	relType := elem.Type()
	for i := 0; i < relType.NumField(); i++ {
		if strings.Contains(relType.Field(i).Name, "XXX") {
			continue
		}
		updateData[relType.Field(i).Name] = elem.Field(i).Interface()
	}
	dao, err := productDao.GetDao()
	if err != nil {
		rsp.Error = &productProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}
	delete(updateData, "id")
	err = dao.Update(req.Product.Product.Id, updateData)
	if err != nil {
		rsp.Error = &productProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}
	rsp.Error = &productProto.Error{
		Code:    http.StatusOK,
		Message: "修改成功！",
	}
	return

}

//获取列表
func (s *service) GetProducts(ctx context.Context, req *productProto.In_GetProducts) (rsp *productProto.Out_GetProducts, err error) {
	defer func() {
		if err := recover(); err != nil {
			log.Logf("%v", err)
			rsp.Error = &productProto.Error{
				Message: "ERROE",
				Code:    http.StatusInternalServerError,
			}
		}
	}()
	rsp = &productProto.Out_GetProducts{}
	//对参数鉴权
	if req.Limit == 0 {
		req.Limit = 10 //默认10个分页
	}
	if req.Limit > 1000 { //每一页数量
		req.Limit = 1000
	}
	if req.Pages <= 0 { //页数
		req.Pages = 1
	}
	orderByStr := "created_time"
	dao, err := productDao.GetDao()
	if err != nil {
		log.Logf("ERROR: %v", err)
		rsp.Error = &productProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}
	limitStrInt64 := strconv.FormatInt(req.Limit, 10)
	limit, _ := strconv.Atoi(limitStrInt64)
	pageStrInt64 := strconv.FormatInt(req.Pages, 10)
	pages, _ := strconv.Atoi(pageStrInt64)
	rsp, err = dao.SimpleQuery(limit, pages, req.SearchKey, req.StartTime, req.EndTime, orderByStr)
	if err != nil || rsp == nil {
		rsp = &productProto.Out_GetProducts{}
		rsp.Error = &productProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}
	var message string
	if rsp.Total > 0 {
		message = "查询成功！"
	} else {
		message = "没有数据了！"
	}
	//统计有多少条
	rsp.Error = &productProto.Error{
		Code:    http.StatusOK,
		Message: message,
	}
	rsp.Limit = req.Limit
	rsp.Pages = req.Pages
	return

}

//批量删除等级列表
func (s *service) DeleteProducts(ctx context.Context, req *productProto.In_DeleteProducts) (rsp *productProto.Out_DeleteProducts, err error) {
	defer func() {
		if err := recover(); err != nil {
			log.Logf("%v", err)
			rsp.Error = &productProto.Error{
				Message: "ERROE",
				Code:    http.StatusInternalServerError,
			}
		}
	}()
	rsp = &productProto.Out_DeleteProducts{}
	if len(req.ProductList) <= 0 {
		rsp.Error = &productProto.Error{
			Code:    http.StatusBadRequest,
			Message: "参数不正确",
		}
		return
	}
	dao, err := productDao.GetDao()
	if err != nil {
		log.Logf("ERROR: %v", err)
		rsp.Error = &productProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	err = dao.Delete(req.ProductList)
	if err != nil {
		rsp.Error = &productProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}
	rsp.Error = &productProto.Error{
		Code:    http.StatusOK,
		Message: "删除成功!",
	}
	return
}

//新建信息
func (s *service) CreateProduct(ctx context.Context, req *productProto.In_CreateProduct) (rsp *productProto.Out_CreateProduct, err error) {
	defer func() {
		if err := recover(); err != nil {
			log.Logf("%v", err)
			rsp.Error = &productProto.Error{
				Message: "ERROE",
				Code:    http.StatusInternalServerError,
			}
		}
	}()
	rsp = &productProto.Out_CreateProduct{}
	//查询该等级是否存在
	dao, err := productDao.GetDao()
	if err != nil {
		log.Logf("ERROR: %v", err)
		rsp.Error = &productProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	err = dao.Insert(req.ProductDetail.Product)
	if err != nil {
		rsp.Error = &productProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}
	rsp.Error = &productProto.Error{
		Code:    http.StatusOK,
		Message: "新增成功!",
	}
	return
}
