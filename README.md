
## 1. 使用说明

- 使用git克隆本项目
   git clone https://github.com/junzhou7/pet.git

```bash
# 使用 go.mod

# 安装go依赖包
go list (go mod tidy)

# 编译
go build
```


由于国内没法安装 go.org/x 包下面的东西，推荐使用 [goproxy.io](https://goproxy.io/zh/)

bash
如果您使用的 Go 版本是 1.13 及以上(推荐)
### 2、启用 Go Modules 功能
go env -w GO111MODULE=on 
### 3、配置 GOPROXY 环境变量
go env -w GOPROXY=https://goproxy.io,direct



### 3、技术选型

- 后端：用`Gin`快速搭建基础restful风格API，`Gin`是一个go语言编写的Web框架。
- 数据库：采用`MySql`(5.6.44)版本，使用`gorm`实现对数据库的基本操作,已添加对sqlite数据库的支持。
- 缓存：使用`Redis`实现记录当前活跃用户的`jwt`令牌并实现多点登录限制。
- 配置文件：使用`fsnotify`和`viper`实现`yaml`格式的配置文件。
- 日志：使用`go-logging`实现日志记录。

### 4.1 目录结构


    ├─server  	        （后端文件夹）
    │  ├─api            （API）
    │  ├─conf           （配置包）
    │  ├─core  	        （內核）
    │  ├─db             （数据库脚本）
    │  ├─global         （全局对象）
    │  ├─initialiaze    （初始化）
    │  ├─middleware     （中间件）
    │  ├─model          （结构体层）
    │  ├─resource       （资源）
    │  ├─router         （路由）
    │  ├─service         (服务)
    │  └─utils	        （公共功能）

