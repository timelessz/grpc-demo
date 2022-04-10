## 包含以下内容：

go kit 框架的所有内容，包括：

- GO-KIT 中间件：service 中间件、endpoints 中间件
- swagger 文档: grpc-gate-way
- JWT认证 JWT
- 服务限流 uber/go-kit/ratelimit
- 服务降级、服务熔断 hystrix
- 服务负载均衡 (轮询、随机、虚拟节点) load balancer
- 服务发现、 服务注册 consul
- 服务链路追踪 jaeger
- 服务状态监控  (支持监控、报警、持久化) prometheus
- 日志管理 logrus
- 日志分析管理 ELK+filebeat

[grpc官网 ](https://grpc.io/docs/languages/go/quickstart/)
[grpc-gateway](https://grpc-ecosystem.github.io/grpc-gateway)

```shell
protoc --go_out=./auth --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=./auth userlogin.proto 
protoc --go_out=./auth  --go-grpc_out=./auth ./userlogin.proto 
protoc --grpc-gateway_out ./auth  --grpc-gateway_opt logtostderr=true userlogin.proto
protoc --openapiv2_out ./auth  --openapiv2_opt logtostderr=true userlogin.proto
```
