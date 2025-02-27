module github.com/PTS0118/go-mall/app/order

go 1.21.5

replace (
	github.com/apache/thrift => github.com/apache/thrift v0.13.0
	github.com/PTS0118/go-mall/common => ../../common
)

require github.com/golang/protobuf v1.5.4 // indirect

