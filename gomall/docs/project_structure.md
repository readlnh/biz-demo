# GoMall REST API 项目结构

```
gomall/
├── api/                    # 对外暴露的API接口
│   └── v1/                # API版本
│       ├── product/       # 产品相关API
│       │   ├── product.go     # 由cwgo生成的API接口
│       │   └── types.go       # 由cwgo生成的请求/响应类型
│       ├── user/          # 用户相关API
│       └── cart/          # 购物车相关API
│
├── cmd/                    # 应用程序入口
│   └── server/
│       └── main.go        # 主程序入口
│
├── configs/               # 配置文件
│   ├── config.go         # 配置结构定义
│   ├── dev.yaml         # 开发环境配置
│   └── prod.yaml        # 生产环境配置
│
├── idl/                   # 接口定义文件
│   ├── api.proto         # API注解定义
│   ├── common.proto      # 通用数据结构
│   └── v1/              # API版本
│       ├── product/
│       │   ├── api.proto    # 产品服务API定义
│       │   └── model.proto  # 产品数据模型
│       ├── user/
│       │   ├── api.proto    # 用户服务API定义
│       │   └── model.proto  # 用户数据模型
│       └── cart/
│           ├── api.proto    # 购物车服务API定义
│           └── model.proto  # 购物车数据模型
│
├── internal/              # 私有应用程序代码
│   ├── handler/          # HTTP处理器
│   │   ├── product/     # 产品相关处理器
│   │   │   ├── create.go
│   │   │   ├── get.go
│   │   │   ├── update.go
│   │   │   ├── delete.go
│   │   │   └── list.go
│   │   ├── user/
│   │   └── cart/
│   │
│   ├── middleware/       # HTTP中间件
│   │   ├── auth.go
│   │   ├── logging.go
│   │   └── recovery.go
│   │
│   ├── service/         # 业务逻辑层
│   │   ├── product/
│   │   │   ├── service.go
│   │   │   └── product.go
│   │   ├── user/
│   │   └── cart/
│   │
│   ├── repository/      # 数据访问层
│   │   ├── mysql/
│   │   │   └── product.go
│   │   └── redis/
│   │       └── cache.go
│   │
│   └── model/          # 内部数据模型
│       ├── product.go
│       ├── user.go
│       └── cart.go
│
├── pkg/                 # 公共代码包
│   ├── auth/           # 认证相关
│   │   └── jwt.go
│   ├── config/         # 配置相关
│   │   └── loader.go
│   ├── database/       # 数据库相关
│   │   ├── mysql.go
│   │   └── redis.go
│   ├── errors/         # 错误处理
│   │   └── errors.go
│   └── logger/         # 日志相关
│       └── logger.go
│
├── scripts/            # 构建和部署脚本
│   ├── build.sh
│   └── deploy.sh
│
├── docs/              # 文档
│   ├── api.md        # API文档
│   └── schema.md     # 数据库schema文档
│
├── go.mod
├── go.sum
├── Makefile          # 项目管理命令
├── .gitignore
└── README.md
```

## 关键目录说明

1. **api/**
   - 由cwgo基于IDL文件生成的API代码
   - 包含API接口定义和类型
   - 按版本和功能模块组织

2. **cmd/**
   - 程序的入口点
   - 负责初始化和启动服务
   - 配置加载和依赖注入

3. **configs/**
   - 配置文件目录
   - 支持多环境配置
   - 配置结构定义

4. **idl/**
   - Proto文件定义
   - 按版本和功能模块组织
   - 包含API定义和数据模型

5. **internal/**
   - 私有应用程序代码
   - 遵循层次架构：handler -> service -> repository
   - 每层职责清晰分离

6. **pkg/**
   - 可以被外部项目导入的公共代码
   - 通用功能和工具
   - 基础设施代码

## 代码生成

使用cwgo生成API代码：

```bash
# 生成所有API代码
cwgo server \
  -I idl \
  --type HTTP \
  --service api \
  --module github.com/readlnh/biz-demo/gomall \
  --idl idl/v1/*/api.proto
```

## 开发工作流

1. 在 `idl/` 目录下定义API接口
2. 使用cwgo生成API代码
3. 在 `internal/handler` 实现API处理逻辑
4. 在 `internal/service` 实现业务逻辑
5. 在 `internal/repository` 实现数据访问

## 项目特点

1. **清晰的层次结构**
   - 展现层（Handler）
   - 业务层（Service）
   - 数据层（Repository）

2. **关注点分离**
   - API定义与实现分离
   - 业务逻辑与数据访问分离
   - 配置与代码分离

3. **可维护性**
   - 模块化的目录结构
   - 统一的错误处理
   - 集中的配置管理

4. **可扩展性**
   - 易于添加新的API模块
   - 便于集成新的服务
   - 支持水平扩展