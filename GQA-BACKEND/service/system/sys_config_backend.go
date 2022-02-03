package system

import (
	"errors"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model/system"
	"gorm.io/gorm"
)

type ServiceConfigBackend struct {}

func (s *ServiceConfigBackend) GetConfigBackendList(getConfigBackendList system.RequestConfigBackendList) (err error, role interface{}, total int64) {
	pageSize := getConfigBackendList.PageSize
	offset := getConfigBackendList.PageSize * (getConfigBackendList.Page - 1)
	db := global.GqaDb.Model(&system.SysConfigBackend{})
	var configList []system.SysConfigBackend
	//配置搜索
	if getConfigBackendList.GqaOption != ""{
		db = db.Where("gqa_option like ?", "%" + getConfigBackendList.GqaOption + "%")
	}
	if getConfigBackendList.Remark != ""{
		db = db.Where("remark like ?", "%" + getConfigBackendList.Remark + "%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(pageSize).Offset(offset).Order(global.OrderByColumn(getConfigBackendList.SortBy, getConfigBackendList.Desc)).Find(&configList).Error
	return err, configList, total
}

func (s *ServiceConfigBackend) EditConfigBackend(toEditConfigBackend system.SysConfigBackend) (err error) {
	// 因为前台只传 custom 字段，这里允许编辑
	err = global.GqaDb.Save(&toEditConfigBackend).Error
	return err
}

func (s *ServiceConfigBackend) AddConfigBackend(toAddConfigBackend system.SysConfigBackend) (err error) {
	var configBackend system.SysConfigBackend
	if !errors.Is(global.GqaDb.Where("gqa_option = ?", toAddConfigBackend.GqaOption).First(&configBackend).Error, gorm.ErrRecordNotFound) {
		return errors.New("此后台配置已存在：" + toAddConfigBackend.GqaOption)
	}
	err = global.GqaDb.Create(&toAddConfigBackend).Error
	return err
}

func (s *ServiceConfigBackend) DeleteConfigBackend(id uint) (err error) {
	var sysConfigBackend system.SysConfigBackend
	if err = global.GqaDb.Where("id = ?", id).First(&sysConfigBackend).Error; err != nil {
		return err
	}
	if sysConfigBackend.Stable == "yes" {
		return errors.New("系统内置不允许删除：" + sysConfigBackend.GqaOption)
	}
	err = global.GqaDb.Where("id = ?", id).Unscoped().Delete(&sysConfigBackend).Error
	return err
}
