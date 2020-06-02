package service

import (
	"context"
	"github.com/go-log/log"
	productProto "github.com/qianxunke/ego-shopping/ego-common-protos/go_out/inventory/product"
	productDao "inventory-service/modules/product/dao"
	"net/http"
)

//获取商品信息
func (s *service) GetProductById(ctx context.Context, req *productProto.In_GetProductById) (rsp *productProto.Out_GetProductById, err error) {
	rsp = &productProto.Out_GetProductById{}
	rsp.Product = &productProto.ProductDetails{}
	rsp.Error = &productProto.Error{}
	if len(req.Id) <= 0 {
		rsp.Error = &productProto.Error{
			Code:    http.StatusBadRequest,
			Message: "请求参数有误！",
		}
		return
	}
	dao, err := productDao.GetProductDao()
	if err != nil {
		rsp.Error = &productProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}
	rsp.Product, err = dao.FindById(req.Id)
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
	rsp = &productProto.Out_UpdateProductInfo{}
	dao, err := productDao.GetProductDao()
	if err != nil {
		rsp.Error = &productProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}
	err = dao.Update(req.Product)
	if err != nil {
		rsp.Error = &productProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}
	go s.sendUpdateProductMassageToEs(req.Product.Product.Id)

	rsp.Error = &productProto.Error{
		Code:    http.StatusOK,
		Message: "修改成功！",
	}
	return

}

//获取列表
func (s *service) GetProducts(ctx context.Context, req *productProto.In_GetProducts) (rsp *productProto.Out_GetProducts, err error) {
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
	orderByStr := "created_time DESC"
	dao, err := productDao.GetProductDao()
	if err != nil {
		log.Logf("ERROR: %v", err)
		rsp.Error = &productProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}
	rsp, err = dao.SimpleQuery(req.Limit, req.Pages, req.SearchKey, req.StartTime, req.EndTime, orderByStr)
	if err != nil {
		log.Logf("ERROR: %v", err)
		rsp.Error = &productProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}
	var message string
	if rsp.Total > 0 && len(rsp.ProductList) > 0 {
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
	rsp = &productProto.Out_DeleteProducts{}
	if len(req.ProductList) <= 0 {
		rsp.Error = &productProto.Error{
			Code:    http.StatusBadRequest,
			Message: "参数不正确",
		}
		return
	}
	dao, err := productDao.GetProductDao()
	if err != nil {
		log.Logf("ERROR: %v", err)
		rsp.Error = &productProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	err = dao.Delete(req.DeleteStatus, req.ProductList)
	if err != nil {
		rsp.Error = &productProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}
	go s.sendDeleteProductMassageToEs(req.ProductList)
	rsp.Error = &productProto.Error{
		Code:    http.StatusOK,
		Message: "删除成功!",
	}
	return
}

//新建信息
func (s *service) CreateProduct(ctx context.Context, req *productProto.In_CreateProduct) (rsp *productProto.Out_CreateProduct, err error) {
	rsp = &productProto.Out_CreateProduct{}
	//查询该等级是否存在
	dao, err := productDao.GetProductDao()
	if err != nil {
		log.Logf("error: %v", err)
		rsp.Error = &productProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	id, err := dao.Insert(req)
	if err != nil {
		rsp.Error = &productProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}
	go s.sendAddProductMassageToEs(id)

	rsp.Error = &productProto.Error{
		Code:    http.StatusOK,
		Message: "新增成功!",
	}
	return
}
