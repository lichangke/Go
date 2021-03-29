[TOC]

更多参见：[从0开始学GO之目录](https://blog.csdn.net/leacock1991/article/details/112853343)

## goroutine

Go 语言支持并发，我们只需要通过 go 关键字来开启 goroutine 即可。

goroutine 是轻量级线程，goroutine 的调度是由 Golang 运行时进行管理的。

goroutine 语法格式：

```go
go 函数名( 参数列表 )
```

例如：

```go
go f(x, y, z)
```

开启一个新的 goroutine:

```go
f(x, y, z)
```

```go
package main
import (
        "fmt"
        "time"
)
func say(s string) {
        for i := 0; i < 5; i++ {
                time.Sleep(100 * time.Millisecond)
                fmt.Println(s)
        }
}
func main() {
        go say("world")
        say("hello")
}
```

![image-20210329172318764](Pictures\从0开始学GO之并发编程\A_从0开始学GO之并发编程.png)

主goroutine退出后，其它的工作goroutine也会自动退出：

```go
package main
import (
	"fmt"
	"time"
)
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(2 * time.Second)
		fmt.Println(s)
	}
}
func main() {
	go say("world")
	//say("hello")
	fmt.Println("main goroutine exit")
}
```

```
main goroutine exit

Process finished with exit code 0
```



## channel

通道（channel）是用来传递数据的一个数据结构。

通道可用于两个 goroutine 之间通过传递一个指定类型的值来同步运行和通讯。操作符 <- 用于指定通道的方向，发送或接收。如果未指定方向，则为双向通道。

定义一个channel时，也需要定义发送到channel的值的类型。channel可以使用内置的make()函数来创建：

```go
make(chan Type) //等价于make(chan Type, 0)
make(chan Type, capacity)
```



channel通过操作符<-来接收和发送数据，发送和接收数据语法：

```go
channel <- value      //发送value到channel
<-channel             //接收并将其丢弃
x := <-channel        //从channel中接收数据，并赋值给x
x, ok := <-channel    //功能同上，同时检查通道是否已关闭或者是否为空
```
默认情况下，channel接收和发送数据都是阻塞的，除非另一端已经准备好，这样就使得goroutine同步变的更加的简单，而不需要显式的lock。







个人能力有限，如有错误或者其他建议，敬请告知欢迎探讨，谢谢!