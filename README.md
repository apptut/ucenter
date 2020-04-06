# GO Web 项目通用模板
golang web项目通用模板，提供日常开发的mvc以及常用的工具集

### 项目目录结构

```bash
├── ansible
├── app
│   ├── app.go
│   ├── controllers
│   ├── middleware
│   ├── models
│   │   └── user.go
│   ├── routes
│   │   └── web.go
│   ├── svc
│   └── utils
├── config
│   └── config.yaml
├── go.mod
├── go.sum
├── logs
└── main.go
```

### 使用方式
项目提供全局的 `app` 单例，如获取配置参数，数据库操作，redis操作均在该命名空间下，默认数据库和redis是全局单例模式，并且是以连接池方式运行，具体参数请修改config/下配置参数。

#### 如获取配置参数

```bash
app.Config...
```

#### Mysql

`mysql` 使用的是 `gorm` 库，具体使用请参考使用说明即可。

```bash
app.Db()...
```

#### redis
`redis` 使用的是 `redis.v5` 包，具体使用请参考使用说明即可。

```bash
app.Redis()...
```

#### 运行项目
配合脚手架工具 `gotool` 直接进入项目目录，运行：
```bash
go run main.go
```
或者指定具体的项目根路径，如果不指定则以当前目录作为项目根目录
```bash
go run main.go -prjHome=./
```

#### 项目编译

编译直接使用 `go build` 命令编译 `main.go` 文件即可：

```go
go build -o ./bin/ucenter main.go
```
