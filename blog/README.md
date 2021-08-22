## Go 进阶训练营 毕业设计

### 题目

对当下自己项目中的业务，进行一个微服务改造，需要考虑如下技术点：
1) 微服务架构（BFF、Service、Admin、Job、Task 分模块）
2) API 设计（包括 API 定义、错误码规范、Error 的使用）
3) gRPC 的使用
4) Go 项目工程化（项目结构、DI、代码分层、ORM 框架）
5) 并发的使用（errgroup 的并行链路请求
6) 微服务中间件的使用（ELK、Opentracing、Prometheus、Kafka）
7) 缓存的使用优化（一致性处理、Pipeline 优化）

### 思路

#### 博客服务

- 博客提供三种服务, 其中 `blog` 调用 `summary` 和 `article` 拼接成完整内容
  
   - 查询博客摘要(summary)
   - 查询博客内容(article)
   - 查询完整博客(blog)
- 查询摘要时直接调用 `summary` 接口
- 查询博客时, `blog` 调用 `summary` 和 `article` 拼接成完整内容, 返回
- 查询摘要时, 先查询 Redis, 没有则到 Sqlite 查, 并将查到的内容发送到 Kafka, 进行缓存
- 摘要缓存更新任务, 消费 Kafka 中的消息, 并缓存到 Redis

### 运行

#### 基础服务
- redis
- kafka
- jaeger

#### Article 服务
```bash
cd blog/cmd/article
go build
./article.exe -conf ../../configs/article/config.yaml
```

#### Summary 服务
```bash
cd blog/cmd/summary
go build
./summary -conf ../../configs/summary/config.yaml
```

#### Blog 服务
```bash
cd blog/cmd/blog
go build
./blog -conf ../../configs/blog/config.yaml
```

#### Job 服务
```bash
cd blog/cmd/job
go build
./job -conf ../../configs/job/config.yaml
```

### 参考

- [Kratos 官方文档](https://github.com/go-kratos/kratos/blob/main/examples/blog)
- [Language Guide (proto3)](https://developers.google.com/protocol-buffers/docs/proto3)
- [`ent` Quick Start](https://entgo.io/docs/getting-started/)

### 测试

- [Postman 测试结果](./doc/postman_result.json)
