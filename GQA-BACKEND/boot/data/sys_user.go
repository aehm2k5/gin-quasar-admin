package data

import (
	"fmt"
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

var SysUser = new(sysUser)

type sysUser struct{}

var sysUserData = []system.SysUser{
	// 超级管理员默认ID为1
	{GqaModel: global.GqaModel{Id: 1, Stable: "yes", Status: "on", Sort: 1, Remark: "我是超管我怕谁？", CreatedAt: time.Now(), CreatedBy: "admin"},
		Avatar: "", Username: "admin", Password: "e10adc3949ba59abbe56e057f20f883e",
		Nickname: "我是超管", RealName: "超级管理员",
		Gender: "u", Mobile: "1234567890", Email: "11111111111",
	},
}

func (s *sysUser) LoadData() error {
	return global.GqaDb.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&system.SysUser{}).Count(&count)
		if count != 0 {
			fmt.Println("[Gin-Quasar-Admin] --> sys_user 表的初始数据已存在，跳过初始化数据！数据量：", count)
			global.GqaLog.Error("[Gin-Quasar-Admin] --> sys_user 表的初始数据已存在，跳过初始化数据！", zap.Any("数据量", count))
			return nil
		}
		if err := tx.Create(&sysUserData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[Gin-Quasar-Admin] --> sys_user 表初始数据成功！")
		global.GqaLog.Error("[Gin-Quasar-Admin] --> sys_user 表初始数据成功！")
		return nil
	})
}
