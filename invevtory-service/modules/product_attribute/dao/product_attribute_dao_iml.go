package dao

import (
	productAttributeProto "github.com/qianxunke/ego-shopping/ego-common-protos/go_out/inventory/product_attribute"
	"github.com/qianxunke/ego-shopping/ego-plugins/db"
)

func (dao *daoIml) FindById(id int64) (product *productAttributeProto.ProductAttribute, err error) {
	product = &productAttributeProto.ProductAttribute{}
	DB := db.MasterEngine()
	err = DB.Where("id = ?", id).First(&product).Error
	return
}

func (dao *daoIml) Insert(product *productAttributeProto.ProductAttribute) (err error) {
	DB := db.MasterEngine()
	err = DB.Create(&product).Error
	return
}

func (dao *daoIml) SimpleQuery(limit int64, pages int64, key string, startTime string, endTime string, order string) (rsp *productAttributeProto.Out_GetProductAttributes, err error) {
	DB := db.MasterEngine()
	rsp = &productAttributeProto.Out_GetProductAttributes{}
	offset := (pages - 1) * limit
	if len(key) == 0 {
		if len(startTime) > 0 && len(endTime) == 0 {
			err = DB.Model(&productAttributeProto.ProductAttribute{}).Where("created_time > ?", endTime).Order(order).Count(&rsp.Total).Error
			if err == nil && rsp.Total > 0 {
				err = DB.Where("created_time > ? ", startTime).Order(order).Offset(offset).Limit(limit).Find(&rsp.ProductAttributeList).Error
			}
		} else if len(startTime) == 0 && len(endTime) > 0 {
			err = DB.Model(&productAttributeProto.ProductAttribute{}).Where("created_time < ? ", endTime).Order(order).Count(&rsp.Total).Error
			if err == nil && rsp.Total > 0 {
				err = DB.Where("created_time < ? ", endTime).Order(order).Offset(offset).Limit(limit).Find(&rsp.ProductAttributeList).Error
			}
		} else if len(startTime) > 0 && len(endTime) > 0 {
			err = DB.Model(&productAttributeProto.ProductAttribute{}).Where("created_time  between ? and ?", startTime, endTime).Order(order).Count(&rsp.Total).Error
			if err == nil && rsp.Total > 0 {
				err = DB.Where("created_time  between ? and ?", startTime, endTime).Order(order).Offset(offset).Limit(limit).Find(&rsp.ProductAttributeList).Error
			}
		} else {
			//先统计
			err = DB.Model(&productAttributeProto.ProductAttribute{}).Order(order).Count(&rsp.Total).Error
			if err == nil && rsp.Total > 0 {
				err = DB.Order(order).Offset(offset).Limit(limit).Find(&rsp.ProductAttributeList).Error
			}
		}
	} else {
		searchKey := "%" + key + "%"
		if len(startTime) > 0 && len(endTime) == 0 {
			err = DB.Model(&productAttributeProto.ProductAttribute{}).Where("(name like ? ) and created_time > ? ", searchKey, startTime).Order(order).Count(&rsp.Total).Error
			if err == nil && rsp.Total > 0 {
				err = DB.Model(&productAttributeProto.ProductAttribute{}).Where("(name like ?) and created_time > ? ", searchKey, startTime).Order(order).Offset(offset).Limit(limit).Find(&rsp.ProductAttributeList).Error
			}
		} else if len(startTime) == 0 && len(endTime) > 0 {
			err = DB.Model(&productAttributeProto.ProductAttribute{}).Where("(name like ?) and created_time < ? ", searchKey, endTime).Order(order).Count(&rsp.Total).Error
			if err == nil && rsp.Total > 0 {
				err = DB.Where("(name like ?) and created_time < ? ", searchKey, endTime).Order(order).Offset(offset).Limit(limit).Find(&rsp.ProductAttributeList).Error
			}
		} else if len(startTime) > 0 && len(endTime) > 0 {
			err = DB.Model(&productAttributeProto.ProductAttribute{}).Where("(name like ?) and created_time between ? and ?", searchKey, startTime, endTime).Order(order).Count(&rsp.Total).Error
			if err == nil && rsp.Total > 0 {
				err = DB.Where("(name like ?) and created_time between ? and ?", searchKey, startTime, endTime).Order(order).Offset(offset).Limit(limit).Find(&rsp.ProductAttributeList).Error
			}
		} else {
			err = DB.Model(&productAttributeProto.ProductAttribute{}).Where("name like ?", searchKey).Order(order).Count(&rsp.Total).Error
			if err == nil && rsp.Total > 0 {
				err = DB.Where("name like ?", searchKey).Order(order).Offset(offset).Limit(limit).Find(&rsp.ProductAttributeList).Error
			}
		}
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
		err = DB.Where("id = ?", ids[i]).Delete(&productAttributeProto.ProductAttribute{}).Error
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
	err = DB.Model(&productAttributeProto.ProductAttribute{}).Where("id = ?", id).Updates(reqMap).Error
	return
}

func (dao *daoIml) GetProductAttributeList(product_attribute_category_id int64, type_value int64) (rsp *productAttributeProto.Out_GetProductAttributeList, err error) {
	rsp = &productAttributeProto.Out_GetProductAttributeList{}
	DB := db.MasterEngine()
	sort := "sort desc"
	err = DB.Model(&productAttributeProto.ProductAttribute{}).Where("product_attribute_category_id = ? and type = ?", product_attribute_category_id, type_value).Count(&rsp.Total).Error
	if err == nil && rsp.Total > 0 {
		err = DB.Where("product_attribute_category_id = ? and type = ?", product_attribute_category_id, type_value).Order(sort).Find(&rsp.ProductAttributeList).Error
	}
	return
}
