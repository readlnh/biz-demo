# RESTful API 项目规划

## IDL 目录结构设计

```
gomall/
├── idl/
│   ├── api.proto      # 通用API注解定义
│   ├── common.proto   # 通用数据结构
│   ├── product/
│   │   ├── api.proto      # 产品服务的API定义
│   │   └── model.proto    # 产品相关的数据模型
│   ├── user/
│   │   ├── api.proto      # 用户服务的API定义
│   │   └── model.proto    # 用户相关的数据模型
│   └── cart/
│       ├── api.proto      # 购物车服务的API定义
│       └── model.proto    # 购物车相关的数据模型
```

## Proto 文件设计规范

### 1. API 定义规范

```protobuf
// 示例：product/api.proto
syntax = "proto3";

package api.product.v1;  // 使用版本化的包名

import "api.proto";      // 导入API注解
import "product/model.proto";  // 导入产品模型

option go_package = "product/api/v1";

// 产品服务API
service ProductService {
  // 获取产品详情
  rpc GetProduct(GetProductRequest) returns (GetProductResponse) {
    option (api.get) = "/v1/products/{id}";  // RESTful路径
  }

  // 创建产品
  rpc CreateProduct(CreateProductRequest) returns (CreateProductResponse) {
    option (api.post) = "/v1/products";
  }

  // 更新产品
  rpc UpdateProduct(UpdateProductRequest) returns (UpdateProductResponse) {
    option (api.put) = "/v1/products/{id}";
  }

  // 删除产品
  rpc DeleteProduct(DeleteProductRequest) returns (DeleteProductResponse) {
    option (api.delete) = "/v1/products/{id}";
  }

  // 搜索产品
  rpc SearchProducts(SearchProductsRequest) returns (SearchProductsResponse) {
    option (api.get) = "/v1/products";
  }
}

// 请求响应定义使用独立message，便于维护和版本控制
message GetProductRequest {
  string id = 1 [(api.path) = "id"];
}

message GetProductResponse {
  Product product = 1;
  ResponseStatus status = 2;
}
```

### 2. 数据模型规范

```protobuf
// 示例：product/model.proto
syntax = "proto3";

package api.product.v1;

import "google/protobuf/timestamp.proto";
import "common.proto";

option go_package = "product/model/v1";

// 产品模型
message Product {
  string id = 1;
  string name = 2;
  string description = 3;
  double price = 4;
  int32 stock = 5;
  repeated string categories = 6;
  google.protobuf.Timestamp created_at = 7;
  google.protobuf.Timestamp updated_at = 8;
}
```

### 3. 通用定义规范

```protobuf
// 示例：common.proto
syntax = "proto3";

package api.common;

option go_package = "api/common";

// 通用响应状态
message ResponseStatus {
  int32 code = 1;
  string message = 2;
}

// 分页请求
message PaginationRequest {
  int32 page = 1;
  int32 page_size = 2;
}

// 分页响应
message PaginationResponse {
  int32 total = 1;
  int32 total_pages = 2;
  int32 current_page = 3;
}
```

## 项目生成规划

### 1. 目录结构

```
gomall/
├── api/            # 生成的API代码
│   ├── product/
│   ├── user/
│   └── cart/
├── internal/       # 内部实现
│   ├── handler/    # API处理器
│   ├── service/    # 业务逻辑
│   ├── repository/ # 数据访问
│   └── model/      # 数据模型
└── pkg/           # 公共包
    ├── errors/    # 错误定义
    └── utils/     # 工具函数
```

### 2. 代码生成命令

```bash
# 生成API代码
cwgo server \
  -I idl \
  --type HTTP \
  --service api \
  --module github.com/readlnh/biz-demo/gomall/api \
  --idl idl/*/api.proto

# 生成模型代码
protoc \
  -I idl \
  --go_out=. \
  --go_opt=paths=source_relative \
  idl/*/model.proto
```

## 最佳实践建议

1. **版本控制**
   - 使用版本化的API路径 (如 `/v1/products`)
   - 在包名中包含版本信息

2. **API设计**
   - 使用标准的HTTP方法（GET, POST, PUT, DELETE）
   - 使用复数形式的资源名称（products而不是product）
   - 设计层次化的API路径

3. **数据模型**
   - 将数据模型与API定义分离
   - 使用通用的字段命名规范
   - 适当使用嵌套消息

4. **错误处理**
   - 使用标准的HTTP状态码
   - 提供详细的错误信息
   - 使用统一的错误响应格式

5. **文档**
   - 使用proto文件注释生成API文档
   - 提供清晰的字段描述
   - 包含示例请求和响应

## 使用说明

1. 创建新API:
   - 在对应服务目录下创建api.proto和model.proto
   - 定义服务接口和数据模型
   - 运行代码生成命令

2. 实现API:
   - 在internal/handler中实现对应的处理函数
   - 在internal/service中实现业务逻辑
   - 在internal/repository中实现数据访问

3. 测试API:
   - 使用生成的测试文件作为基础
   - 添加单元测试和集成测试
   - 使用curl或Postman测试API

4. 部署:
   - 使用Docker构建镜像
   - 配置环境变量
   - 设置API网关和负载均衡