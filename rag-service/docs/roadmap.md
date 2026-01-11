# 项目进度与规划 (Roadmap)

## 当前状态 (Status)
- **阶段**: 初始化 (Initialization)
- **完成度**: 10%

## 已完成特性 (Completed)
- [x] **工程结构初始化**: 遵循标准 Go 目录布局。
- [x] **配置管理模块**: 基于 Viper 实现，支持 YAML 读取和环境变量。
- [x] **日志模块**: 基于 Zap 封装，支持 Debug/Prod 模式切换。
- [x] **服务入口骨架**: `cmd/api-server` 可正常启动并加载配置。

## 待办事项 (TODO)

### 1. API 接口层 (Next Step)
- [ ] 引入 Gin Web 框架。
- [ ] 实现 HTTP Server 启动与优雅关闭。
- [ ] 实现健康检查接口 (`/health`)。
- [ ] 定义统一的 API 响应结构 (`Result`, `Error`).

### 2. 基础设施层
- [ ] **Database**: 初始化 GORM/Sqlx 连接 PostgreSQL。
- [ ] **VectorStore**: 实现 Milvus/Qdrant 客户端封装。
- [ ] **LLM Client**: 封装 OpenAI/LocalLLM 调用接口。

### 3. RAG 核心业务
- [ ] **文档处理**:
  - [ ] PDF/Text Loader
  - [ ] Text Splitter (按字符/Token)
- [ ] **检索模块**:
  - [ ] Vector Retriever
  - [ ] Keyword Retriever
- [ ] **对话接口**:
  - [ ] `/v1/chat/completions` 实现。

## 版本记录
- **v0.1.0 (Current)**: 基础骨架搭建完成。
