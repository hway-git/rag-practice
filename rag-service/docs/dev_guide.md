# 开发指南 (Development Guide)

## 1. 环境准备
- **Go**: 1.21+
- **Docker & Docker Compose**: 用于启动依赖服务（PostgreSQL, Milvus 等）。

## 2. 快速开始

### 2.1 初始化依赖
```bash
go mod tidy
```

### 2.2 启动服务
```bash
# 开发模式运行
go run cmd/api-server/main.go
```
启动成功后，将看到类似如下日志：
```
INFO    logger/logger.go:44     RAG Service Initializing...     {"version": "v0.1.0", "env": "debug"}
INFO    logger/logger.go:44     Configuration loaded successfully       {"port": 8080}
```

## 3. 开发规范

### 3.1 配置文件
- 配置文件位于 `configs/config.yaml`。
- 如果需要新增配置项：
  1. 修改 `configs/config.yaml` 添加键值对。
  2. 修改 `internal/config/config.go` 中的结构体定义，确保 `mapstructure` tag 与 yaml key 一致。

### 3.2 日志记录
- **严禁**使用 `fmt.Println` 或 `log.Println`。
- **必须**使用 `pkg/logger`。
- 示例：
  ```go
  import "rag-service/pkg/logger"
  import "go.uber.org/zap"

  // 简单日志
  logger.Info("Starting process...")

  // 结构化日志（推荐）
  logger.Info("User login", 
      zap.String("user_id", "u123"), 
      zap.Int("attempt", 1),
  )
  
  // 错误日志
  if err != nil {
      logger.Error("Database connection failed", zap.Error(err))
  }
  ```

### 3.3 错误处理
- 内部函数返回 `error`。
- HTTP Handler 层统一处理 error 并转换为 HTTP 响应。
