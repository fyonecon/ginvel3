
## 了解（当前版本：gv3.^）

>介绍📖：Ginvel是一个MVC+RPC的Golang微服务框架，Http基于Gin，DDD服务基于gRPC。框架需要「 Go1.17^ + MySQL5.7^ + Redis5^ + CPU核心2^ 」的运行环境，开发和部署比较灵活。
> 
> 特点🪜：Ginvel会利用尽可能多的默认值来代替或减少报错提醒；Ginvel采用"宽进严出"的原则；Ginvel封装了大量数据处理的函数，做到开箱即用，语意化表达。
> 
> 姊妹👩：Ginvel是GinLaravel的难度版（升级版）。Ginvel使用起来可能比较难，入门难、操作难、阅读难。
>

Ginvel3框架仓库：
>（大陆🇨🇳）https://bitbucket.org/ginvel/ginvel3
> 
>（推荐🔥）https://github.com/fyonecon/ginvel3

GinLaravel框架仓库：
>https://github.com/fyonecon/ginlaravel

## 现已支持
> Go-MySQL、GORM（v2）、Go-Redis、热更（fresh）、Swagger、MVC（不推荐使用）、HTML模版输出、Http访问频率拦截、Http熔断机制、HttpCors、对称加密（可中文）、http拦截器、熔断、多层路由、分词与全文检索、运行时监控、图形验证码、生成和读取Excel、全局定时器等。

> 测试过的客户端环境等：Vue3+Axios、Fetch、POST（x-www-form-urlencoded）、GET、Centos7、Mac。

## 设计理念
> 宽进严出，功能组合式，面向Api，适合复杂项目，任何参数或服务都会有默认值。整个项目运行需go(1.17+<2021~>)+mysql+redis！

> state状态码示意：1请求成功+接口有数据，StatusCode为200；0请求成功+接口无数据，包含请求条件不足，StatusCode为200；403被拒绝访问，StatusCode为403；404接口名称错误，StatusCode为404；429请求频率限制或请求拥挤，StatusCode为429；500服务器代码错误，StatusCode为500。

> 默认端口：http默认9500，grpc默认9600

## 路由周期
> var全局参数+init()+main ➡️ 请求路由名 ➡️ header过滤 ➡️ 拦截请求频率 ➡️ 请求熔断机制 ➡️ 校验请求方法和Token参数 ➡️ 运行目标函数 ➡️ Controller程序并返回Json ➡️ 治理或运行时监控服务 ➡️ 程序结束

## 项目目录
+ /app/ ※ 你的项目目录

+ /app/request/ ※ http（api、ws）、rpc控制器目录，这是你写项目代码的地方

+ /app/arm/ ※ 项目运行时治理（辅助数据采集）

+ /bootstrap/ ※ 框架启动，驱动启动

+ /common/helper/ ※ 公共函数，助手函数

+ /common/kits/ ※ 公共封装插件

+ /docs/ ※ swagger文件目录

+ /internal/ ※ 框架基础

+ /internal/reader/ ※ 读取配置文件

+ /internal/processes/ ※ 框架运行时

+ /internal/middlewares/ ※ 框架中间件

+ /internal/providers/ ※ 框架服务提供（初始化服务）

+ /routes/xxx_group/ ※ xxx分组路由，包括api、html模板、websocket三个路由，非rpc路由

+ /storage/ ※ 系统日志、文件上传、静态缓存、toml配置文件等。目录需权限777（chmod -R 777 storage）。

+ Dockerfile ※ 执行Docker命令的文件

+ runner.conf ※ fresh自动刷新参数配置文件

+ ginvel ※ 或名为"main"，项目生产的二进制文件，在生产环境用。目录需权限773。

+ go.mod ※ 项目所依赖的module路径、第三方库等的引入

+ main.go ※ 跑起本项目的入口main文件。

## 命名原则
+ 自定义函数：大驼峰
+ 自定义变量：小驼峰
+ 自定义类：大驼峰
+ 自定义结构体和结构体成员：大驼峰
+ MySQL、Redis：小写+下划线
+ 接口名：小写+下划线
+ 系统集文件夹名或文件名：小写
+ 自定义文件夹名：大驼峰


## 基础环境
> 1⃣️Go运行环境
>
> ️搭建Go和基础Gin环境请参考：https://blog.csdn.net/weixin_41827162/article/details/115693925

> 2⃣️MySQL5.7+
>
> 请将/项目资料/ginlaravel.sql 文件导入到你的数据库

> 3⃣️Redis
>
> 请提前开启你的Redis服务

> cmd中运行「 go run main.go 」即可启动项目。
>
> 或使用热更方式启动http服务，在cmd中目录运行"fresh"。v1.4开始为适配swaggo遂将server.go更名为main.go。

> 访问
>
> http://127.0.0.1:9500

> 项目上线
>
> serverConfig["ENV"]的值改成release，然后使用以上同样方法运行。

## 如何初始化项目
以当前目录 /Users/fyonecon/go/src/ 为例
```cmd
获取源代码：
cd go/src
git clone https://github.com/fyonecon/ginvel3.git

初始化项目：
go mod init

清除不需要的vendor中的第三方扩展：
go mod tidy
    
重新载入vendor中的第三方扩展：
go mod vendor

将项目打包成二进制文件：
go build -mod=mod
（运行二进制文件需要ginvel的文件的权限为：chmod 773 ginvel）

在/stoarge/config_toml/local.framkwork.toml配置框架参数、mysql、redis、es等信息。

启动http服务：
go run main.go

或 启动二进制文件http服务：
./ginvel


```

[comment]: <> (## Go包管理工具Vendor)

[comment]: <> (```apacheconf)

[comment]: <> (下载govendor工具到本地：)

[comment]: <> (go get -u github.com/kardianos/govendor)

[comment]: <> (生成vendor文件夹（存放你项目需要的依赖包）和vendor.json：)

[comment]: <> (govendor init)

[comment]: <> (将GOPATH文件夹中的包添加到vendor目录下：)

[comment]: <> (govendor add +external)

[comment]: <> (或者govendor add +e)

[comment]: <> (或者govendor add +e)

[comment]: <> (govendor update +vendor # 更新vendor的包命令)

[comment]: <> (govendor list)

[comment]: <> (# govendor fetch github.com/xxx)

[comment]: <> (```)

## 运行fresh热更服务（Mac环境）
以项目目录 /Users/fyonecon/go/src/ginvel 为例
```cmd
安装fresh：
go get github.com/pilu/fresh

去.bash_profile文件目录：
cd ~

重新编译配置文件：
source ~/.bash_profile

切换到项目根目录：
cd go/src/ginvel

开启热更（其中项目"/storage/"目录及其子目录需要777权限。）：
fresh

退出http服务用快捷键：Ctrl + C 。或直接关闭终端窗口。

注意，在Mac开发环境中：
以上所用命令在Goland的terminal里面只运行fresh就可以了，
但是为了方便，也可在iterm2软件里面设置shell一次性运行：
iterm2——设置——Profiles——点左侧新建——General
——Cammand（选择 Login Shell）填入括号里面
【cd ~; source ~/.bash_profile; cd /Users/fyonecon/go/src/ginvel; fresh;】
内容即可。

```
以上即可项目开启的fresh热更服务。
若想一直开启终端窗口，请使用screen（yum install screen）来保持窗口。

> 热更服务文档：https://github.com/gravityblast/fresh 。

## Swagger接口文档
>官方教程：https://github.com/swaggo/gin-swagger

安装教程，在/routes/route_default.go路由文件里面，不需要再次安装和引入。）如下：
```cmd
1。进入项目跟目录

2。安装swag命令：
go get -u github.com/swaggo/swag/cmd/swag

3。查看是否命令是否安装成功：
swag -v

3.1。安装docs
swag init

3.2。切换到vendor目录，将swaggo安装在vendor目录下而不是全局安装：
cd vendor

4。安装 gin-swagger：
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files
go mod vendor

5。在route里面import引入swag：
import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "ginlaravel/docs" // docs is generated by Swag CLI, you have to import it.
)

5.1。书写路由：
url := ginSwagger.URL("http://127.0.0.1:8090/swagger/doc.json") // The url pointing to API definition
route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

5.2。启动http服务：go run main.go

5.3。访问：http://127.0.0.1:8090/swagger/index.html
```

## 将项目打包成二进制文件（Mac、Centos环境）
```cmd
#1。在ginvel项目根目录为例：
go build -mod=mod

（或 go build main.go ，main文二进制文件将是目标文件。）

#此时，在项目目录生成或更新了名字叫做"ginvel"二进制文件。其中"ginvel"是go module的名字，这是约定好的名称。
    
#2。开启文件的可执行权限：
chmod 773 ginvel
        
#3。在项目根目录运行：
./ginvel

（或 ./main）
        
#即可开启二进制服务。

# ⚠️提示：利用二进制服务开启web服务的最小代码和目录为：
/storage/ 目录及其子目录，权限777（chmod -R 777 storage）
runner.conf 文件，权限755
ginvel或main二进制文件。 文件，权限773或777。（chmod 773 ginvel或chmod 773 main）

```

## 将项目编译成exe可执行文件
```cmd
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go

```

## 将Go项目部署在Centos7上（Go环境搭建、服务器环境搭建、数据库环境搭建）
> 教程：https://blog.csdn.net/weixin_41827162/article/details/116048754

## 如何将GOPATH里面的"go get xxx"扩展引用到vendor里面
> 先运行"go get 扩展github地址"，在项目中import引入"xxx"扩展
>
> 再在项目跟目录运行"go mod vendor"
>
> 这样就可以将扩展自动引用到vendor目录下而不用govendor。

## 将go项目部署到centos7的docker中
```dockerfile

#1.1）docker images查看制作
FROM golang:latest as builder
# 修改系统为上海时区
RUN echo "Asia/Shanghai" > /etc/timezone \
 && rm /etc/localtime && dpkg-reconfigure -f noninteractive tzdata
#作者
MAINTAINER fyonecon "fyonecon@gmail.com"
#设置工作目录（暂时规定最终文件夹名与module名相同）
WORKDIR $GOPATH/src/ginvel
#将服务器的go工程代码加入到docker容器中
ADD . $GOPATH/src/ginvel
#修改env参数
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN export GIN_MODE=release
#初始化框架所需扩展
# RUN go mod vendor
#编译成二进制文件
RUN go build main.go
#RUN go build -o /tmp/main main.go
#暴露端口
EXPOSE 9500
EXPOSE 9600
#最终运行docker的命令（运行项目生成的二进制文件）
ENTRYPOINT  ["./main"]

#FROM debian:slim
#WORKDIR /go/bin
#COPY --from=builder /tmp/main ./main
#ENTRYPOINT ["/go/bin/main"]

######### Dockerfile教程 #########

#1.2）制作docker镜像，单独在项目跟目录运行（这个docker镜像名可以自定义，一般与文件夹名或module名一样即可）：
#将Dockfile文件放置在项目跟目录，然后运行
#docker build -t ginvel .
#查看镜像
#docker images
#查看正在运行的docker
#docker ps
#查看docker占用内存大小
#docker ps -as

#2）docker运行镜像，运行生成的二进制文件（二进制文件为项目go.mod里面的module名，不可更改。及其重启后，docker的端口映射需要重新运行。）
#第一个端口为宿主机端口，第二个端口为docker端口，注意httpserver的host是0.0.0.0。
#开启docker（前台方式）
#docker run -p 9500:9500 ginvel
#开启docker（后台方式，推荐）
#docker run -p 9500:9500 -d ginvel

# 如有开启rpc等必要，请开启9600端口（参考http-9500方法）

#2.1）查看已运行的docker
#docker ps

#3）查看docker日志
#docker logs -f ginvel

#4）其他（可以使用docker build的名或镜像id，用docker images查看）
#停止服务
#docker stop ginvel
#删除container实例
#docker rm ginvel
#删除container镜像
#docker rmi ginvel
#强行删除container镜像
#docker rmi -f ginvel

#5）进入docker的命令
#查看正在运行的docker的id（docker images不行）
#sudo docker ps
#进入docker
#sudo docker exec -it 正在运行的docker的id /bin/bash

```
## 单元测试testing
运行单元测试需要cd到相关文件的目录（有xxx_test.go的目录，例如cd internal/helper/），然后运行：
```cmd
go test
```
注意：本框架本身一般不需要单元测试（因为在源代码已经做了各种可能条件的判断、过滤、限制）。你可以用单元测试来测试你的自定义函数即可。

## 作者Author
> https://github.com/fyonecon
> 
> fyonecon@gmail.com

## Ginvel版权
Apache2.0 - go.ginvel.com
