module github.com/PTS0118/go-mall/app/product

go 1.21.5

replace github.com/apache/thrift => github.com/apache/thrift v0.13.0

require github.com/golang/protobuf v1.5.4 // indirect

require gorm.io/plugin/soft_delete v1.2.1
