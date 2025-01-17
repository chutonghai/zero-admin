// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	cart "zero-admin/front-api/internal/handler/cart"
	category "zero-admin/front-api/internal/handler/category"
	collection "zero-admin/front-api/internal/handler/collection"
	history "zero-admin/front-api/internal/handler/history"
	home "zero-admin/front-api/internal/handler/home"
	memberaddress "zero-admin/front-api/internal/handler/member/address"
	membercoupon "zero-admin/front-api/internal/handler/member/coupon"
	membermember "zero-admin/front-api/internal/handler/member/member"
	order "zero-admin/front-api/internal/handler/order"
	product "zero-admin/front-api/internal/handler/product"
	"zero-admin/front-api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/add",
				Handler: memberaddress.MemberAddressAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/list",
				Handler: memberaddress.MemberAddressListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/update",
				Handler: memberaddress.MemberAddressUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/delete",
				Handler: memberaddress.MemberAddressDeleteHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/member/address"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: membermember.LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/register",
				Handler: membermember.RegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/info",
				Handler: membermember.InfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/updatePassword",
				Handler: membermember.UpdatePasswordHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/updateMember",
				Handler: membermember.UpdateMemberHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/member"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/add",
				Handler: cart.CartItemAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/delete",
				Handler: cart.CartItemDeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/updateQuantity",
				Handler: cart.CartUpdateQuantityHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/updateAttr",
				Handler: cart.CartUpdateAttrHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/list",
				Handler: cart.CarItemListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/list/promotion",
				Handler: cart.CarItemtListPromotionHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/itemt",
				Handler: cart.CartProductHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/clear",
				Handler: cart.CarItemClearHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/cart"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/productCateList/:parentId",
				Handler: category.ProductCateListHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/home"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/index",
				Handler: home.HomeIndexHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/recommendProductList",
				Handler: home.RecommendProductListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/recommendBrandList",
				Handler: home.RecommendBrandListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/recommendNewProductList",
				Handler: home.RecommendNewProductListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/recommendHotProductList",
				Handler: home.RecommendHotProductListHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/home"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/returnApply",
				Handler: order.ReturnApplyHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/order"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/queryProduct/:id",
				Handler: product.QueryProductHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/queryProductList",
				Handler: product.QueryProductListHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/product"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/add",
				Handler: collection.AddProductCollectionAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/delete",
				Handler: collection.ProductCollectionDeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/list",
				Handler: collection.ProductCollectionListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/clear",
				Handler: collection.ProductCollectionClearHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/collection"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/add",
				Handler: history.AddReadHistoryAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/delete",
				Handler: history.ReadHistoryDeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/list",
				Handler: history.ReadHistoryListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/clear",
				Handler: history.ReadHistoryClearHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/history"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/add",
				Handler: membercoupon.CouponAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/list",
				Handler: membercoupon.CouponHistoryListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/listByProductId",
				Handler: membercoupon.CouponListByProductIdHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/member/coupon"),
	)
}
