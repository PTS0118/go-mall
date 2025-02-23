package rpc

import (
	"github.com/PTS0118/go-mall/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/PTS0118/go-mall/rpc_gen/kitex_gen/order/orderservice"
	"github.com/PTS0118/go-mall/rpc_gen/kitex_gen/product/productcatalogservice"

	//"context"
	//"github.com/cloudwego/biz-demo/gomall/common/mtl"
	"log"
	"sync"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/discovery"
	consul "github.com/kitex-contrib/registry-consul"

	"github.com/PTS0118/go-mall/api/conf"

	//"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/checkout/checkoutservice"
	//"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/order/orderservice"
	//"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/product"
	//"github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/PTS0118/go-mall/rpc_gen/kitex_gen/user/userservice"
	//"github.com/cloudwego/kitex/pkg/circuitbreak"
	//"github.com/cloudwego/kitex/pkg/fallback"
	//"github.com/cloudwego/kitex/pkg/rpcinfo"
	//prometheus "github.com/kitex-contrib/monitor-prometheus"
)

var (
	ProductClient productcatalogservice.Client
	UserClient    userservice.Client
	CartClient    cartservice.Client
	//CheckoutClient checkoutservice.Client
	OrderClient  orderservice.Client
	once         sync.Once
	err          error
	registryAddr string
	//commonSuite  client.Option
	resolver discovery.Resolver
)

func InitClient() {
	once.Do(func() {
		registryAddr = conf.GetConf().Hertz.RegistryAddr
		resolver, err = consul.NewConsulResolver("127.0.0.1:8500")
		if err != nil {
			log.Fatal(err)
		}
		//commonSuite = client.WithSuite(clientsuite.CommonGrpcClientSuite{
		//	RegistryAddr:       registryAddr,
		//	CurrentServiceName: "api",
		//})
		initProductClient()
		initUserClient()
		//initCartClient()
		//initCheckoutClient()
		initOrderClient()
	})
}

func initProductClient() {
	ProductClient, err = productcatalogservice.NewClient("product", client.WithResolver(resolver), client.WithRPCTimeout(time.Second*3))
	if err != nil {
		hlog.Fatal(err)
	}
	//var opts []client.Option
	//
	//cbs := circuitbreak.NewCBSuite(func(ri rpcinfo.RPCInfo) string {
	//	return circuitbreak.RPCInfo2Key(ri)
	//})
	//cbs.UpdateServiceCBConfig("shop-frontend/product/GetProduct", circuitbreak.CBConfig{Enable: true, ErrRate: 0.5, MinSample: 2})
	//
	//opts = append(opts, commonSuite, client.WithCircuitBreaker(cbs), client.WithFallback(fallback.NewFallbackPolicy(fallback.UnwrapHelper(func(ctx context.Context, req, resp interface{}, err error) (fbResp interface{}, fbErr error) {
	//	methodName := rpcinfo.GetRPCInfo(ctx).To().Method()
	//	if err == nil {
	//		return resp, err
	//	}
	//	if methodName != "ListProducts" {
	//		return resp, err
	//	}
	//	return &product.ListProductsResp{
	//		Products: []*product.Product{
	//			{
	//				Price:       6.6,
	//				Id:          3,
	//				Picture:     "/static/image/t-shirt.jpeg",
	//				Name:        "T-Shirt",
	//				Description: "CloudWeGo T-Shirt",
	//			},
	//		},
	//	}, nil
	//}))))
	//opts = append(opts, client.WithTracer(prometheus.NewClientTracer("", "", prometheus.WithDisableServer(true), prometheus.WithRegistry(mtl.Registry))))
	//
	//ProductClient, err = productcatalogservice.NewClient("product", opts...)
	//frontendutils.MustHandleError(err)
}

func initUserClient() {
	UserClient, err = userservice.NewClient("user", client.WithResolver(resolver), client.WithRPCTimeout(time.Second*3))
	if err != nil {
		hlog.Fatal(err)
	}
}

//	func initCartClient() {
//		CartClient, err = cartservice.NewClient("cart", commonSuite)
//		frontendutils.MustHandleError(err)
//	}
//
//	func initCheckoutClient() {
//		CheckoutClient, err = checkoutservice.NewClient("checkout", commonSuite)
//		frontendutils.MustHandleError(err)
//	}
func initOrderClient() {
	OrderClient, err = orderservice.NewClient("order", client.WithResolver(resolver), client.WithRPCTimeout(time.Second*3))
	if err != nil {
		hlog.Fatal(err)
	}
}
