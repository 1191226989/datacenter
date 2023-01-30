## datacenter

`datacenter` 是基于golang微服务框架 [go-zero](https://github.com/zeromicro/go-zero) 进行微服务模块化设计的数据中台项目，封装了常用的功能，使用简单，致力于进行快速的业务研发，同时增加了更多限制，约束项目组开发成员，规避混乱无序及自由随意的编码。

集成组件：

1. 支持 [tokenlimit](https://go-zero.dev/cn/docs/blog/governance/tokenlimit) 令牌桶限流 
1. 支持 [middleware](https://go-zero.dev/cn/docs/advance/middleware) 中间件使用 
1. 支持 [jwt](https://go-zero.dev/cn/docs/advance/jwt) 鉴权 
1. 支持 [Prometheus](https://github.com/prometheus/client_golang) 指标记录 
1. 支持 [Swagger](https://github.com/swaggo/gin-swagger) 接口文档生成 
1. 支持 trace 项目内部链路追踪 
1. 支持 [pprof](https://github.com/gin-contrib/pprof) 性能剖析
1. 支持 [errorx](https://go-zero.dev/cn/docs/advance/error-handle) 统一定义错误码 
1. 支持 [zap](https://go.uber.org/zap) 日志收集 
1. 支持 [go-redis](https://github.com/go-redis/redis/v7) 组件
1. 支持 RESTful API 返回值规范
1. 支持 [rpc编写与调用](https://go-zero.dev/cn/docs/advance/rpc-call)
1. 支持 生成数据表 [CURD](https://go-zero.dev/cn/docs/advance/model-gen)、[控制器方法](https://go-zero.dev/cn/docs/goctl/api) 等代码生成器

## 目录说明

```text
├── README.md
├── gateway ----------------------------------- 网关服务目录，实现对外访问业务api
│   ├── api ----------------------------------- 拆分网关各模块api定义
│   │   ├── game.api
│   │   ├── public.api
│   │   └── user.api
│   ├── common -------------------------------- 公共扩展方法
│   │   ├── errorx
│   │   │   └── baseerror.go
│   │   └── response
│   │       ├── basecode.go
│   │       └── response.go
│   ├── etc ----------------------------------- 网关模块配置
│   │   └── gateway-api.yaml
│   ├── gateway.api --------------------------- 网关api定义汇总
│   ├── gateway.go
│   └── internal
│       ├── config
│       │   └── config.go
│       ├── handler
│       │   ├── game
│       │   ├── public
│       │   ├── routes.go
│       │   └── user
│       ├── logic
│       │   ├── game
│       │   ├── public
│       │   └── user
│       ├── middleware ------------------------ 中间件
│       │   ├── checkpassmiddleware.go
│       │   └── tokenlimitermiddleware.go
│       ├── svc ------------------------------- 依赖配置
│       │   └── servicecontext.go
│       └── types
│           └── types.go
├── go.mod
├── go.sum
└── user -------------------------------------- 用户服务目录
    ├── model --------------------------------- 用户相关model目录
    │   ├── user.sql -------------------------- 用户表sql文件
    │   ├── usermodel.go ---------------------- 用户相关model扩展文件，可增加自定义方法
    │   ├── usermodel_gen.go
    │   └── vars.go
    └── rpc ----------------------------------- 用户rpc
        ├── etc ------------------------------- 用户rpc配置
        │   └── user.yaml
        ├── internal
        │   ├── config
        │   ├── logic
        │   ├── server
        │   └── svc
        ├── types
        │   └── user
        ├── user.go
        ├── user.proto ------------------------ proto定义文件
        └── userclient
            └── user.go
```

## 常用命令

- goctl一键安装protoc & protoc-gen-go
```
goctl env check -i -f --verbose
```

- goctl创建api模板
```
cd gateway/
goctl api -o gateway.api
```

- goctl生成api服务
```
cd gateway/
goctl api go -api gateway.api -dir .
```

- goctl创建rpc模板
```
cd rpc/
goctl rpc -o=user.proto
```

- gotcl生成rpc服务
```
cd rpc/
goctl rpc protoc user.proto --go_out=./types --go-grpc_out=./types --zrpc_out=.
```

- gotcl生成model
```
cd model/
goctl model mysql ddl -src user.sql -dir . -c
```

## 其他

- docker单机部署etcd
```
docker run -d --name etcd-server \
    --publish 2379:2379 \
    --publish 2380:2380 \
    --env ALLOW_NONE_AUTHENTICATION=yes \
    --env LISTEN-CLIENT-URLS=http://0.0.0.0:2379 \
    bitnami/etcd:latest
```

