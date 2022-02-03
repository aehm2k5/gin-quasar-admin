package data

import (
	"fmt"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model/system"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

var SysRole = new(sysRole)

type sysRole struct{}

var sysRoleData = []system.SysRole{
	{GqaModel: global.GqaModel{Stable: "yes", Status: "on", Sort: 1001, Remark: "超级管理员角色组", CreatedAt: time.Now(), CreatedBy: "admin"},
		RoleCode: "super-admin", RoleName: "超级管理员组", DeptDataPermissionType: "all",
	},
}

func (s *sysRole) LoadData() error {
	return global.GqaDb.Transaction(func(tx *gorm.DB) error {
		var count int64
		tx.Model(&system.SysRole{}).Count(&count)
		if count != 0 {
			fmt.Println("[Gin-Quasar-Admin] --> sys_role 表的初始数据已存在，跳过初始化数据！数据量：", count)
			global.GqaLog.Warn("[Gin-Quasar-Admin] --> sys_role 表的初始数据已存在，跳过初始化数据！", zap.Any("数据量", count))
			return nil
		}
		if err := tx.Create(&sysRoleData).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		fmt.Println("[Gin-Quasar-Admin] --> sys_role 表初始数据成功！")
		global.GqaLog.Info("[Gin-Quasar-Admin] --> sys_role 表初始数据成功！")
		return nil
	})
}
