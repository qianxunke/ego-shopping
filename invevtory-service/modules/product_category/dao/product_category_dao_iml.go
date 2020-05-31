package dao

import (
	productCategoryProto "github.com/qianxunke/ego-shopping/ego-common-protos/go_out/inventory/product_category"
	"github.com/qianxunke/ego-shopping/ego-plugins/db"
)

func (dao *daoIml) FindById(id int64) (product *productCategoryProto.ProductCategory, err error) {
	product = &productCategoryProto.ProductCategory{}
	DB := db.MasterEngine()
	err = DB.Where("id = ?", id).First(&product).Error
	return
}

func (dao *daoIml) Insert(product *productCategoryProto.ProductCategory) (err error) {
	DB := db.MasterEngine()
	err = DB.Create(&product).Error
	return
}

func (dao *daoIml) SimpleQuery(limit int64, pages int64, key string, startTime string, endTime string, order string) (rsp *productCategoryProto.Out_GetProductCategorys, err error) {
	DB := db.MasterEngine()
	rsp = &productCategoryProto.Out_GetProductCategorys{}
	offset := (pages - 1) * limit
	if len(key) == 0 {
		if len(startTime) > 0 && len(endTime) == 0 {
			err = DB.Model(&productCategoryProto.ProductCategory{}).Where("created_time > ?", endTime).Order(order).Count(&rsp.Total).Error
			if err == nil && rsp.Total > 0 {
				err = DB.Where("created_time > ? ", startTime).Order(order).Offset(offset).Limit(limit).Find(&rsp.ProductCategoryList).Error
			}
		} else if len(startTime) == 0 && len(endTime) > 0 {
			err = DB.Model(&productCategoryProto.ProductCategory{}).Where("created_time < ? ", endTime).Order(order).Count(&rsp.Total).Error
			if err == nil && rsp.Total > 0 {
				err = DB.Where("created_time < ? ", endTime).Order(order).Offset(offset).Limit(limit).Find(&rsp.ProductCategoryList).Error
			}
		} else if len(startTime) > 0 && len(endTime) > 0 {
			err = DB.Model(&productCategoryProto.ProductCategory{}).Where("created_time  between ? and ?", startTime, endTime).Order(order).Count(&rsp.Total).Error
			if err == nil && rsp.Total > 0 {
				err = DB.Where("created_time  between ? and ?", startTime, endTime).Order(order).Offset(offset).Limit(limit).Find(&rsp.ProductCategoryList).Error
			}
		} else {
			//先统计
			err = DB.Model(&productCategoryProto.ProductCategory{}).Order(order).Count(&rsp.Total).Error
			if err == nil && rsp.Total > 0 {
				err = DB.Order(order).Offset(offset).Limit(limit).Find(&rsp.ProductCategoryList).Error
			}
		}
	} else {
		searchKey := "%" + key + "%"
		if len(startTime) > 0 && len(endTime) == 0 {
			err = DB.Model(&productCategoryProto.ProductCategory{}).Where("(name like ? ) and created_time > ? ", searchKey, startTime).Order(order).Count(&rsp.Total).Error
			if err == nil && rsp.Total > 0 {
				err = DB.Model(&productCategoryProto.ProductCategory{}).Where("(name like ?) and created_time > ? ", searchKey, startTime).Order(order).Offset(offset).Limit(limit).Find(&rsp.ProductCategoryList).Error
			}
		} else if len(startTime) == 0 && len(endTime) > 0 {
			err = DB.Model(&productCategoryProto.ProductCategory{}).Where("(name like ?) and created_time < ? ", searchKey, endTime).Order(order).Count(&rsp.Total).Error
			if err == nil && rsp.Total > 0 {
				err = DB.Where("(name like ?) and created_time < ? ", searchKey, endTime).Order(order).Offset(offset).Limit(limit).Find(&rsp.ProductCategoryList).Error
			}
		} else if len(startTime) > 0 && len(endTime) > 0 {
			err = DB.Model(&productCategoryProto.ProductCategory{}).Where("(name like ?) and created_time between ? and ?", searchKey, startTime, endTime).Order(order).Count(&rsp.Total).Error
			if err == nil && rsp.Total > 0 {
				err = DB.Where("(name like ?) and created_time between ? and ?", searchKey, startTime, endTime).Order(order).Offset(offset).Limit(limit).Find(&rsp.ProductCategoryList).Error
			}
		} else {
			err = DB.Model(&productCategoryProto.ProductCategory{}).Where("name like ?", searchKey).Order(order).Count(&rsp.Total).Error
			if err == nil && rsp.Total > 0 {
				err = DB.Where("name like ?", searchKey).Order(order).Offset(offset).Limit(limit).Find(&rsp.ProductCategoryList).Error
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
		err = DB.Where("id = ?", ids[i]).Delete(&productCategoryProto.ProductCategory{}).Error
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
	err = DB.Model(&productCategoryProto.ProductCategory{}).Where("id = ?", id).Updates(reqMap).Error
	return
}

func (dao *daoIml) GetProductCategoryDetailsList(limit int64, pages int64, order string) (rsp *productCategoryProto.Out_GetProductCategoryDetailsList, err error) {
	DB := db.MasterEngine()
	rsp = &productCategoryProto.Out_GetProductCategoryDetailsList{}
	err = DB.Model(&productCategoryProto.ProductCategory{}).Where("parent_id = ?", 0).Order(order).Count(&rsp.Total).Error
	sql := " select c1.id id,c1.name name,c2.id  child_id, c2.name child_name from product_categories c1 left join product_categories c2 on c1.id = c2.parent_id where c1.parent_id = 0"
	if err == nil && rsp.Total > 0 {
		err = DB.Raw(sql).Scan(&rsp.ProductCategoryList).Error
	}
	return
}
