package data

import (
	"fmt"
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var PluginXkSysRoleMenu = new(sysRoleMenu)

type sysRoleMenu struct{}

var sysRoleMenuData = []system.SysRoleMenu{
	// 为 super-admin 设置所有 sys_menu 的总数
	{"super-admin", "GqaPluginXk"},
	{"super-admin", "plugin-xk-news"},
	{"super-admin", "plugin-xk-project"},
	{"super-admin", "plugin-xk-honour"},
	{"super-admin", "plugin-xk-document"},
	{"super-admin", "plugin-xk-download"},
}

func (s *sysRoleMenu) LoadData() error {
	return global.GqaDb.Table("sys_role_menu").Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&system.SysMenu{}).Unscoped().Where("sys_menu_name in ?",
			[]string{"GqaPluginXk", "plugin-xk-news", "plugin-xk-project"}).Count(&count)
		if count != 0 {
			fmt.Println("[GQA-Plugin] --> sys_role_menu 表中xk插件数已存在，跳过初始化数据！数据量：", count)
			global.GqaLog.Error("[GQA-Plugin] --> sys_role_menu 表中xk插件数据已存在，跳过初始化数据！", zap.Any("数据量", count))
			return nil
		}
		if err := tx.Save(&sysRoleMenuData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[GQA-Plugin] --> xk插件初始数据进入 sys_role_menu 表成功！")
		global.GqaLog.Error("[GQA-Plugin] --> xk插件初始数据进入 sys_role_menu 表成功！")
		return nil
	})
}
