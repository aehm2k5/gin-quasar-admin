package private_service

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/service/system"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/Junvary/gqa-plugin-xk/model"
	"gorm.io/gorm"
)

func GetHonourList(getHonourList model.RequestHonourList, username string) (err error, honour []model.GqaPluginXkHonour, total int64) {
	pageSize := getHonourList.PageSize
	offset := getHonourList.PageSize * (getHonourList.Page - 1)
	var honourList []model.GqaPluginXkHonour
	var db *gorm.DB
	if err, db = system.DeptDataPermission(username, global.GqaDb.Model(&model.GqaPluginXkHonour{})); err != nil {
		return err, honourList, 0
	}
	//配置搜索
	if getHonourList.Title != ""{
		db = db.Where("title like ?", "%" + getHonourList.Title + "%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(pageSize).Offset(offset).Order(global.OrderByColumn(getHonourList.SortBy, getHonourList.Desc)).Preload("CreatedByUser").Find(&honourList).Error
	return err, honourList, total
}

func  EditHonour(toEditHonour model.GqaPluginXkHonour, username string) (err error) {
	var db *gorm.DB
	if err, db = system.DeptDataPermission(username, global.GqaDb.Model(&model.GqaPluginXkHonour{})); err != nil {
		return err
	}
	var honour model.GqaPluginXkHonour
	if err = db.Where("id = ?", toEditHonour.Id).First(&honour).Error; err != nil {
		return err
	}
	//err = db.Updates(&toEditHonour).Error
	err = db.Updates(utils.MergeMap(utils.GlobalModelToMap(&toEditHonour.GqaModel), map[string]interface{}{
		"title":      toEditHonour.Title,
		"attachment": toEditHonour.Attachment,
	})).Error
	return err
}

func AddHonour(toAddHonour model.GqaPluginXkHonour, username string) (err error) {
	var db *gorm.DB
	if err, db = system.DeptDataPermission(username, global.GqaDb.Model(&model.GqaPluginXkHonour{})); err != nil {
		return err
	}
	err = db.Create(&toAddHonour).Error
	return err
}

func DeleteHonour(id uint, username string) (err error) {
	var db *gorm.DB
	if err, db = system.DeptDataPermission(username, global.GqaDb.Model(&model.GqaPluginXkHonour{})); err != nil {
		return err
	}
	var honour model.GqaPluginXkHonour
	if err = db.Where("id = ?", id).First(&honour).Error; err != nil {
		return err
	}
	err = db.Where("id = ?", id).Unscoped().Delete(&honour).Error
	return err
}

func QueryHonourById(id uint, username string) (err error, honourInfo model.GqaPluginXkHonour) {
	var honour model.GqaPluginXkHonour
	var db *gorm.DB
	if err, db = system.DeptDataPermission(username, global.GqaDb.Model(&model.GqaPluginXkHonour{})); err != nil {
		return err, honour
	}
	err = db.Preload("CreatedByUser").Preload("UpdatedByUser").First(&honour, "id = ?", id).Error
	return err, honour
}