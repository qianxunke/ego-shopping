package service

import (
	"context"
	"github.com/micro/go-micro/util/log"
	"net/http"
	brandDao "inventory-service/modules/brand/dao"
	branProto "github.com/qianxunke/ego-shopping/ego-common-protos/go_out/inventory/brand"
	"reflect"
)

//获取信息
func (s *service) GetBrandById(ctx context.Context,req *branProto.In_GetBrandById) (rsp *branProto.Out_GetBrandById,err error) {
	rsp = &branProto.Out_GetBrandById{}
	rsp.Error = &branProto.Error{}
	if req.Id <= 0 {
		rsp.Error = &branProto.Error{
			Code:    http.StatusBadRequest,
			Message: "请求参数有误！",
		}
		return
	}
	dao, err := brandDao.GetBrandDao()
	if err != nil {
		rsp.Error = &branProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}
	rsp.Brand, err = dao.FindById(req.Id)
	if err != nil {
		rsp.Error = &branProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}
	rsp.Error = &branProto.Error{
		Code:    http.StatusOK,
		Message: "查询成功！",
	}
	return

}

//修改信息
func (s *service) UpdateBrandInfo(ctx context.Context,req *branProto.In_UpdateBrandInfo) (rsp *branProto.Out_UpdateBrandInfo,err error) {
	rsp = &branProto.Out_UpdateBrandInfo{}
	updataData := map[string]interface{}{}
	elem := reflect.ValueOf(&req.Brand).Elem()
	relType := elem.Type()
	for i := 0; i < relType.NumField(); i++ {
		updataData[relType.Field(i).Name] = elem.Field(i).Interface()
	}
	dao, err := brandDao.GetBrandDao()
	if err != nil {
		rsp.Error = &branProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}
	delete(updataData, "id")
	err = dao.Update(req.Brand.Id, updataData)
	if err != nil {
		rsp.Error = &branProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}
	rsp.Error = &branProto.Error{
		Code:    http.StatusOK,
		Message: "修改成功！",
	}
	return
}

//获取列表
func (s *service) GetBrands(ctx context.Context,req *branProto.In_GetBrands) (rsp *branProto.Out_GetBrands,err error) {
	rsp = &branProto.Out_GetBrands{}
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
	dao, err := brandDao.GetBrandDao()
	if err != nil {
		log.Logf("ERROR: %v", err)
		rsp.Error = &branProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}
	rsp, err = dao.SimpleQuery(req.Limit, req.Pages, req.SearchKey, req.StartTime, req.EndTime, orderByStr)
	if err != nil {
		log.Logf("ERROR: %v", err)
		rsp.Error = &branProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}
	var message string
	if rsp.Total > 0 && len(rsp.BrandList) > 0 {
		message = "查询成功！"
	} else {
		message = "没有数据了！"
	}
	//统计有多少条
	rsp.Error = &branProto.Error{
		Code:    http.StatusOK,
		Message: message,
	}
	rsp.Limit = req.Limit
	rsp.Pages = req.Pages
	return

}

//删除列表
func (s *service) DeleteBrands(ctx context.Context,req *branProto.In_DeleteBrands) (rsp *branProto.Out_DeleteBrands,err error) {
	rsp = &branProto.Out_DeleteBrands{}
	if len(req.BrandList) <= 0 {
		rsp.Error = &branProto.Error{
			Code:    http.StatusBadRequest,
			Message: "参数不正确",
		}
		return
	}
	dao, err := brandDao.GetBrandDao()
	if err != nil {
		log.Logf("ERROR: %v", err)
		rsp.Error = &branProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	err = dao.Delete(req.BrandList)
	if err != nil {
		rsp.Error = &branProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}
	rsp.Error = &branProto.Error{
		Code:    http.StatusOK,
		Message: "删除成功!",
	}
	return
}

//新建信息
func (s *service) CreateBrand(ctx context.Context,req *branProto.In_CreateBrand) (rsp *branProto.Out_CreateBrand,err error) {
	rsp = &branProto.Out_CreateBrand{}
	//查询该等级是否存在
	dao, err := brandDao.GetBrandDao()
	if err != nil {
		log.Logf("ERROR: %v", err)
		rsp.Error = &branProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	err = dao.Insert(req.Brand)
	if err != nil {
		rsp.Error = &branProto.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		return
	}
	rsp.Error = &branProto.Error{
		Code:    http.StatusOK,
		Message: "新增成功!",
	}
	return
}
