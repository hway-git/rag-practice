# rag-practice

```
rag-service/
├── api/                     # 接口层：对外提供 RAG 服务入口
│   ├── handler/             # 路由处理器
│   │   ├── chat_handler.go  # 对话接口（问答+上下文管理）
│   │   ├── doc_handler.go   # 文档上传/导入接口
│   │   └── search_handler.go # 纯检索接口（用于调试）
│   ├── middleware/          # 接口中间件（鉴权、日志、限流、请求校验）
│   ├── request/             # 入参模型
│   └── response/            # 出参模型（统一响应格式）
├── cmd/                     # 程序入口：多实例部署支持
│   ├── api-server/          # RAG 服务主程序
│   │   └── main.go          # 入口
│   └── worker/              # 异步任务 worker（文档解析、向量入库）
│       └── main.go
├── configs/                 # 配置文件：支持多环境
│   ├── config.yaml          # 主配置（向量库地址、模型地址、数据库配置）
│   ├── dev.yaml             # 开发环境配置
│   └── prod.yaml            # 生产环境配置
├── internal/                # 内部核心层：业务逻辑不对外暴露
│   ├── rag/                 # RAG 核心流程：检索增强引擎
│   │   ├── retriever/       # 检索模块
│   │   │   ├── vector_retriever.go  # 向量检索（对接 Milvus/Qdrant）
│   │   │   ├── keyword_retriever.go # 关键词检索（备用/混合检索）
│   │   │   └── hybrid_retriever.go  # 混合检索策略
│   │   ├── generator/       # 生成模块（LLM 调用）
│   │   │   ├── llm_client.go  # LLM 通用客户端（OpenAI/本地模型）
│   │   │   └── prompt_builder.go # Prompt 模板构建（上下文+问题组装）
│   │   ├── reranker/        # 重排序模块（提升检索精度）
│   │   │   └── cross_encoder_reranker.go
│   │   └── rag_service.go   # RAG 流程编排（检索→重排→生成）
│   ├── document/            # 文档处理模块
│   │   ├── loader/          # 文档加载器（多格式支持）
│   │   │   ├── pdf_loader.go
│   │   │   ├── txt_loader.go
│   │   │   └── excel_loader.go
│   │   ├── parser/          # 文档解析（文本提取、分段）
│   │   │   ├── text_splitter.go # 文本分块（按字数/语义）
│   │   │   └── metadata_extractor.go # 元数据提取（标题、作者、时间）
│   │   └── indexer/         # 文档入库（文本→向量→向量库）
│   ├── vectorstore/         # 向量数据库抽象层
│   │   ├── interface.go     # 向量库通用接口（增删改查）
│   │   ├── milvus_client.go # Milvus 实现
│   │   ├── qdrant_client.go # Qdrant 实现
│   │   └── pgvector_client.go # PGVector 实现
│   ├── knowledge/           # 知识库管理
│   │   ├── repo.go          # 知识库 CRUD（对接 MySQL/PostgreSQL）
│   │   └── manager.go       # 知识库权限、版本管理
│   └── cache/               # 缓存层（检索结果缓存、会话缓存）
│       ├── redis_client.go
│       └── cache_service.go
├── pkg/                     # 公共工具包：可对外复用
│   ├── logger/              # 日志工具（结构化日志、链路追踪）
│   ├── tracer/              # 追踪工具
│   ├── encrypt/             # 加密工具（敏感文档/配置）
│   └── utils/               # 通用工具（字符串、时间、网络）
├── scripts/                 # 脚本目录
│   ├── init_db.sql          # 数据库初始化脚本
│   ├── build.sh             # 构建脚本
│   └── deploy.sh            # 部署脚本
├── test/                    # 测试目录：单元测试+集成测试
│   ├── unit/                # 单元测试（核心组件测试）
│   └── integration/         # 集成测试（接口+流程测试）
├── go.mod                   # 依赖管理
├── Dockerfile               # 容器化构建
├── docker-compose.yaml      # 本地开发环境编排
└── README.md                # 工程说明
```