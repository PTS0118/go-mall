// Code generated by hertz generator. DO NOT EDIT.

package order

import (
	order "github.com/PTS0118/go-mall/api/biz/handler/order"
	"github.com/cloudwego/hertz/pkg/app/server"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	root.GET("/list", append(_listordersMw(), order.ListOrders)...)
	root.POST("/markPaid", append(_markorderpaidMw(), order.MarkOrderPaid)...)
	root.POST("/place", append(_placeorderMw(), order.PlaceOrder)...)
	root.POST("/update", append(_updateorderMw(), order.UpdateOrder)...)
}
