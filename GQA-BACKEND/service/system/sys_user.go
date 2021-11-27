package system

import (
	"encoding/json"
	"errors"
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
	"gin-quasar-admin/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"sort"
)

type ServiceUser struct {
}

func (s *ServiceUser) GetUserList(requestUserList system.RequestUserList) (err error, user interface{}, total int64) {
	pageSize := requestUserList.PageSize
	offset := requestUserList.PageSize * (requestUserList.Page - 1)
	db := global.GqaDb.Model(&system.SysUser{})
	var userList []system.SysUser
	//配置搜索
	if requestUserList.Username != ""{
		db = db.Where("username like ?", "%" + requestUserList.Username + "%")
	}
	if requestUserList.RealName != ""{
		db = db.Where("real_name like ?", "%" + requestUserList.RealName + "%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(pageSize).Offset(offset).Order(global.OrderByColumn(requestUserList.SortBy, requestUserList.Desc)).Find(&userList).Error
	return err, userList, total
}

func (s *ServiceUser) EditUser(toEditUser system.SysUser) (err error) {
	var sysUser system.SysUser
	if err = global.GqaDb.Where("id = ?", toEditUser.Id).First(&sysUser).Error; err != nil {
		return err
	}
	if sysUser.Stable == "yes" {
		return errors.New("系统内置不允许编辑：" + toEditUser.Username)
	}
	err = global.GqaDb.Updates(&toEditUser).Error
	return err
}

func (s *ServiceUser) AddUser(toAddUser *system.SysUser) (err error) {
	var user system.SysUser
	if !errors.Is(global.GqaDb.Where("username = ?", toAddUser.Username).First(&user).Error, gorm.ErrRecordNotFound) {
		return errors.New("此用户已存在：" + toAddUser.Username)
	}
	defaultPassword := utils.GetConfigBackend("defaultPassword")
	if defaultPassword == ""{
		toAddUser.Password = utils.EncodeMD5("123456")
		err = global.GqaDb.Create(&toAddUser).Error
		return errors.New("未找到配置默认密码，初始密码设置为：123456")
	}else {
		toAddUser.Password = utils.EncodeMD5(defaultPassword)
		err = global.GqaDb.Create(&toAddUser).Error
		return err
	}
}

func (s *ServiceUser) DeleteUser(id uint) (err error) {
	var sysUser system.SysUser
	if err = global.GqaDb.Where("id = ?", id).First(&sysUser).Error; err != nil {
		return err
	}
	if sysUser.Stable == "yes" {
		return errors.New("系统内置不允许删除：" + sysUser.Username)
	}
	err = global.GqaDb.Where("id = ?", id).Unscoped().Delete(&sysUser).Error
	return err
}

func (s *ServiceUser) GetUserByUsername(username string) (err error, userInfo system.SysUser) {
	var user system.SysUser
	err = global.GqaDb.First(&user, "username = ?", username).Error
	return err, user
}

func (s *ServiceUser) QueryUserById(id uint) (err error, userInfo system.SysUser) {
	var user system.SysUser
	err = global.GqaDb.Preload("CreatedByUser").Preload("UpdatedByUser").First(&user, "id = ?", id).Error
	return err, user
}

func (s *ServiceUser) GetUserMenu(c *gin.Context) (err error, menu []system.SysMenu) {
	username := utils.GetUsername(c)
	var user system.SysUser
	err = global.GqaDb.Preload("Role").Where("username=?", username).First(&user).Error
	if err != nil {
		return err, nil
	}
	var role []system.SysRole
	err = global.GqaDb.Model(&user).Association("Role").Find(&role)
	if err != nil {
		return err, nil
	}
	var menus []system.SysMenu
	err = global.GqaDb.Model(&role).Association("Menu").Find(&menus)
	if err != nil {
		return err, nil
	}
	//menus切片去重
	type distinctMenu []system.SysMenu
	resultMenu := map[string]bool{}
	for _, v := range menus{
		data, _ := json.Marshal(v)
		resultMenu[string(data)] = true
	}
	result := distinctMenu{}
	for mm := range resultMenu{
		var m system.SysMenu
		_ = json.Unmarshal([]byte(mm), &m)
		result = append(result, m)
	}
	//result切片排序
	sort.Slice(result, func(i, j int) bool {
		return result[i].Sort < result[j].Sort
	})
	return nil, result
}

func (s *ServiceUser) GetUserRole(c *gin.Context) (err error, role []system.SysRole) {
	username := utils.GetUsername(c)
	var user system.SysUser
	err = global.GqaDb.Preload("Role").Where("username=?", username).First(&user).Error
	if err != nil {
		return err, nil
	}
	var userRole []system.SysRole
	err = global.GqaDb.Model(&user).Association("Role").Find(&userRole)
	if err != nil {
		return err, nil
	}
	return nil, userRole
}
