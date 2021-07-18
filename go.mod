module github.com/storezhang/uoa

go 1.16

require (
	github.com/cpuguy83/go-md2man/v2 v2.0.0 // indirect
	github.com/google/uuid v1.1.2 // indirect
	github.com/rs/xid v1.3.0
	github.com/storezhang/gox v1.5.3
	github.com/storezhang/pangu v1.2.8
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common v1.0.209
	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sts v1.0.209
	github.com/tencentyun/cos-go-sdk-v5 v0.7.24
	golang.org/x/crypto v0.0.0-20200820211705-5c72a883971a // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

// replace github.com/storezhang/gox => ../gox
// replace github.com/storezhang/gox => ../pangu
