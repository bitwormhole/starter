# Starter Framework

Starter 是一个用 Go(Golang) 实现的依赖注入应用框架。它通过代码生成的方式工作，避免了反射对性能的影响，以及在某些情况下产生的兼容性问题。


## 安装


在安装 Starter 包之前, 需要先安装Go开发环境并准备一个基本的Go工程。

1. 首先安装 Go 开发环境 (需要 1.16 或更新的版本), 然后你就能通过下列命令来安装 Starter 了。
 
        $ go get -u github.com/bitwormhole/starter

2. 把它导入到你的代码:
 
        import "github.com/bitwormhole/starter"

<!-- 3. (可选的) Import net/http. This is required for example if using constants such as http.StatusOK.

        import "net/http" -->

## 快速开始
在工程文件夹下新建一个源文件，例如：example.go, 然后输入下列代码。

    package main

    import "github.com/bitwormhole/starter"

    func main() {
        i := starter.InitApp()
        i.Use(starter.Module())
        i.Run()
    }

运行这个例子：

    $ go run example.go


## 更多

要了解更多内容，请访问：
https://bitwormhole.com/starter/
