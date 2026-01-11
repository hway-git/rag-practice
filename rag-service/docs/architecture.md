# RAG Service 架构设计文档

## 1. 整体架构
本项目采用标准的分层架构（Layered Architecture），旨在实现高内聚、低耦合的代码结构。

### 1.1 目录结构说明
遵循 Go Standard Project Layout：

- **cmd/**: 应用程序入口。
  - `api-server/`: 主服务入口，负责启动 HTTP Server。
  - `worker/`: 异步任务入口（待实现），负责文档解析等耗时任务。
- **internal/**: 私有应用代码，不对外暴露。
  - `config/`: 配置加载逻辑。
  - `rag/`: RAG 核心业务逻辑（Retriever, Generator, Reranker）。
  - `document/`: 文档处理流程（Loader, Parser, Indexer）。
  - `vectorstore/`: 向量数据库适配层。
  - `knowledge/`: 知识库管理（CRUD）。
- **pkg/**: 公共库，可被外部引用。
  - `logger/`: 统一日志库封装。
  - `utils/`: 通用工具函数。
- **api/**: 接口定义。
  - `handler/`: HTTP 请求处理器。
  - `middleware/`: HTTP 中间件。
- **configs/**: 配置文件（YAML）。

## 2. 核心组件设计

### 2.1 配置管理 (Configuration)
- **技术选型**: [Viper](https://github.com/spf13/viper)
- **设计思路**:
  - 支持多环境配置（`dev`, `prod`）。
  - 支持环境变量覆盖（`ENV_VAR` 覆盖 `config.yaml`），便于容器化部署。
  - 强类型映射：将 YAML 配置映射为 `Config` 结构体，提供编译期类型检查。
- **关键代码**: `internal/config/config.go`

### 2.2 日志系统 (Logger)
- **技术选型**: [Zap](https://github.com/uber-go/zap)
- **设计思路**:
  - **高性能**: 使用 Zap 的结构化日志，减少内存分配。
  - **环境感知**:
    - `debug` 模式：输出彩色控制台日志，方便开发调试。
    - `release` 模式：输出 JSON 格式日志，方便 ELK 等日志系统采集。
  - **全局封装**: 在 `pkg/logger` 中封装全局 `Log` 实例，避免到处传递 Logger 对象，同时保持 API 简洁（`logger.Info(...)`）。

### 2.3 依赖注入 (待实现)
- 计划使用 Wire 或手动依赖注入，在 `main.go` 中完成各组件（Database, VectorStore, Service）的组装。

## 3. RAG 核心流程 (规划中)
1. **文档处理**: Upload -> Load -> Split -> Embed -> Vector Store
2. **检索生成**: Query -> Embed -> Vector Search -> Rerank -> Prompt -> LLM -> Response
