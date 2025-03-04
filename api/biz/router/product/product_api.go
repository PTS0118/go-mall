// Code generated by hertz generator. DO NOT EDIT.

package product

import (
	product "github.com/PTS0118/go-mall/api/biz/handler/product"
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
	root.POST("/create", append(_createproductMw(), product.CreateProduct)...)
	root.DELETE("/delete", append(_deleteproductMw(), product.DeleteProduct)...)
	root.POST("/list", append(_listproductsMw(), product.ListProducts)...)
	root.GET("/product", append(_getproductMw(), product.GetProduct)...)
	root.GET("/search", append(_searchproducsMw(), product.SearchProducs)...)
	root.POST("/update", append(_updateproductMw(), product.UpdateProduct)...)
}
