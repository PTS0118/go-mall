module github.com/PTS0118/go-mall/app/order

go 1.21.5

replace (
	github.com/PTS0118/go-mall/common => ../../common
	github.com/apache/thrift => github.com/apache/thrift v0.13.0
)

require (
	github.com/IBM/sarama v1.43.3
	github.com/golang/protobuf v1.5.4 // indirect
)
