## nfgo-ddd-showcase

![license](https://img.shields.io/badge/license-Apache--2.0-green.svg)
[![Language](https://img.shields.io/badge/Language-Go-blue.svg)](https://golang.org/)
[![Build Status](https://github.com/nf-go/nfgo-ddd-showcase/workflows/Go/badge.svg)](https://github.com/nf-go/nfgo-ddd-showcase/actions)


基于DDD分层架构模型的go语言的微服务代码分层模型探索和示例


1. interfaces(用户接口层): 用户接口层负责向用户显示信息和解释用户指令。这里的用户可能是：用户、其他应用程序等。

1. application(应用层): 应用层是很薄的一层，理论上不应该有业务规则或逻辑。因为领域层包含多个聚合，所以它可以协调多个聚合的服务和领域对象完成服务编排和组合，协作完成业务操作。此外，应用层也是微服务之间交互的通道，它可以调用其它微服务的应用服务，完成微服务之间的服务组合和编排。
    * `备注: 当前的nfgo-ddd-showcase省略了application层`
    * 在设计和开发时，不要将本该放在领域层的业务逻辑放到应用层中实现。因为庞大的应用层会使领域模型失焦，时间一长你的微服务就会演化为传统的三层架构，业务逻辑会变得混乱。
    * 应用服务是在应用层的，它负责服务的组合、编排和转发，负责处理业务用例的执行顺序以及结果的拼装

1. domain(领域层): 领域层的作用是实现企业核心业务逻辑，通过各种校验手段保证业务的正确性。领域层主要体现领域模型的业务能力，它用来表达业务概念、业务状态和业务规则。
    * 领域层包含聚合根、实体、值对象、领域服务等领域模型中的领域对象。
    * 领域模型的业务逻辑主要是由实体和领域服务来实现的，其中实体会采用充血模型来实现所有与之相关的业务功能。
    * 实体和领域服务在实现业务逻辑上不是同级的，当领域中的某些功能，单一实体（或者值对象）不能实现时，领域服务就会出马，它可以组合聚合内的多个实体（或者值对象），实现复杂的业务逻辑。

1. infra(基础层,infrastructure): 基础层是贯穿所有层的，它的作用就是为其它各层提供通用的技术和基础服务，包括第三方工具、驱动、消息中间件、网关、文件、缓存以及数据库等。比较常见的功能还是提供数据库持久化。


`这里的DDD分层架构示例选型松散分层架构，允许用户接口层直接与领域层交互`，另外还有严格分层架构，领域服务只能被应用服务调用，而应用服务只能被用户接口层调用，服务是逐层对外封装或组合的，依赖关系清晰。

环境准备:

* go 1.19
* 安装下面的工具:

[protoc](https://github.com/protocolbuffers/protobuf/releases/download/v3.15.6/protoc-3.15.6-linux-x86_64.zip)

```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@ad51f572fd270f2323e3aa2c1d2775cab9087af2
go install github.com/envoyproxy/protoc-gen-validate@v0.6.1
go install github.com/google/wire/cmd/wire@v0.5.0
go install github.com/swaggo/swag/cmd/swag@v1.8.4
go install golang.org/x/tools/cmd/stringer@v0.8.0
```

代码生成:

```
go mod download
go generate ./internal/...
go generate ./...
```

### 参考链接

* [DDD实战课-基于DDD的微服务拆分与设计](https://time.geekbang.org/column/intro/100037301)
* [后端开发实践系列——领域驱动设计(DDD)编码实践](https://insights.thoughtworks.cn/backend-development-ddd/)