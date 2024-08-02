<p align="center">
 <img src="https://img.shields.io/badge/golang-1.22.3-blue" alt="Downloads">
 <img src="https://img.shields.io/badge/gin-1.10.0-blue" alt="Downloads">
 <img src="https://img.shields.io/badge/gorm-1.25.11-red" alt="Downloads"/>
 <img src="https://img.shields.io/badge/casbin-2.98.0-green" alt="Downloads"/>
 <img src="https://img.shields.io/badge/goredis-9.6.1-red" alt="Downloads"/>
 <img src="https://img.shields.io/badge/viper-1.19.0-green" alt="Downloads"/>
 <img src="https://img.shields.io/badge/zap-1.27.0-green" alt="Downloads"/>
</p>


# 项目介绍

gdmin 是一个基于 [gin](https://gin-gonic.com) 和 [gorm](https://gorm.io/docs/index.html) 开发的权限管理平台，集成了jwt鉴权，动态路由，菜单权限，按钮权限，RBAC等功能。

服务端基于golang语言开发，主要使用的技术：[gin](https://gin-gonic.com) + [gorm](https://gorm.io/docs/index.html) + [casbin](https://casbin.org/docs/overview) + [jwt](https://pkg.go.dev/github.com/golang-jwt/jwt/v5) + [go-redis](https://redis.io/docs/latest/develop/connect/clients/go/) + [viper](https://github.com/spf13/viper) + [zap](https://github.com/uber-go/zap) 等；前端采用github开源的后台管理模板 SoybeanAdmin，详见：https://github.com/soybeanjs/soybean-admin 。

# 项目截图
1. 登录页面
![登录](/_images/login.jpg)
2. 首页
![首页](/_images/home.jpg)
3. 菜单管理
![菜单管理](/_images/menu_manage.jpg)
4. 分配菜单
![分配菜单](/_images/assign-menus.jpg)

# 快速开始

### 环境准备
```angular2html
- golang 1.22+
- node.js 18.19.0+
- mysql 8.0+
- redis 7.0+
```
### 数据初始化
1. 下载项目
```
git clone https://gitee.com/nichanghao/gdmin.git
```
2. 使用docker compose启动mysql和redis
```angular2html
cd docker/
docker-compose up -d
```
3. 创建数据库
```
CREATE DATABASE gdmin;
```
4. 导入数据文件：sql/mysql/gdmin.sql

### 运行项目
1. 后端
```angular2html
cd server
go mod tidy
go run main.go
```
2. 前端
```angular2html
cd ui
pnpm i
pnpm dev
```
3. 打开浏览器访问 http://localhost:9527  
username：`gdmin`  
password：`123456`


