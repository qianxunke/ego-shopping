package dao

import (
	branProto "github.com/qianxunke/ego-shopping/ego-common-protos/go_out/inventory/brand"
	"github.com/qianxunke/ego-shopping/ego-plugins/db"
)

func (dao *productDaoIml) FindById(id int64) (product *branProto.Brand, err error) {
	product = &branProto.Brand{}
	DB := db.MasterEngine()
	err = DB.Where("id = ?", id).First(&product).Error
	return
}

func (dao *productDaoIml) Insert(product *branProto.Brand) (err error) {
	DB := db.MasterEngine()
	err = DB.Create(&product).Error
	return
}

func (dao *productDaoIml) SimpleQuery(limit int64, pages int64, key string, startTime string, endTime string, order string) (rsp *branProto.Out_GetBrands, err error) {
	DB := db.MasterEngine()
	rsp = &branProto.Out_GetBrands{}
	offset := (pages - 1) * limit
	if len(key) == 0 {
		if len(startTime) > 0 && len(endTime) == 0 {
			err = DB.Model(&branProto.Brand{}).Where("created_time > ?", endTime).Order(order).Count(&rsp.Total).Error
			if err == nil && rsp.Total > 0 {
				err = DB.Where("created_time > ? ", startTime).Order(order).Offset(offset).Limit(limit).Find(&rsp.BrandList).Error
			}
		} else if len(startTime) == 0 && len(endTime) > 0 {
			err = DB.Model(&branProto.Brand{}).Where("created_time < ? ", endTime).Order(order).Count(&rsp.Total).Error
			if err == nil && rsp.Total > 0 {
				err = DB.Where("created_time < ? ", endTime).Order(order).Offset(offset).Limit(limit).Find(&rsp.BrandList).Error
			}
		} else if len(startTime) > 0 && len(endTime) > 0 {
			err = DB.Model(&branProto.Brand{}).Where("created_time  between ? and ?", startTime, endTime).Order(order).Count(&rsp.Total).Error
			if err == nil && rsp.Total > 0 {
				err = DB.Where("created_time  between ? and ?", startTime, endTime).Order(order).Offset(offset).Limit(limit).Find(&rsp.BrandList).Error
			}
		} else {
			//先统计
			err = DB.Model(&branProto.Brand{}).Order(order).Count(&rsp.Total).Error
			if err == nil && rsp.Total > 0 {
				err = DB.Order(order).Offset(offset).Limit(limit).Find(&rsp.BrandList).Error
			}
		}
	} else {
		searchKey := "%" + key + "%"
		if len(startTime) > 0 && len(endTime) == 0 {
			err = DB.Model(&branProto.Brand{}).Where("(name like ? ) and created_time > ? ", searchKey, startTime).Order(order).Count(&rsp.Total).Error
			if err == nil && rsp.Total > 0 {
				err = DB.Model(&branProto.Brand{}).Where("(name like ?) and created_time > ? ", searchKey, startTime).Order(order).Offset(offset).Limit(limit).Find(&rsp.BrandList).Error
			}
		} else if len(startTime) == 0 && len(endTime) > 0 {
			err = DB.Model(&branProto.Brand{}).Where("(name like ?) and created_time < ? ", searchKey, endTime).Order(order).Count(&rsp.Total).Error
			if err == nil && rsp.Total > 0 {
				err = DB.Where("(name like ?) and created_time < ? ", searchKey, endTime).Order(order).Offset(offset).Limit(limit).Find(&rsp.BrandList).Error
			}
		} else if len(startTime) > 0 && len(endTime) > 0 {
			err = DB.Model(&branProto.Brand{}).Where("(name like ?) and created_time between ? and ?", searchKey, startTime, endTime).Order(order).Count(&rsp.Total).Error
			if err == nil && rsp.Total > 0 {
				err = DB.Where("(name like ?) and created_time between ? and ?", searchKey, startTime, endTime).Order(order).Offset(offset).Limit(limit).Find(&rsp.BrandList).Error
			}
		} else {
			err = DB.Model(&branProto.Brand{}).Where("name like ?", searchKey).Order(order).Count(&rsp.Total).Error
			if err == nil && rsp.Total > 0 {
				err = DB.Where("name like ?", searchKey).Order(order).Offset(offset).Limit(limit).Find(&rsp.BrandList).Error
			}
		}
	}
	return
}

func (dao *productDaoIml) Delete(ids []int64) (err error) {
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
		err = DB.Where("id = ?", ids[i]).Delete(&branProto.Brand{}).Error
		if err != nil {
			break
		}
	}
	if err != nil {
		DB.Commit()
	}
	return
}

func (dao *productDaoIml) Update(id int64, reqMap map[string]interface{}) (err error) {
	DB := db.MasterEngine()
	err = DB.Model(&branProto.Brand{}).Where("id = ?", id).Updates(reqMap).Error
	return
}
