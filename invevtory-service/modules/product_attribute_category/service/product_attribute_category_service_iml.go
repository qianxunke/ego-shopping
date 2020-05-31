package service

import (
	"context"
	"github.com/micro/go-micro/util/log"
	"net/http"
	productAttributeCategoryDao "inventory-service/modules/product_attribute_category/dao"
	productAttributeCategoryProto "github.com/qianxunke/ego-shopping/ego-common-protos/go_out/inventory/product_attribute_category"
	"reflect"
)

//获取信息
func (s *service) GetProductAttributeCategoryById(ctx context.Context,req *productAttributeCategoryProto.In_GetProductAttributeCategoryById) (rsp *productAttributeCategoryProto.Out_GetProductAttributeCategoryById,err error) {
	rsp = &productAttributeCategoryProto.Out_GetProductAttributeCategoryById{}
	rsp.Error = &productAttributeCategoryProto.Error{}
	if req.Id <= 0 {
		rsp.Error = &productAttributeCategoryProto.Error{
			Code:    http.StatusBadRequest,
			Message: "请求参数有误！",
		}
		return
	}
	dao, err := productAttributeCategoryDao.GetDao()
	if err != nil {
		rsp.Error = &productAttributeCategoryProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}
	rsp.ProductAttributeCategory, err = dao.FindById(req.Id)
	if err != nil {
		rsp.Error = &productAttributeCategoryProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}
	rsp.Error = &productAttributeCategoryProto.Error{
		Code:    http.StatusOK,
		Message: "查询成功！",
	}
	return

}

//修改信息
func (s *service) UpdateProductAttributeCategoryInfo(ctx context.Context,req *productAttributeCategoryProto.In_UpdateProductAttributeCategoryInfo) (rsp *productAttributeCategoryProto.Out_UpdateProductAttributeCategoryInfo,err error) {
	rsp = &productAttributeCategoryProto.Out_UpdateProductAttributeCategoryInfo{}
	updataData := map[string]interface{}{}
	elem := reflect.ValueOf(&req.ProductAttributeCategory).Elem()
	relType := elem.Type()
	for i := 0; i < relType.NumField(); i++ {
		updataData[relType.Field(i).Name] = elem.Field(i).Interface()
	}
	dao, err := productAttributeCategoryDao.GetDao()
	if err != nil {
		rsp.Error = &productAttributeCategoryProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}
	delete(updataData, "id")
	err = dao.Update(req.ProductAttributeCategory.Id, updataData)
	if err != nil {
		rsp.Error = &productAttributeCategoryProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}
	rsp.Error = &productAttributeCategoryProto.Error{
		Code:    http.StatusOK,
		Message: "修改成功！",
	}
	return
}

//获取列表
func (s *service) GetProductAttributeCategorys(ctx context.Context,req *productAttributeCategoryProto.In_GetProductAttributeCategorys) (rsp *productAttributeCategoryProto.Out_GetProductAttributeCategorys,err error) {
	rsp = &productAttributeCategoryProto.Out_GetProductAttributeCategorys{}
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
	orderByStr := "created_time DESC"
	dao, err := productAttributeCategoryDao.GetDao()
	if err != nil {
		log.Logf("ERROR: %v", err)
		rsp.Error = &productAttributeCategoryProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}
	rsp, err = dao.SimpleQuery(req.Limit, req.Pages, req.SearchKey, req.StartTime, req.EndTime, orderByStr)
	if err != nil {
		log.Logf("ERROR: %v", err)
		rsp.Error = &productAttributeCategoryProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}
	var message string
	if rsp.Total > 0 && len(rsp.ProductAttributeCategoryList) > 0 {
		message = "查询成功！"
	} else {
		message = "没有数据了！"
	}
	//统计有多少条
	rsp.Error = &productAttributeCategoryProto.Error{
		Code:    http.StatusOK,
		Message: message,
	}
	rsp.Limit = req.Limit
	rsp.Pages = req.Pages
	return

}

//删除列表
func (s *service) DeleteProductAttributeCategorys(ctx context.Context,req *productAttributeCategoryProto.In_DeleteProductAttributeCategorys) (rsp *productAttributeCategoryProto.Out_DeleteProductAttributeCategorys,err error) {
	rsp = &productAttributeCategoryProto.Out_DeleteProductAttributeCategorys{}
	if len(req.ProductAttributeCategoryList) <= 0 {
		rsp.Error = &productAttributeCategoryProto.Error{
			Code:    http.StatusBadRequest,
			Message: "参数不正确",
		}
		return
	}
	dao, err := productAttributeCategoryDao.GetDao()
	if err != nil {
		log.Logf("ERROR: %v", err)
		rsp.Error = &productAttributeCategoryProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	err = dao.Delete(req.ProductAttributeCategoryList)
	if err != nil {
		rsp.Error = &productAttributeCategoryProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}
	rsp.Error = &productAttributeCategoryProto.Error{
		Code:    http.StatusOK,
		Message: "删除成功!",
	}
	return
}

//新建信息
func (s *service) CreateProductAttributeCategory(ctx context.Context,req *productAttributeCategoryProto.In_CreateProductAttributeCategory) (rsp *productAttributeCategoryProto.Out_CreateProductAttributeCategory,err error) {
	rsp = &productAttributeCategoryProto.Out_CreateProductAttributeCategory{}
	//查询该等级是否存在
	dao, err := productAttributeCategoryDao.GetDao()
	if err != nil {
		log.Logf("ERROR: %v", err)
		rsp.Error = &productAttributeCategoryProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	err = dao.Insert(req.ProductAttributeCategory)
	if err != nil {
		rsp.Error = &productAttributeCategoryProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}
	rsp.Error = &productAttributeCategoryProto.Error{
		Code:    http.StatusOK,
		Message: "新增成功!",
	}
	return
}
