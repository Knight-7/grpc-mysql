# grpc mysql

## 一个简单的grpc项目

使用gRPC框架调用服务端的curd

由于 Golang 1.15 版本开始废弃 CommonName，因此推荐使用 SAN 证书。 如果想兼容之前的方式，需要设置环境变量 GODEBUG 为 x509ignoreCN=0