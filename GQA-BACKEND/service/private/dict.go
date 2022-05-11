package private

import (
	"errors"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"gorm.io/gorm"
)

type ServiceDict struct{}

func (s *ServiceDict) GetDictList(requestDictList model.RequestGetDictList) (err error, role interface{}, total int64, parentCode string) {
	pageSize := requestDictList.PageSize
	offset := requestDictList.PageSize * (requestDictList.Page - 1)
	var db *gorm.DB
	if requestDictList.ParentCode == "" {
		db = global.GqaDb.Find(&model.SysDict{})
	} else {
		db = global.GqaDb.Where("parent_code = ?", requestDictList.ParentCode).Find(&model.SysDict{})
	}
	var dictList []model.SysDict
	//配置搜索
	if requestDictList.DictCode != "" {
		db = db.Where("dict_code like ?", "%"+requestDictList.DictCode+"%")
	}
	if requestDictList.DictLabel != "" {
		db = db.Where("dict_label like ?", "%"+requestDictList.DictLabel+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(pageSize).Offset(offset).Order(model.OrderByColumn(requestDictList.SortBy, requestDictList.Desc)).Find(&dictList).Error
	return err, dictList, total, requestDictList.ParentCode
}

func (s *ServiceDict) EditDict(toEditDict model.SysDict) (err error) {
	var sysDict model.SysDict
	if err = global.GqaDb.Where("id = ?", toEditDict.Id).First(&sysDict).Error; err != nil {
		return err
	}
	if sysDict.Stable == "yes" {
		return errors.New("系统内置不允许编辑：" + toEditDict.DictCode)
	}
	//err = global.GqaDb.Updates(&toEditDict).Error
	err = global.GqaDb.Save(&toEditDict).Error
	return err
}

func (s *ServiceDict) AddDict(toAddDict model.SysDict) (err error) {
	var dict model.SysDict
	if !errors.Is(global.GqaDb.Where("dict_code = ?", toAddDict.DictCode).First(&dict).Error, gorm.ErrRecordNotFound) {
		return errors.New("此字典已存在：" + toAddDict.DictCode)
	}
	err = global.GqaDb.Create(&toAddDict).Error
	return err
}

func (s *ServiceDict) DeleteDictById(id uint) (err error) {
	var dict model.SysDict
	if err = global.GqaDb.Where("id = ?", id).First(&dict).Error; err != nil {
		return err
	}
	if dict.Stable == "yes" {
		return errors.New("系统内置不允许删除：" + dict.DictCode)
	}
	err = global.GqaDb.Where("id = ?", id).Unscoped().Delete(&dict).Error
	return err
}

func (s *ServiceDict) QueryDictById(id uint) (err error, dictInfo model.SysDict) {
	var dict model.SysDict
	err = global.GqaDb.Preload("CreatedByUser").Preload("UpdatedByUser").First(&dict, "id = ?", id).Error
	return err, dict
}
