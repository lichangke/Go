[TOC]

更多参见：[从0开始学GO之目录](https://blog.csdn.net/leacock1991/article/details/112853343)

## Go语言环境安装（Win）



### Go下载：

[https://golang.google.cn/dl/](https://golang.google.cn/dl/)

![image-20210119165412491](\Pictures\从0开始学GO之环境搭建\A_从0开始学GO之环境搭建.png)

### 环境变量

需要配置 3 个环境变量，也就是 GOROOT、GOPATH 和 GOBIN

- **GOROOT**：Go 语言安装根目录的路径，也就是 GO 语言的安装路径
- **GOPATH**：若干工作区目录的路径。是我们自己定义的工作空间
- **GOBIN**：GO 程序生成的可执行文件（executable file）的路径。

对于GOPATH 可以简单理解成 Go 语言的**工作目录**，它的值可以只是一个目录路径，也可以是多个目录路径，并且每个目录都代表 Go 语言的一个**工作区（workspace）**，工作区用于放置 Go 语言的**源码文件（source file）**以及安装（install）后的**归档文件（archive file，也就是以“.a”为扩展名的文件）**和**可执行文件（executable file）**。



#### GOPATH

![image-20210119171424646](\Pictures\从0开始学GO之环境搭建\E_从0开始学GO之环境搭建.png)

当然你也可以放在系统变量中，并且配置多个路径

#### GOBIN

![image-20210119171629407](\Pictures\从0开始学GO之环境搭建\F_从0开始学GO之环境搭建.png)

#### GOROOT

![image-20210119170814236](\Pictures\从0开始学GO之环境搭建\B_从0开始学GO之环境搭建.png)

我的安装根目录是`D:\Go` 

**path** 中 添加 `%GOROOT%\bin`

![image-20210119171005164](\Pictures\从0开始学GO之环境搭建\C_从0开始学GO之环境搭建.png)

通过cmd 输入<font color=red>`go env` </font> 或者 <font color=red>`go version` </font> 检查环境变量是否设置成功

![image-20210119171229471](\Pictures\从0开始学GO之环境搭建\D_从0开始学GO之环境搭建.png)



### IDE 安装（GoLand）

下载地址 [https://www.jetbrains.com/go/](https://www.jetbrains.com/go/)

请支持正版 [http://www.520xiazai.com/soft/jetbrains-2020-pojie.html](http://www.520xiazai.com/soft/jetbrains-2020-pojie.html)



### HelloWorld

```go
package main
import "fmt"
func main() {
	var helloWorld = "Hello World"
	fmt.Println(helloWorld)
}
```

![image-20210119174210704](\Pictures\从0开始学GO之环境搭建\G_从0开始学GO之环境搭建.png)



个人能力有限，如有错误或者其他建议，敬请告知欢迎探讨，谢谢!