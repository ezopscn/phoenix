<!--suppress HtmlDeprecatedAttribute -->
<h1 align="center">🥳 PHOENIX 后端开发文档</h1>
<h3 align="center">使用 Go + React 前后端分离开发的 Kubernetes 运维管理平台后端</h3>

<p align="center">
  <a>
    <img src="https://img.shields.io/badge/-MySQL-blue?style=flat-square&logo=mysql&logoColor=white" alt="">
  </a>
  <a>
    <img src="https://img.shields.io/badge/-Redis-c14438?style=flat-square&logo=redis&logoColor=white&link=mailto:ezops.cn@gmail.com" alt="">
  </a>
</p>

<hr>

### 🤔 技术栈

- [x] Go：Google 开发的开源编程语言，诞生于 2006 年 1 月 2 日 15 点 4 分 5 秒 [:octocat:](https://github.com/golang/go)
- [x] Cobra：CLI 开发参数处理工具 [:octocat:](https://github.com/spf13/cobra)
- [x] Embed：go 1.16 新增的文件嵌入属性, 轻松将静态文件打包到编译后的二进制应用中
- [x] Gin：用 Go 编写的高性能 HTTP Web 框架 [:octocat:](https://github.com/gin-gonic/gin)
- [x] Viper：配置管理工具, 支持多种配置文件类型 [:octocat:](https://github.com/spf13/viper)
- [x] Gorm：数据库 ORM 管理框架, 可自行扩展多种数据库类型 [:octocat:](https://gorm.io/gorm)
- [x] Carbon：简单、语义化且对开发人员友好的 datetime 包 [:octocat:](https://github.com/golang-module/carbon)
- [x] Redis：Redis 客户端 [:octocat:](https://github.com/redis/go-redis)
- [x] Jwt：用户认证, 登入登出一键搞定 [:octocat:](https://github.com/appleboy/gin-jwt)
- [x] GoMail：发送邮件 [:octocat:](https://github.com/itrepablik/gomail)

<br>

### ⚡ 开发依赖

* 项目开发所需依赖的第三方包安装方法：

```bash
# 命令行工具
go get -u github.com/spf13/cobra

# Golang web 开发框架
go get -u github.com/gin-gonic/gin

# YAML 配置文件解析成结构体
go get -u github.com/spf13/viper

# 数据库 GORM
go get -u gorm.io/gorm

# MySQL 连接驱动
go get -u gorm.io/driver/mysql

# 日期时间
go get -u github.com/golang-module/carbon/v2

# Redis 客户端
go get -u github.com/redis/go-redis/v9

# JWT 认证
go get -u github.com/appleboy/gin-jwt/v2

# 邮件发送
go get -u github.com/itrepablik/gomail
```

<br>

### 😽 项目部署

* 数据库初始化

```bash
phoenix migrate
```