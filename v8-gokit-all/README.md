## 包含以下内容：

go kit 框架的所有内容，包括：

- [x] 为已经实现
- [ ] 尚未实现

- -[x] GO-KIT 中间件：service 中间件、endpoints 中间件
- -[X] swagger 文档: grpc-gate-way
- -[ ] JWT认证 JWT
- -[ ] 服务限流 uber/go-kit/ratelimit
- -[ ] 服务降级、服务熔断 hystrix
- -[ ] 服务负载均衡 (轮询、随机、虚拟节点) load balancer
- -[ ] 服务发现
- -[X] 服务注册 consul
- -[ ] 服务链路追踪 jaeger
- -[ ] 服务状态监控  (支持监控、报警、持久化) prometheus grafana
- -[ ] TLS
- - [X] 日志管理 logrus
- -[ ] 日志收集、日志分析、日志持久化 ELK+filebeat ？


    


---
- [grpc官网 ](https://grpc.io/docs/languages/go/quickstart/)
- [grpc-gateway](https://grpc-ecosystem.github.io/grpc-gateway)
- [GRPC Health Checking Protocol](https://github.com/grpc/grpc/blob/master/doc/health-checking.md)
- [负载均衡算法](https://hedzr.com/golang/algorithm/go-load-balancer-1/)
- [grpc-middleware](https://github.com/grpc-ecosystem/go-grpc-middleware)
- [tls 详解](https://www.jianshu.com/p/1fc7130eb2c2)
- [从gRPC安全设计理解双向证书方案]( https://blog.csdn.net/weixin_47208161/article/details/109475931)
- [公钥、私钥、签名、数字证书的关系](https://www.jianshu.com/p/3c5212b47bec)
- [golang grpc中metadata的使用](https://blog.csdn.net/lff1123/article/details/122913719)   jaeger传递span使用
- [Interceptor拦截器 -- gRPC生态里的中间件](https://blog.csdn.net/kevin_tech/article/details/116141626)

---

auth 授权

```shell
protoc --go_out=./auth --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=./auth userlogin.proto 
protoc --grpc-gateway_out ./auth  --grpc-gateway_opt logtostderr=true userlogin.proto
protoc --openapiv2_out ./auth  --openapiv2_opt logtostderr=true userlogin.proto
```

bookservice

```shell
protoc --go_out=./book --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=./book book.proto
protoc --grpc-gateway_out ./book  --grpc-gateway_opt logtostderr=true book.proto
protoc --openapiv2_out ./book  --openapiv2_opt logtostderr=true book.proto
```

grpcui 用户界面
```shell
grpcui -bind 192.168.2.168 -plaintext 127.0.0.1:8881
```

docker consul run

```shell
docker pull consul
docker run    -d    -p 8500:8500    -p 8600:8600/udp    --name=badger   consul agent -server -ui -node=server-1 -bootstrap-expect=1 -client=0.0.0.0
curl -X PUT http://127.0.0.1:8500/v1/agent/service/deregister/book
```


生成公私钥

```shell
openssl genrsa -out ca.key 2048
openssl req -x509 -new -nodes -key ca.key -days 10000 -out ca.crt -subj "/CN=ca"

openssl ecparam -genkey -name secp384r1 -out server.key
openssl req -new -x509 -sha256 -key server.key -out server.pem -days 3650
```

### jaeger docker 运行

```jaeger
docker pull jaegertracing/all-in-one 
docker run -d -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 -p5775:5775/udp -p6831:6831/udp -p6832:6832/udp -p5778:5778 -p16686:16686 -p14268:14268 -p9411:9411 jaegertracing/all-in-one:latest
```

### 关于metadata：  
gRPC让我们可以像本地调用一样实现远程调用，对于每一次的RPC调用中，都可能会有一些在header中传递的数据，而这些数据就可以通过metadata来传递。