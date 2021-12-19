package system

import (
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
	"gin-quasar-admin/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ApiMenu struct{}

func (a *ApiMenu) GetMenuList(c *gin.Context) {
	var requestMenuList system.RequestMenuList
	if err := c.ShouldBindJSON(&requestMenuList); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err, menuList, total := ServiceMenu.GetMenuList(requestMenuList); err != nil {
		global.GqaLog.Error("获取菜单列表失败！", zap.Any("err", err))
		global.ErrorMessage("获取菜单列表失败，"+err.Error(), c)
	} else {
		global.SuccessData(system.ResponsePage{
			Records:  menuList,
			Page:     requestMenuList.Page,
			PageSize: requestMenuList.PageSize,
			Total:    total,
		}, c)
	}
}

func (a *ApiMenu) EditMenu(c *gin.Context) {
	var toEditMenu system.SysMenu
	if err := c.ShouldBindJSON(&toEditMenu); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	toEditMenu.UpdatedBy = utils.GetUsername(c)
	if err := ServiceMenu.EditMenu(toEditMenu); err != nil {
		global.GqaLog.Error("编辑菜单失败！", zap.Any("err", err))
		global.ErrorMessage("编辑菜单失败，"+err.Error(), c)
	} else {
		global.SuccessMessage("编辑菜单成功！", c)
	}
}

func (a *ApiMenu) AddMenu(c *gin.Context) {
	var toAddMenu system.RequestAddMenu
	if err := c.ShouldBindJSON(&toAddMenu); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	addMenu := &system.SysMenu{
		GqaModel: global.GqaModel{
			CreatedBy: utils.GetUsername(c),
			Status:    toAddMenu.Status,
			Sort:      toAddMenu.Sort,
			Remark:    toAddMenu.Remark,
		},
		ParentCode: toAddMenu.ParentCode,
		Name:       toAddMenu.Name,
		Path:       toAddMenu.Path,
		Component:  toAddMenu.Component,
		Hidden:     toAddMenu.Hidden,
		KeepAlive:  toAddMenu.KeepAlive,
		Title:      toAddMenu.Title,
		Icon:       toAddMenu.Icon,
		IsLink:     toAddMenu.IsLink,
	}
	if err := ServiceMenu.AddMenu(*addMenu); err != nil {
		global.GqaLog.Error("添加菜单失败！", zap.Any("err", err))
		global.ErrorMessage("添加菜单失败，"+err.Error(), c)
	} else {
		global.SuccessMessage("添加菜单成功！", c)
	}
}

func (a *ApiMenu) DeleteMenu(c *gin.Context) {
	var toDeleteId system.RequestQueryById
	if err := c.ShouldBindJSON(&toDeleteId); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err := ServiceMenu.DeleteMenu(toDeleteId.Id); err != nil {
		global.GqaLog.Error("删除菜单失败！", zap.Any("err", err))
		global.ErrorMessage("删除菜单失败，"+err.Error(), c)
	} else {
		global.SuccessMessage("删除菜单成功！", c)
	}
}

func (a *ApiMenu) QueryMenuById(c *gin.Context) {
	var toQueryId system.RequestQueryById
	if err := c.ShouldBindJSON(&toQueryId); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err, menu := ServiceMenu.QueryMenuById(toQueryId.Id); err != nil {
		global.GqaLog.Error("查找菜单失败！", zap.Any("err", err))
		global.ErrorMessage("查找菜单失败，"+err.Error(), c)
	} else {
		global.SuccessMessageData(gin.H{"records": menu}, "查找菜单成功！", c)
	}
}
