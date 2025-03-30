# RBAC 认证系统

## 项目概述

这是一个基于角色访问控制(RBAC)的认证系统，包含后端API服务和前端管理界面。

## 技术栈

### 后端
- 框架: Gin
- ORM: Gorm
- 数据库: PostgreSQL
- 其他: Go 1.18+

### 前端
- 框架: React
- 构建工具: Rsbuild
- UI组件库: Arco Design
- 包管理: Bun

## 安装与运行

### 后端
1. 安装Go 1.18+环境
2. 配置数据库连接信息
3. 运行 `go run main.go`

### 前端
1. 安装Bun
2. 运行 `bun install` 安装依赖
3. 开发模式: `bun dev`
4. 生产构建: `bun build`

## 目录结构

```
├── api/            # 后端代码
│   ├── core/       # 核心模块
│   ├── db/         # 数据库相关
│   └── main.go     # 入口文件
├── web/            # 前端代码
│   ├── public/     # 静态资源
│   └── src/        # 源代码
└── README.md       # 项目文档
```

## 贡献指南

欢迎提交Pull Request，请遵循现有代码风格。