package service

import (
	"context"
	"github.com/micro/go-micro/util/log"
	productAttributeProto "github.com/qianxunke/ego-shopping/ego-common-protos/go_out/inventory/product_attribute"
	productAttributeDao "inventory-service/modules/product_attribute/dao"
	"net/http"
	"reflect"
)

//获取信息
func (s *service) GetProductAttributeById(ctx context.Context, req *productAttributeProto.In_GetProductAttributeById) (rsp *productAttributeProto.Out_GetProductAttributeById, err error) {
	rsp = &productAttributeProto.Out_GetProductAttributeById{}
	rsp.Error = &productAttributeProto.Error{}
	if req.Id <= 0 {
		rsp.Error = &productAttributeProto.Error{
			Code:    http.StatusBadRequest,
			Message: "请求参数有误！",
		}
		return
	}
	dao, err := productAttributeDao.GetDao()
	if err != nil {
		rsp.Error = &productAttributeProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}
	rsp.ProductAttribute, err = dao.FindById(req.Id)
	if err != nil {
		rsp.Error = &productAttributeProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}
	rsp.Error = &productAttributeProto.Error{
		Code:    http.StatusOK,
		Message: "查询成功！",
	}
	return

}

//修改信息
func (s *service) UpdateProductAttributeInfo(ctx context.Context, req *productAttributeProto.In_UpdateProductAttributeInfo) (rsp *productAttributeProto.Out_UpdateProductAttributeInfo, err error) {
	rsp = &productAttributeProto.Out_UpdateProductAttributeInfo{}
	updataData := map[string]interface{}{}
	elem := reflect.ValueOf(&req.ProductAttribute).Elem()
	relType := elem.Type()
	for i := 0; i < relType.NumField(); i++ {
		updataData[relType.Field(i).Name] = elem.Field(i).Interface()
	}
	dao, err := productAttributeDao.GetDao()
	if err != nil {
		rsp.Error = &productAttributeProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}
	delete(updataData, "id")
	err = dao.Update(req.ProductAttribute.Id, updataData)
	if err != nil {
		rsp.Error = &productAttributeProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}
	rsp.Error = &productAttributeProto.Error{
		Code:    http.StatusOK,
		Message: "修改成功！",
	}
	return
}

//获取列表
func (s *service) GetProductAttributes(ctx context.Context, req *productAttributeProto.In_GetProductAttributes) (rsp *productAttributeProto.Out_GetProductAttributes, err error) {
	rsp = &productAttributeProto.Out_GetProductAttributes{}
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
	orderByStr := "sort ASC"
	dao, err := productAttributeDao.GetDao()
	if err != nil {
		log.Logf("ERROR: %v", err)
		rsp.Error = &productAttributeProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}
	rsp, err = dao.SimpleQuery(req.Limit, req.Pages, req.SearchKey, req.StartTime, req.EndTime, orderByStr)
	if err != nil {
		log.Logf("ERROR: %v", err)
		rsp.Error = &productAttributeProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}
	var message string
	if rsp.Total > 0 && len(rsp.ProductAttributeList) > 0 {
		message = "查询成功！"
	} else {
		message = "没有数据了！"
	}
	//统计有多少条
	rsp.Error = &productAttributeProto.Error{
		Code:    http.StatusOK,
		Message: message,
	}
	rsp.Limit = req.Limit
	rsp.Pages = req.Pages
	return

}

//删除列表
func (s *service) DeleteProductAttributes(ctx context.Context, req *productAttributeProto.In_DeleteProductAttributes) (rsp *productAttributeProto.Out_DeleteProductAttributes, err error) {
	rsp = &productAttributeProto.Out_DeleteProductAttributes{}
	if len(req.ProductAttributeList) <= 0 {
		rsp.Error = &productAttributeProto.Error{
			Code:    http.StatusBadRequest,
			Message: "参数不正确",
		}
		return
	}
	dao, err := productAttributeDao.GetDao()
	if err != nil {
		log.Logf("ERROR: %v", err)
		rsp.Error = &productAttributeProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	err = dao.Delete(req.ProductAttributeList)
	if err != nil {
		rsp.Error = &productAttributeProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}
	rsp.Error = &productAttributeProto.Error{
		Code:    http.StatusOK,
		Message: "删除成功!",
	}
	return
}

//新建信息
func (s *service) CreateProductAttribute(ctx context.Context, req *productAttributeProto.In_CreateProductAttribute) (rsp *productAttributeProto.Out_CreateProductAttribute, err error) {
	rsp = &productAttributeProto.Out_CreateProductAttribute{}
	//查询该等级是否存在
	dao, err := productAttributeDao.GetDao()
	if err != nil {
		log.Logf("ERROR: %v", err)
		rsp.Error = &productAttributeProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	err = dao.Insert(req.ProductAttribute)
	if err != nil {
		rsp.Error = &productAttributeProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}
	rsp.Error = &productAttributeProto.Error{
		Code:    http.StatusOK,
		Message: "新增成功!",
	}
	return
}

func (s *service) GetProductAttributeList(ctx context.Context, req *productAttributeProto.In_GetProductAttributeList) (rsp *productAttributeProto.Out_GetProductAttributeList, err error) {
	log.Log("Received ProductAttribute.GetProductAttributeList ")
	rsp = &productAttributeProto.Out_GetProductAttributeList{}
	dao, err := productAttributeDao.GetDao()
	if err != nil {
		log.Logf("ERROR: %v", err)
		rsp.Error = &productAttributeProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}
	rsp, err = dao.GetProductAttributeList(req.Cid, req.TypeValue)
	if err != nil {
		log.Logf("ERROR: %v", err)
		rsp.Error = &productAttributeProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}
	var message string
	if rsp.Total > 0 && len(rsp.ProductAttributeList) > 0 {
		message = "查询成功！"
	} else {
		message = "没有数据了！"
	}
	//统计有多少条
	rsp.Error = &productAttributeProto.Error{
		Code:    http.StatusOK,
		Message: message,
	}
	return
}
