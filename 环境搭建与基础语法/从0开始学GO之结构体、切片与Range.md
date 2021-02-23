[TOC]

更多参见：[从0开始学GO之目录](https://blog.csdn.net/leacock1991/article/details/112853343)

## 结构体



Go 语言中数组可以存储同一类型的数据，但在结构体中可以为不同项定义不同的数据类型。

结构体是一种复合的数据类型，它是由一系列具有相同类型或不同类型的数据构成的数据集合。每个数据称为结构体的成员。



### 定义结构体

结构体定义需要使用 type 和 struct 语句。struct 语句定义一个新的数据类型，结构体中有一个或多个成员。type 语句设定了结构体的名称。结构体的格式如下：

```go
type struct_variable_type struct {
   member definition
   member definition
   ...
   member definition
}
```

### 结构体初始化

普通变量与指针变量

```go
package main
import "fmt"

type Books struct {
	title string
	author string
	subject string
	book_id int
}

func main() {
	// 普通变量
	Book1 := Books{"Go 语言", "www.runoob.com", "Go 语言教程", 12345} // 也可以如下
	// var Book1 Books        // 普通变量
	//Book1.title = "Go 语言"
	//Book1.author = "www.runoob.com"
	//Book1.subject = "Go 语言教程"
	//Book1.book_id = 6495407
	// var Book1 Books  = Books{"Go 语言", "www.runoob.com", "Go 语言教程", 12345}
	// 指针变量
	Book2 :=  &Books{"Python 教程", "www.runoob.com", "Python 语言教程", 54321}      // 也可以如下
	// var Book2 *Books =  &Books{"Python 教程", "www.runoob.com", "Python 语言教程", 54321}
	fmt.Println(Book1)
	fmt.Println(Book2)

}
```

![image-20210127220736573](\Pictures\从0开始学GO之结构体、切片与Range\A_从0开始学GO之结构体、切片与Range.png)



### 结构体作为函数参数

值传递与引用传递

```go
package main
import "fmt"

type Books struct {
	title string
	author string
	subject string
	book_id int
}

func printBookAndChange1( book Books ) {
	fmt.Printf( "Book title : %s\n", book.title)
	fmt.Printf( "Book author : %s\n", book.author)
	fmt.Printf( "Book subject : %s\n", book.subject)
	fmt.Printf( "Book book_id : %d\n", book.book_id)
	book.title = book.title + "_Add"
	book.book_id = -12345
}

func printBookAndChange2( book *Books ) {
	fmt.Printf( "Book title : %s\n", book.title)
	fmt.Printf( "Book author : %s\n", book.author)
	fmt.Printf( "Book subject : %s\n", book.subject)
	fmt.Printf( "Book book_id : %d\n", book.book_id)
	book.title = book.title + "_Add"
	book.book_id = -12345
}

func main() {
	Book1 := Books{"Go 语言", "www.runoob.com", "Go 语言教程", 12345}
	Book2 :=  &Books{"Python 教程", "www.runoob.com", "Python 语言教程", 54321}
	fmt.Println("======================")
	fmt.Println(Book1)
	fmt.Println(Book2)
	fmt.Println("======================")
	printBookAndChange1(Book1)
	fmt.Println("======================")
	printBookAndChange2(Book2)
	fmt.Println("======================")
	fmt.Println(Book1)
	fmt.Println(Book2)
	fmt.Println("======================")
	//p.成员 和(*p).成员 操作是等价的
	Book2.book_id = -1
	(*Book2).title = "Python"
	fmt.Println(Book2)
}
```



![image-20210127221802756](Pictures\从0开始学GO之结构体、切片与Range\B_从0开始学GO之结构体、切片与Range.png)





## 切片

Go 语言切片是对数组的抽象。数组的长度在定义之后无法再次修改；数组是值类型，每次传递都将产生一份副本。Go 中提供了一种灵活，功能强悍的内置类型切片("动态数组")，与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。

**它通过内部指针和相关属性引⽤数组⽚段，以实现变⻓⽅案**。slice并不是真正意义上的动态数组，而是一个引用类型。slice总是指向一个底层array，slice的声明也可以像array一样，只是不需要长度。

![img](Pictures\从0开始学GO之结构体、切片与Range\C_从0开始学GO之结构体、切片与Range.png)

图片参见 [https://www.jianshu.com/p/659da9efd55f](https://www.jianshu.com/p/659da9efd55f)

### 定义切片

切片不需要说明长度，可以声明一个未指定大小的数组来定义切片：

```go
var identifier []type
```

或使用 **make()** 函数来创建切片:

```go
var slice1 []type = make([]type, len)
也可以简写为
slice1 := make([]type, len)
```

也可以指定容量，其中 **capacity** 为可选参数。

```go
make([]T, length, capacity)
```

### 切片初始化

```go
s :=[] int {1,2,3 } 
直接初始化切片，[] 表示是切片类型，{1,2,3} 初始化值依次是 1,2,3，其 cap=len=3。
s := arr[:] 
初始化切片 s，是数组 arr 的引用。
s := arr[startIndex:endIndex] 
将 arr 中从下标 startIndex 到 endIndex-1 下的元素创建为一个新的切片。
s := arr[startIndex:] 
默认 endIndex 时将表示一直到arr的最后一个元素。
s := arr[:endIndex] 
默认 startIndex 时将表示从 arr 的第一个元素开始。
s1 := s[startIndex:endIndex] 
通过切片 s 初始化切片 s1。
s :=make([]int,len,cap) 
通过内置函数 make() 初始化切片s，[]int 标识为其元素类型为 int 的切片。
```

### 切片的操作

| 操作            | 含义                                                         |
| --------------- | ------------------------------------------------------------ |
| s[n]            | 切片s中索引位置为n的项                                       |
| s[:]            | 从切片s的索引位置0到len(s)-1处所获得的切片                   |
| s[low:]         | 从切片s的索引位置low到len(s)-1处所获得的切片                 |
| s[:high]        | 从切片s的索引位置0到high处所获得的切片，len=high             |
| s[low:high]     | 从切片s的索引位置low到high处所获得的切片，len=high-low       |
| s[low:high:max] | 从切片s的索引位置low到high处所获得的切片，len=high-low，cap=max-low |
| len(s)          | 切片s的长度，总是<=casp(s)                                   |
| cap(s)          | 切片s的容量，总是>=len(s)                                    |

示例说明：

```
    array := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
```

| 操作        | 结果                  | len  | cap  | 说明            |
| ----------- | --------------------- | ---- | ---- | --------------- |
| array[:6:8] | [0 1 2 3 4 5]         | 6    | 8    | 省略 low        |
| array[5:]   | [5 6 7 8 9]           | 5    | 5    | 省略 high、 max |
| array[:3]   | [0 1 2]               | 3    | 10   | 省略 high、 max |
| array[:]    | [0 1 2 3 4 5 6 7 8 9] | 10   | 10   | 全部省略        |

###  切片和底层数组关系

```go
package main

import "fmt"

func main() {
	s1 :=[] int {1,2,3 }
	fmt.Println("s1 = ",s1)  // s1 =  [1 2 3]
	arr := [] int {1,2,3,4,5,6}
	s2 := arr[:] //  s2，是数组 arr 的引用。
	s3 := arr[1:5]   // 将 arr 中从下标 startIndex 到 endIndex-1 下的元素创建为一个新的切片。
	s4:= arr[2:]
	s5:= arr[:5]
	s6 := s3[1:3]
	fmt.Println("s2 = ",s2) // s2 =  [1 2 3 4 5 6]
	fmt.Println("s3 = ",s3) // s3 =  [2 3 4 5]
	fmt.Println("s4 = ",s4) // s4 =  [3 4 5 6]
	fmt.Println("s5 = ",s5) // s5 =  [1 2 3 4 5]
	fmt.Println("s6 = ",s6) // s6 =  [3 4]
	s2[0] = -2
	fmt.Println("arr = ",arr) // arr =  [-2 2 3 4 5 6]
	s3[1] = -3
	fmt.Println("arr = ",arr) // arr =  [-2 2 -3 4 5 6]
	s4[2] = -4
	fmt.Println("arr = ",arr) // arr =  [-2 2 -3 4 -4 6]
	s5[3] = -5
	fmt.Println("arr = ",arr) // arr =  [-2 2 -3 -5 -4 6]
	s6[1] = -6
	fmt.Println("arr = ",arr) // arr =  [-2 2 -3 -6 -4 6]
}
```



![image-20210127224436787](\Pictures\从0开始学GO之结构体、切片与Range\D_从0开始学GO之结构体、切片与Range.png)



### len() 、cap() 、append、copy函数

切片是可索引的，并且可以由 len() 方法获取长度。

切片提供了计算容量的方法 cap() 可以测量切片最长可以达到多少。

append函数向 slice 尾部添加数据，返回新的 slice 对象。append函数会智能地底层数组的容量增长，一旦超过原底层数组容量，通常**以2倍容量重新分配底层数组，并复制原来的数据**。

函数 copy 在两个 slice 间复制数据，**复制⻓度以 len 小的为准**，两个 slice 可指向同⼀底层数组。

```go
package main

import "fmt"

func main() {
	s := make([]int, 0, 1) // 长度 0  容量 1
	c := cap(s)
	printSlice(s)
	for i := 0; i < 50; i++ {
		s = append(s, i)
		if n := cap(s); n > c {
			printSlice(s)
			c = n
		}
	}
	printSlice(s)
	fmt.Println("==1===============")
	s1 := make([]int, 5, 7)
	copy(s1, s)    // dst:s1, src:s
	printSlice(s) // len=50 cap=64
	printSlice(s1) // len=5 cap=7
	fmt.Println("==2===============")
	s1[0] = -1  // s1 改变 s 不改变
	printSlice(s)
	printSlice(s1) // len=5 cap=7 slice=[-1 1 2 3 4]
	fmt.Println("==3===============")
	data := [...]int{0, 1, 2, 3, 4, 5}  // 长度 6
	fmt.Println("data = ",data)
	s2 := make([]int, 5, 7) // 长度 5  容量 7
	printSlice(s2) // len=5 cap=7 slice=[0 0 0 0 0]
	s2 = data[:4] //{0, 1, 2, 3}
	s2[0] = 100 // 对 data 有影响
	fmt.Println("data = ",data) // data =  [100 1 2 3 4 5]
	c2 := cap(s2)
	printSlice(s2) // len=4 cap=6 slice=[0 1 2 3]
	fmt.Println("==4===============")
	for i := 0; i < 10; i++ {
		s2 = append(s2, i)
		if n := cap(s2); n > c2 {
			printSlice(s2)
			c = n
		}
	}
	fmt.Println("==5===============")
	// 以2倍容量重新分配底层数组，并复制原来的数据
	s2[0] = -1  // 对 data 无影响
	printSlice(s2) // len=14 cap=24 slice=[-1 1 2 3 0 1 2 3 4 5 6 7 8 9]
	fmt.Println("data = ",data) // data =  [0 1 2 3 0 1]
}

func printSlice(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}
```



![image-20210127231432725](Pictures\从0开始学GO之结构体、切片与Range\E_从0开始学GO之结构体、切片与Range.png)



## range 范围

Go 语言中 range 关键字用于 for 循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素。在数组和切片中它返回元素的索引和索引对应的值，在集合中返回 key-value 对。



```go
package main
import "fmt"
func main() {
    //这是我们使用range去求一个slice的和。使用数组跟这个很类似
    nums := []int{2, 3, 4}
    sum := 0
    for _, num := range nums {
        sum += num
    }
    fmt.Println("sum:", sum)
    //在数组上使用range将传入index和值两个变量。上面那个例子我们不需要使用该元素的序号，所以我们使用空白符"_"省略了。有时侯我们确实需要知道它的索引。
    for i, num := range nums {
        if num == 3 {
            fmt.Println("index:", i)
        }
    }
    //range也可以用在map的键值对上。
    kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }
    //range也可以用来枚举Unicode字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
    for i, c := range "go" {
        fmt.Println(i, c)
    }
}
```



![image-20210127231619240](Pictures\从0开始学GO之结构体、切片与Range\F_从0开始学GO之结构体、切片与Range.png)





## 参考：

[https://www.runoob.com/go/go-structures.html](https://www.runoob.com/go/go-structures.html)

[https://www.runoob.com/go/go-range.html](https://www.runoob.com/go/go-range.html)



个人能力有限，如有错误或者其他建议，敬请告知欢迎探讨，谢谢!