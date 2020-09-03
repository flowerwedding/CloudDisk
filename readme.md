# CloudDisk

[swagger接口文档](http://127.0.0.1:8000/swagger/index.html)

## 初衷

这是暑期考核的一个选项，但是我当时选了弹幕没选云盘。现在回来复盘一遍，顺便再加点我自己的理解和其他功能。

还是想去红岩网校吧，真的好想。

## 技术选型

1. web:[gin](https://github.com/gin-gonic/gin)
2. orm:[gorm](https://github.com/jinzhu/gorm)
3. database:[MySQL](https://github.com/go-sql-driver/mysql)
4. 配置文件 [go-yaml](https://github.com/go-yaml/yaml)

## 项目结构

```
-CloudDisk
    |-configs 配置文件目录
    |-docs 文档目录
    |-global 全局变量目录
    |-internal 内部模块目录
        |-dao 数据访问目录
        |-middleware HTTP中间件目录
        |-model 模型目录
        |-routers 路由相关逻辑目录
        |-service 项目核心业务逻辑目录
    |-pkg 项目相关模块包目录
    |-scripts 各类构建、安装、分析等操作的脚本目录
    |-storage 项目生成的临时文件目录
    |-third_party 第三方资源工具目录
    |-vendor 项目依赖其他开源项目目录
    |-view 模板文件目录
    |-main.go 程序执行入口
    |-setup.go
```

## TODO

- [x] 权限管理
- [x] 图形验证码
- [x] 一次性快传
- [x] 断点续传
- [x] 加密分享链接
- [x] 二维码分享链接
- [x] 下载限速
- [x] 文件列表
- [ ] 文件监控统计
- [x] 回收站

## 安装部署

本项目使用govendor管理依赖包，[govendor](https://github.com/kardianos/govendor)安装方法

```
go get -u github.com/kardianos/govendor
```

```
git clone https://github.com/flowerwedding/CloudDisk
cd CloudDisk
govendor sync
go run main.go
```

## 使用方法

### 使用说明

默认管理员username 为 hello，管理员调用对应的方法和请求可以增删改查用户权限

使用 /auth 接口登录，每次发送请求携带所生成的 token，或放在请求头中

### 注意事项

1. 如果需求上传图片功能请自行申请七牛云存储空间，并修改配置文件填写
   -  qiniu_accesskey
   -  qiniu_secretkey
   -  qiniu_fileserver 七牛访问地址
   -  qiniu_bucket 空间名称
2. 如果需要使用邮件功能，请自行填写
   - smtp_username
   - smtp_password
   - smtp_from
   - smtp_to
3. 如果需要下载文件功能，请自行修改，否则默认为本地桌面
   - file_downpath
   - file_limitrate

## 效果图

![1598585254431](https://github.com/flowerwedding/CloudDisk/blob/master/view/img/1598585254431.png)
