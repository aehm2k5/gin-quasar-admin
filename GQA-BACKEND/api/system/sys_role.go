package system

import (
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
	"gin-quasar-admin/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ApiRole struct{}

func (a *ApiRole) GetRoleList(c *gin.Context) {
	var requestRoleList system.RequestRoleList
	if err := c.ShouldBindJSON(&requestRoleList); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err, roleList, total := ServiceRole.GetRoleList(requestRoleList); err != nil {
		global.GqaLog.Error("获取角色列表失败！", zap.Any("err", err))
		global.ErrorMessage("获取角色列表失败，"+err.Error(), c)
	} else {
		global.SuccessData(system.ResponsePage{
			Records:  roleList,
			Page:     requestRoleList.Page,
			PageSize: requestRoleList.PageSize,
			Total:    total,
		}, c)
	}
}

func (a *ApiRole) EditRole(c *gin.Context) {
	var toEditRole system.SysRole
	if err := c.ShouldBindJSON(&toEditRole); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	toEditRole.UpdatedBy = utils.GetUsername(c)
	if err := ServiceRole.EditRole(toEditRole); err != nil {
		global.GqaLog.Error("编辑角色失败！", zap.Any("err", err))
		global.ErrorMessage("编辑角色失败，"+err.Error(), c)
	} else {
		global.SuccessMessage("编辑角色成功！", c)
	}
}

func (a *ApiRole) AddRole(c *gin.Context) {
	var toAddRole system.RequestAddRole
	if err := c.ShouldBindJSON(&toAddRole); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	addRole := &system.SysRole{
		GqaModel: global.GqaModel{
			CreatedBy: utils.GetUsername(c),
			Status:    toAddRole.Status,
			Sort:      toAddRole.Sort,
			Remark:    toAddRole.Remark,
		},
		RoleCode: toAddRole.RoleCode,
		RoleName: toAddRole.RoleName,
	}
	if err := ServiceRole.AddRole(*addRole); err != nil {
		global.GqaLog.Error("添加角色失败！", zap.Any("err", err))
		global.ErrorMessage("添加角色失败，"+err.Error(), c)
	} else {
		global.SuccessMessage("添加角色成功！", c)
	}
}

func (a *ApiRole) DeleteRole(c *gin.Context) {
	var toDeleteId system.RequestQueryById
	if err := c.ShouldBindJSON(&toDeleteId); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err := ServiceRole.DeleteRole(toDeleteId.Id); err != nil {
		global.GqaLog.Error("删除角色失败！", zap.Any("err", err))
		global.ErrorMessage("删除角色失败，"+err.Error(), c)
	} else {
		global.SuccessMessage("删除角色成功！", c)
	}
}

func (a *ApiRole) QueryRoleById(c *gin.Context) {
	var toQueryId system.RequestQueryById
	if err := c.ShouldBindJSON(&toQueryId); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err, role := ServiceRole.QueryRoleById(toQueryId.Id); err != nil {
		global.GqaLog.Error("查找角色失败！", zap.Any("err", err))
		global.ErrorMessage("查找角色失败，"+err.Error(), c)
	} else {
		global.SuccessMessageData(gin.H{"records": role}, "查找角色成功！", c)
	}
}

func (a *ApiRole) GetRoleMenuList(c *gin.Context) {
	var roleCode system.RequestRoleCode
	if err := c.ShouldBindJSON(&roleCode); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err, menuList := ServiceRole.GetRoleMenuList(&roleCode); err != nil {
		global.GqaLog.Error("获取角色菜单列表失败！", zap.Any("err", err))
		global.ErrorMessage("获取角色菜单列表失败，"+err.Error(), c)
	} else {
		global.SuccessData(gin.H{"records": menuList}, c)
	}
}

func (a *ApiRole) EditRoleMenu(c *gin.Context) {
	var roleMenu system.RequestRoleMenuEdit
	if err := c.ShouldBindJSON(&roleMenu); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if roleMenu.RoleCode == "super-admin" {
		global.ErrorMessage("超级管理员角色不允许编辑！", c)
		return
	}
	if err := ServiceRole.EditRoleMenu(&roleMenu); err != nil {
		global.GqaLog.Error("编辑角色菜单失败！", zap.Any("err", err))
		global.ErrorMessage("编辑角色菜单失败，"+err.Error(), c)
	} else {
		global.SuccessMessage("编辑角色菜单成功！", c)
	}
}

func (a *ApiRole) GetRoleApiList(c *gin.Context) {
	var roleCode system.RequestRoleCode
	if err := c.ShouldBindJSON(&roleCode); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err, apiList := ServiceRole.GetRoleApiList(&roleCode); err != nil {
		global.GqaLog.Error("获取角色API列表失败！", zap.Any("err", err))
		global.SuccessMessage("获取角色API列表失败，"+err.Error(), c)
	} else {
		global.SuccessData(gin.H{"records": apiList}, c)
	}
}

func (a *ApiRole) EditRoleApi(c *gin.Context) {
	var roleApi system.RequestRoleApiEdit
	if err := c.ShouldBindJSON(&roleApi); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if roleApi.RoleCode == "super-admin" {
		global.ErrorMessage("超级管理员角色不允许编辑！", c)
		return
	}
	if err := ServiceRole.EditRoleApi(&roleApi); err != nil {
		global.GqaLog.Error("编辑角色API失败！", zap.Any("err", err))
		global.ErrorMessage("编辑角色API失败，"+err.Error(), c)
	} else {
		global.GqaCasbin = utils.Casbin(global.GqaDb)
		global.SuccessMessage("编辑角色API成功！", c)
	}
}

func (a *ApiRole) QueryUserByRole(c *gin.Context) {
	var roleCode system.RequestRoleCode
	if err := c.ShouldBindJSON(&roleCode); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err, userList := ServiceRole.QueryUserByRole(&roleCode); err != nil {
		global.GqaLog.Error("查找角色用户失败！", zap.Any("err", err))
		global.ErrorMessage("查找角色用户失败，"+err.Error(), c)
	} else {
		global.SuccessData(gin.H{"records": userList}, c)
	}
}

func (a *ApiRole) RemoveRoleUser(c *gin.Context) {
	var toDeleteRoleUser system.RequestRoleUser
	if err := c.ShouldBindJSON(&toDeleteRoleUser); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if toDeleteRoleUser.Username == "admin" && toDeleteRoleUser.RoleCode == "super-admin" {
		global.ErrorMessage("抱歉，你不能把超级管理员从超级管理员组中移除！", c)
		return
	}
	if err := ServiceRole.RemoveRoleUser(&toDeleteRoleUser); err != nil {
		global.GqaLog.Error("移除角色用户失败！", zap.Any("err", err))
		global.ErrorMessage("移除角色用户失败，"+err.Error(), c)
	} else {
		global.SuccessMessage("移除角色用户成功！", c)
	}
}

func (a *ApiRole) AddRoleUser(c *gin.Context) {
	var toAddRoleUser system.RequestRoleUserAdd
	if err := c.ShouldBindJSON(&toAddRoleUser); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err := ServiceRole.AddRoleUser(&toAddRoleUser); err != nil {
		global.GqaLog.Error("添加角色用户失败！", zap.Any("err", err))
		global.ErrorMessage("添加角色用户失败，"+err.Error(), c)
	} else {
		global.SuccessMessage("添加角色用户成功！", c)
	}
}

func (a *ApiRole) EditRoleDeptDataPermission(c *gin.Context) {
	var toEditRoleDeptDataPermission system.RequestRoleDeptDataPermission
	if err := c.ShouldBindJSON(&toEditRoleDeptDataPermission); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err := ServiceRole.EditRoleDeptDataPermission(&toEditRoleDeptDataPermission); err != nil {
		global.GqaLog.Error("编辑角色部门数据权限失败！", zap.Any("err", err))
		global.ErrorMessage("编辑角色部门数据权限失败，"+err.Error(), c)
	} else {
		global.SuccessMessage("编辑角色部门数据权限成功！", c)
	}
}
