# seckill

本项目为软件工程第一小组课程项目。

## 开发环境配置

1. 安装go，参见[相关文档](https://github.com/Windyabc/seckill#%E7%9B%B8%E5%85%B3%E6%96%87%E6%A1%A3)
2. 设置GOROOT和GOPATH，请自行查阅相关资料
3. 安装beego
   1. `go env -w GOPROXY=https://goproxy.cn`
   2. `go get github.com/astaxie/beego`
   3. `go get github.com/beego/bee`
   4. 将$GOPATH/bin加入到系统环境变量
4. 进入$GOPATH目录
5. `mkdir seckill`
6. `cd seckill`
7. `git init`
8. `git add remote origin https://github.com/Windyabc/seckill.git`
9. `git pull remote main`
10. `bee run`
11. 若能正常访问`localhost:8080`，则环境配置成功

## 协作流程及规范

1. 若无相关issue，则新建issue，对要实现的功能进行简单描述。前后端接口也应在此明确。
2. 本地建立新分支，在新分支上进行编码、本地测试
3. commit
4. push到新分支，新分支以`姓名/功能概括`命名
5. 提交pr
6. 组长/产品验收后merge，若发现问题，则提醒相关开发者进行修改

## 相关文档

* golang

  [go的安装（官方文档）](https://golang.google.cn/doc/install)

* beego

  [beego官方中文文档](https://beego.me/docs/quickstart/new.md)

  
