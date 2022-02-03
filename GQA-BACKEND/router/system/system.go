package system

import "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/api"

type RouterSystem struct {
	RouterMenu
	RouterUser
	RouterRole
	RouterDept
	RouterDict
	RouterApi
	RouterUpload
	RouterConfigBackend
	RouterConfigFrontend
	RouterLog
	RouterNotice
	RouterTodoNote
}

var ApiSystem = api.GroupApiApp.ApiSystem
var ApiMenu = ApiSystem.ApiMenu
var ApiUser = ApiSystem.ApiUser
var ApiRole = ApiSystem.ApiRole
var ApiDept = ApiSystem.ApiDept
var ApiDict = ApiSystem.ApiDict
var ApiApi = ApiSystem.ApiApi
var ApiUpload = ApiSystem.ApiUpload
var ApiConfigBackend = ApiSystem.ApiConfigBackend
var ApiConfigFrontend = ApiSystem.ApiConfigFrontend
var ApiLogLogin = ApiSystem.ApiLogLogin
var ApiLogOperation = ApiSystem.ApiLogOperation
var ApiNotice = ApiSystem.ApiNotice
var ApiTodoNote = ApiSystem.ApiTodoNote
