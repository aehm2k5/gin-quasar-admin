package data

import (
	"fmt"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model/system"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var SysRoleMenu = new(sysRoleMenu)

type sysRoleMenu struct{}

var sysRoleMenuData = []system.SysRoleMenu{
	// 为 super-admin 设置所有 sys_menu 的 MenuName
	{"super-admin", "profile"},
	{"super-admin", "dashboard"},
	{"super-admin", "system"},
	{"super-admin", "dept"},
	{"super-admin", "user"},
	{"super-admin", "role"},
	{"super-admin", "menu"},
	{"super-admin", "dict"},
	{"super-admin", "config-frontend"},
	{"super-admin", "config-backend"},
	{"super-admin", "log"},
	{"super-admin", "log-login"},
	{"super-admin", "log-operation"},
	{"super-admin", "example"},
	{"super-admin", "example-icon"},
	{"super-admin", "notice"},
}

func (s *sysRoleMenu) LoadData() error {
	return global.GqaDb.Table("sys_role_menu").Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&system.SysRoleMenu{}).Count(&count)
		if count != 0 {
			fmt.Println("[Gin-Quasar-Admin] --> sys_role_menu 表的初始数据已存在，跳过初始化数据！数据量：", count)
			global.GqaLog.Warn("[Gin-Quasar-Admin] --> sys_role_menu 表的初始数据已存在，跳过初始化数据！", zap.Any("数据量", count))
			return nil
		}
		if err := tx.Create(&sysRoleMenuData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[Gin-Quasar-Admin] --> sys_role_menu 表初始数据成功!")
		global.GqaLog.Info("[Gin-Quasar-Admin] --> sys_role_menu 表初始数据成功!")
		return nil
	})
}
