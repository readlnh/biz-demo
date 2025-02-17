# IDL 分离设计指南

## 为什么要分离 RPC 和 REST API 的 IDL？

### 1. 不同的关注点

#### RPC Service IDL
- 面向内部服务通信
- 注重服务间的高效通信
- 可以使用更细粒度的接口
- 更关注性能和类型安全
- 通常包含更多的业务细节

#### REST API IDL
- 面向外部客户端
- 注重资源的抽象和表示
- 需要更粗粒度的接口设计
- 更关注可理解性和兼容性
- 可能需要组合多个RPC调用

### 2. 目录结构建议

```
gomall/
├── idl/
│   ├── api/              # REST API 相关IDL
│   │   ├── api.proto    # API通用注解
│   │   ├── common.proto # 通用数据结构
│   │   └── v1/         # API版本
│   │       ├── product/
│   │       │   ├── api.proto    # 产品REST API定义
│   │       │   └── model.proto  # API数据模型
│   │       ├── user/
│   │       └── cart/
│   │
│   └── rpc/              # RPC Service 相关IDL
│       ├── product/
│       │   ├── service.proto # 产品服务定义
│       │   └── model.proto   # 内部数据模型
│       ├── user/
│       │   ├── service.proto
│       │   └── model.proto
│       └── cart/
           ├── service.proto
           └── model.proto
```

### 3. 示例对比

#### RPC Service IDL Example (idl/rpc/product/service.proto)
```protobuf
syntax = "proto3";

package rpc.product;

option go_package = "product/service";

import "rpc/product/model.proto";

// 内部产品服务
service ProductService {
    // 创建产品（内部服务使用）
    rpc CreateProduct(CreateProductRequest) returns (CreateProductResponse);
    
    // 批量获取产品详情
    rpc BatchGetProducts(BatchGetProductsRequest) returns (BatchGetProductsResponse);
    
    // 更新产品库存
    rpc UpdateProductStock(UpdateProductStockRequest) returns (UpdateProductStockResponse);
    
    // 内部商品搜索
    rpc SearchProducts(SearchProductsRequest) returns (SearchProductsResponse);
}

// 更细粒度的请求响应定义
message UpdateProductStockRequest {
    string product_id = 1;
    int32 stock_delta = 2;    // 库存变化量
    string operator_id = 3;   // 操作人
    string operation_type = 4; // 操作类型
    string transaction_id = 5; // 事务ID
}
```

#### REST API IDL Example (idl/api/v1/product/api.proto)
```protobuf
syntax = "proto3";

package api.v1.product;

option go_package = "api/v1/product";

import "api/v1/product/model.proto";
import "api/common.proto";

// 对外产品API
service ProductAPI {
    // 创建产品
    rpc CreateProduct(CreateProductRequest) returns (CreateProductResponse) {
        option (api.post) = "/v1/products";
    }
    
    // 获取产品详情
    rpc GetProduct(GetProductRequest) returns (GetProductResponse) {
        option (api.get) = "/v1/products/{id}";
    }
    
    // 搜索产品
    rpc SearchProducts(SearchProductsRequest) returns (SearchProductsResponse) {
        option (api.get) = "/v1/products";
    }
}

// 更高层次的抽象
message GetProductRequest {
    string id = 1 [(api.path) = "id"];
}

message GetProductResponse {
    Product product = 1;
    common.ResponseStatus status = 2;
}
```

### 4. 代码生成

#### RPC Service
```bash
# 生成RPC服务代码
cwgo server \
  -I idl \
  --type RPC \
  --service product \
  --module github.com/readlnh/biz-demo/gomall/internal/rpc \
  --idl idl/rpc/product/service.proto
```

#### REST API
```bash
# 生成REST API代码
cwgo server \
  -I idl \
  --type HTTP \
  --service api \
  --module github.com/readlnh/biz-demo/gomall/api \
  --idl idl/api/v1/*/api.proto
```

### 5. 最佳实践建议

1. **清晰的职责划分**
   - RPC IDL 专注于内部服务通信
   - REST API IDL 专注于外部接口抽象

2. **独立的版本控制**
   - RPC 服务可以根据内部需求快速迭代
   - REST API 需要考虑向后兼容性

3. **数据模型转换**
   - 在 service 层处理 RPC 模型到 API 模型的转换
   - 避免内部数据模型直接暴露给外部

4. **文档维护**
   - RPC IDL 文档面向开发团队
   - REST API 文档面向 API 消费者

通过这样的分离，可以使项目结构更清晰，职责更明确，同时也便于管理和维护。