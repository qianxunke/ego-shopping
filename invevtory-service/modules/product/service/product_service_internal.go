package service

import (
	"github.com/micro/go-micro/util/log"
	//"qianxunke/es-srv/proto/product_es"
	//"qianxunke/inventory-srv/clients"
	productDao "inventory-service/modules/product/dao"
	"reflect"
	"strings"
)

func (s *service) sendAddProductMassageToEs(id string) (err error) {
	defer func() {
		if re := recover(); re != nil {
			//发送错误信息给
		}
	}()
	dao, err := productDao.GetProductDao()
	if err != nil {
		log.Logf("error: %v", err)
	}
	inProduct, err := dao.FindById(id)
	if err != nil {
		log.Logf("error : %s", err)
		return
	}
	m := make(map[string]interface{})
	elem := reflect.ValueOf(inProduct).Elem()
	relType := elem.Type()
	for i := 0; i < relType.NumField(); i++ {
		if !strings.Contains(relType.Field(i).Name, "XXX") {
			m[relType.Field(i).Name] = elem.Field(i).Interface()
		}
	}
	//将 map 转换为指定的结构体
	/*
	item := &product_es.In_CreateProduct{}
	item.Product = &product_es.Product{}
	err = mapstructure.Decode(m, item.Product)
	if err != nil {
		log.Logf("error : %s", err)
		return
	}
	err = clients.AddPusher.Publish(context.Background(), item)
	if err != nil {
		log.Logf("error : %s", err)
		return
	}

	 */
	return
}

func (s *service) sendUpdateProductMassageToEs(id string) (err error) {
	defer func() {
		if re := recover(); re != nil {
			//发送错误信息给
		}
	}()
	dao, err := productDao.GetProductDao()
	if err != nil {
		log.Logf("error: %v", err)
	}
	inProduct, err := dao.FindById(id)
	if err != nil {
		log.Logf("error : %s", err)
		return
	}
	m := make(map[string]interface{})
	elem := reflect.ValueOf(inProduct).Elem()
	relType := elem.Type()
	for i := 0; i < relType.NumField(); i++ {
		if !strings.Contains(relType.Field(i).Name, "XXX") {
			m[relType.Field(i).Name] = elem.Field(i).Interface()
		}
	}
	//将 map 转换为指定的结构体
	/*
	item := &product_es.In_UpdateProductInfo{}
	item.Product = &product_es.Product{}
	err = mapstructure.Decode(m, item.Product)
	if err != nil {
		log.Logf("error : %s", err)
		return
	}
	err = clients.UpdateProductPusher.Publish(context.Background(), item)
	if err != nil {
		log.Logf("error : %s", err)
		return
	}

	 */
	return
}

func (s *service) sendDeleteProductMassageToEs(ids []string) (err error) {
	defer func() {
		if re := recover(); re != nil {
			//发送错误信息给
		}
	}()
	/*
	msg := &product_es.In_DeleteProducts{}
	msg.ProductList = ids
	err = clients.DeleteProductPusher.Publish(context.Background(), msg)
	if err != nil {
		log.Logf("error : %s", err)
		return
	}
	 */
	return
}
