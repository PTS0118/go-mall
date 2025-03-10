// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	auth "github.com/PTS0118/go-mall/api/biz/router/auth"
	cart "github.com/PTS0118/go-mall/api/biz/router/cart"
	order "github.com/PTS0118/go-mall/api/biz/router/order"
	product "github.com/PTS0118/go-mall/api/biz/router/product"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	order.Register(r)

	cart.Register(r)

	product.Register(r)

	auth.Register(r)
}
