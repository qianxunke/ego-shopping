package service

import (
	"context"
	"github.com/micro/go-micro/util/log"
	"net/http"
	productCategoryDao "inventory-service/modules/product_category/dao"
	productCategoryProto "github.com/qianxunke/ego-shopping/ego-common-protos/go_out/inventory/product_category"
	"reflect"
)

//获取信息
func (s *service) GetProductCategoryById(ctx context.Context,req *productCategoryProto.In_GetProductCategoryById) (rsp *productCategoryProto.Out_GetProductCategoryById,err error) {
	rsp = &productCategoryProto.Out_GetProductCategoryById{}
	rsp.Error = &productCategoryProto.Error{}
	if req.Id <= 0 {
		rsp.Error = &productCategoryProto.Error{
			Code:    http.StatusBadRequest,
			Message: "请求参数有误！",
		}
		return
	}
	dao, err := productCategoryDao.GetDao()
	if err != nil {
		rsp.Error = &productCategoryProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}
	rsp.ProductCategory, err = dao.FindById(req.Id)
	if err != nil {
		rsp.Error = &productCategoryProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}
	rsp.Error = &productCategoryProto.Error{
		Code:    http.StatusOK,
		Message: "查询成功！",
	}
	return

}

//修改信息
func (s *service) UpdateProductCategoryInfo(ctx context.Context,req *productCategoryProto.In_UpdateProductCategoryInfo) (rsp *productCategoryProto.Out_UpdateProductCategoryInfo,err error) {
	rsp = &productCategoryProto.Out_UpdateProductCategoryInfo{}
	updataData := map[string]interface{}{}
	elem := reflect.ValueOf(&req.ProductCategory).Elem()
	relType := elem.Type()
	for i := 0; i < relType.NumField(); i++ {
		updataData[relType.Field(i).Name] = elem.Field(i).Interface()
	}
	dao, err := productCategoryDao.GetDao()
	if err != nil {
		rsp.Error = &productCategoryProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}
	delete(updataData, "id")
	err = dao.Update(req.ProductCategory.Id, updataData)
	if err != nil {
		rsp.Error = &productCategoryProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}
	rsp.Error = &productCategoryProto.Error{
		Code:    http.StatusOK,
		Message: "修改成功！",
	}
	return
}

//获取列表
func (s *service) GetProductCategorys(ctx context.Context,req *productCategoryProto.In_GetProductCategorys) (rsp *productCategoryProto.Out_GetProductCategorys,err error) {
	rsp = &productCategoryProto.Out_GetProductCategorys{}
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
	dao, err := productCategoryDao.GetDao()
	if err != nil {
		log.Logf("ERROR: %v", err)
		rsp.Error = &productCategoryProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}
	rsp, err = dao.SimpleQuery(req.Limit, req.Pages, req.SearchKey, req.StartTime, req.EndTime, orderByStr)
	if err != nil {
		log.Logf("ERROR: %v", err)
		rsp.Error = &productCategoryProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}
	var message string
	if rsp.Total > 0 && len(rsp.ProductCategoryList) > 0 {
		message = "查询成功！"
	} else {
		message = "没有数据了！"
	}
	//统计有多少条
	rsp.Error = &productCategoryProto.Error{
		Code:    http.StatusOK,
		Message: message,
	}
	rsp.Limit = req.Limit
	rsp.Pages = req.Pages
	return

}

//删除列表
func (s *service) DeleteProductCategorys(ctx context.Context,req *productCategoryProto.In_DeleteProductCategorys) (rsp *productCategoryProto.Out_DeleteProductCategorys,err error) {
	rsp = &productCategoryProto.Out_DeleteProductCategorys{}
	if len(req.ProductCategoryList) <= 0 {
		rsp.Error = &productCategoryProto.Error{
			Code:    http.StatusBadRequest,
			Message: "参数不正确",
		}
		return
	}
	dao, err := productCategoryDao.GetDao()
	if err != nil {
		log.Logf("ERROR: %v", err)
		rsp.Error = &productCategoryProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	err = dao.Delete(req.ProductCategoryList)
	if err != nil {
		rsp.Error = &productCategoryProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}
	rsp.Error = &productCategoryProto.Error{
		Code:    http.StatusOK,
		Message: "删除成功!",
	}
	return
}

//新建信息
func (s *service) CreateProductCategory(ctx context.Context,req *productCategoryProto.In_CreateProductCategory) (rsp *productCategoryProto.Out_CreateProductCategory,err error) {
	rsp = &productCategoryProto.Out_CreateProductCategory{}
	//查询该等级是否存在
	dao, err := productCategoryDao.GetDao()
	if err != nil {
		log.Logf("ERROR: %v", err)
		rsp.Error = &productCategoryProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	err = dao.Insert(req.ProductCategory)
	if err != nil {
		rsp.Error = &productCategoryProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}
	rsp.Error = &productCategoryProto.Error{
		Code:    http.StatusOK,
		Message: "新增成功!",
	}
	return
}

func (s *service) GetProductCategoryDetailsList(ctx context.Context,req *productCategoryProto.In_GetProductCategoryDetailsList) (rsp *productCategoryProto.Out_GetProductCategoryDetailsList,err error) {
	rsp = &productCategoryProto.Out_GetProductCategoryDetailsList{}
	orderByStr := "created_time DESC"
	dao, err := productCategoryDao.GetDao()
	if err != nil {
		log.Logf("ERROR: %v", err)
		rsp.Error = &productCategoryProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}
	rsp, err = dao.GetProductCategoryDetailsList(1000, 1, orderByStr)
	if err != nil {
		log.Logf("ERROR: %v", err)
		rsp.Error = &productCategoryProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}
	var message string
	if rsp.Total > 0 && len(rsp.ProductCategoryList) > 0 {
		message = "查询成功！"
	} else {
		message = "没有数据了！"
	}
	//统计有多少条
	rsp.Error = &productCategoryProto.Error{
		Code:    http.StatusOK,
		Message: message,
	}
	rsp.Limit = 1000
	rsp.Pages = 1
	return

}
