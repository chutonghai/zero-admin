// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	config "go-zero-admin/api/internal/handler/config"
	dept "go-zero-admin/api/internal/handler/dept"
	dict "go-zero-admin/api/internal/handler/dict"
	log "go-zero-admin/api/internal/handler/log"
	menu "go-zero-admin/api/internal/handler/menu"
	role "go-zero-admin/api/internal/handler/role"
	user "go-zero-admin/api/internal/handler/user"
	"go-zero-admin/api/internal/svc"

	"github.com/tal-tech/go-zero/rest"
)

func RegisterHandlers(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/user/login",
				Handler: user.UserLoginHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/user/currentUser",
				Handler: user.UserInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/user/add",
				Handler: user.UserAddHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/user/list",
				Handler: user.UserListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/user/update",
				Handler: user.UserUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/user/delete",
				Handler: user.UserDeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/user/updateUserRole",
				Handler: user.UpdateUserRoleHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/user/reSetPassword",
				Handler: user.ReSetPasswordHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/user/UpdateUserStatus",
				Handler: user.UpdateUserStatusHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/role/add",
				Handler: role.RoleAddHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/role/list",
				Handler: role.RoleListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/role/update",
				Handler: role.RoleUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/role/delete",
				Handler: role.RoleDeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/role/roleMenuIds",
				Handler: role.RoleMenuIdsHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/role/updateRoleMenu",
				Handler: role.UpdateRoleMenuHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/menu/add",
				Handler: menu.MenuAddHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/menu/list",
				Handler: menu.MenuListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/menu/update",
				Handler: menu.MenuUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/menu/delete",
				Handler: menu.MenuDeleteHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/dict/add",
				Handler: dict.DictAddHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/dict/list",
				Handler: dict.DictListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/dict/update",
				Handler: dict.DictUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/dict/delete",
				Handler: dict.DictDeleteHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/dept/add",
				Handler: dept.DeptAddHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/dept/list",
				Handler: dept.DeptListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/dept/update",
				Handler: dept.DeptUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/dept/delete",
				Handler: dept.DeptDeleteHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/api/loginLog/list",
				Handler: log.LoginLogListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/loginLog/delete",
				Handler: log.LoginLogDeleteHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/api/sysLog/list",
				Handler: log.SysLogListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/sysLog/delete",
				Handler: log.SysLogDeleteHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/config/add",
				Handler: config.ConfigAddHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/config/list",
				Handler: config.ConfigListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/config/update",
				Handler: config.ConfigUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/config/delete",
				Handler: config.ConfigDeleteHandler(serverCtx),
			},
		},
	)
}
