[TOC]

更多参见：[从0开始学GO之目录](https://blog.csdn.net/leacock1991/article/details/112853343)



Go语言内置以下这些基础类型：  

> 布尔类型、整型、浮点类型、复数类型、字符串、字符类型、错误类型

也支持以下这些复合类型：

> 指针（ pointer）、数组（ array）、切片（ slice）、字典（ map）、通道（ chan）、结构体（ struct）、接口（ interface）



# 基础类型

## 布尔类型  

布尔类型与其他语言基本一致，关键字也为bool，可赋值为true和false 。布尔类型不能接受其他类型的赋值，不支持自动或强制的类型转换。  

```go
	var v1 bool
	v1 = true
	v2 := (1 == 2) // v2也会被推导为bool类型

	//布尔类型不能接受其他类型的赋值，不支持自动或强制的类型转换
	var b bool
	b = 1 // err, 编译错误 cannot use 1 (type untyped int) as type bool in assignment
	b = bool(1) // err, 编译错误 cannot convert 1 (type untyped int) to type bool
```



## 整型

|     类型     | 长度（字节） | 值范围 |
| :----------: | :----------: | :----: |
|     int8     | 1 | - 128 ~ 127 |
| uint8（byte) | 1 | 0 ~ 255 |
|    int16     | 2 | - 32 768 ~ 32 767 |
| uint16 | 2 |0 ~ 65 535|
| int32 | 4 |- 2 147 483 648 ~ 2 147 483 647|
| uint32 | 4 |0 ~ 4 294 967 295|
| int64 | 8 |- 9 223 372 036 854 775 808 ~ 9 223 372 036 854 775 807|
| uint64 | 8 |0 ~ 18 446 744 073 709 551 615|
| int | 平台相关 |平台相关|
| uint | 平台相关 |平台相关|
| uintptr | 同指针 |在32位平台下为4字节， 64位平台下为8字节|

int和int32在Go语言里被认为是两种不同的类型，编译器也不会帮你自动做类型转换  

```go
	var value2 int32
	value1 := 64 // value1将会被自动推导为int类型
	value2 = value1 // 编译错误 cannot use value1 (type int) as type int32 in assignment
	value2 = int32(value1) // 可以使用强制类型转换
```

在进行强制类型转换时注意数据精度损失  



## 浮点型

浮点型用于表示包含小数点的数据  Go语言定义了两个类型float32和float64，其中float32等价于C语言的float类型，float64等价于C语言的double类型  

|  类型   | 长度（字节） |                   说明                    |
| :-----: | :----------: | :---------------------------------------: |
| float32 |      4       |  小数位精确到7位，等价于C语言的float类型  |
| float64 |      8       | 小数位精确到15位，等价于C语言的double类型 |

```go
	var f1 float32
	f1 = 12
	f2 := 12.0 //  f2 被推导为 float64，不是 float32
	f1 = f2 // err  cannot use f2 (type float64) as type float32 in assignment
```



## 复数类型

复数实际上由两个实数（在计算机中用浮点数表示）构成，一个表示实部（real），一个表示虚部（imag）。

|    类型    | 长度（字节） |
| :--------: | :----------: |
| complex64  |      8       |
| complex128 |      16      |

```go
package main
import "fmt"
func main() {
	var v1 complex64 // 由2个float32构成的复数类型
	v1 = 3.2 + 12i
	v2 := 3.2 + 12i        // v2是complex128类型 ,  浮点型 被推导为 float64
	v3 := complex(3.2, 12) // v3结果同v2
	fmt.Println(v1, v2, v3)
	//内置函数real(v1)获得该复数的实部
	//通过imag(v1)获得该复数的虚部
	fmt.Println("real(v1) = ",real(v1))
	fmt.Println("imag(v1) = ",imag(v1))
}
```

![image-20210120111347503](\Pictures\从0开始学GO之基础与复合类型\A_从0开始学GO之基础与复合类型.png)



对于一个复数z = complex(x, y)，就可以通过Go语言内置函数real(z)获得该复数的实部，也就是x，通过imag(z)获得该复数的虚部，也就是y。  



## 字符串



```go
package main
import "fmt"
func main() {
	var str string                                    // 声明一个字符串变量
	str = "Hello World"                               // 字符串赋值
	ch := str[6]                                      // 取字符串的第一个字符
	fmt.Printf("str = %s, len = %d\n", str, len(str)) //内置的函数len()来取字符串的长度
	fmt.Printf("str[6] = %c, ch = %c\n", str[6], ch)
	str1 := str[0:5] // 不包括 下标5的字符
	fmt.Printf("str1 = %s, len = %d\n", str1, len(str1))
}
```

![image-20210120111904472](\Pictures\从0开始学GO之基础与复合类型\B_从0开始学GO之基础与复合类型.png)

可以参考 python 的字符串理解其特性



## 字符类型 

在Go语言中支持两个字符类型，一个是byte（实际上是uint8的别名），代表utf-8字符串的单个字节的值；另一个是rune，代表单个unicode字符。

```go
package main
import "fmt"

func main() {
	var ch1, ch2, ch3 ,ch4 byte
	ch1 = 'a'  // 字符赋值
	ch2 = 99   // 字符的ascii码赋值
	ch3 = '\n' // 转义字符
	ch4 = 'A' + 1 // ascii码 
	fmt.Printf("ch1 = %c, ch2 = %c, ch3 = %c, ch4 = %c", ch1, ch2, ch3, ch4)
}
```

![image-20210120112820893](\Pictures\从0开始学GO之基础与复合类型\C_从0开始学GO之基础与复合类型.png)



# 复合类型

## 数组



数组是指一系列同一类型数据的集合。数组中包含的每个数据被称为数组元素（element），一个数组包含的元素个数被称为数组的长度。

数组⻓度必须是常量，且是类型的组成部分。 [2]int 和 [3]int 是不同类型。

```go
package main
import "fmt"

func main() {
	var a [10]int
	for i := 0; i < 10; i++ { // 循环
		a[i] = i + 1
	}
	for i, v := range a { // 遍历
		fmt.Printf(" a[%d] = %d ", i, v) // 1 - 10
	}
	fmt.Printf("\n")
	// 切片 类似 python
	a1 := a[:5]              // 1- 5
	a2 := a[5:]              // 6 -10
	for i := 0; i < 5; i++ { // 循环
		fmt.Printf(" a1[%d] = %d ", i, a1[i])
		fmt.Printf(" a2[%d] = %d ", i, a2[i])
	}
	fmt.Printf("\n")
	b := [10]int{} //ok 未初始化元素值为 0
	c := [2] struct { x, y int32 }{{11,12},{22,23}} // 复杂类型数组
	d := [5]*int{&a[5],&a[6],&a[7]} // 指针数组
	e := [...]int{1, 2, 3}      // 通过初始化值确定数组长度
	f := [5]int{2: 100, 4: 200} // 通过索引号初始化元素，未初始化元素值为 0

	// 多维数组 未初始化元素值为 0
	g := [4][2]int{{10, 11}, {12, 13}, {14, 15}, {16, 17}}
	h := [...][2]float32{{10.0, 11.1}, {20.0, 21.1}}
	i := [2][2][2]int{0:{1:{0:-1}},1:{1:{0:2}}}
	//var f [2][2][2]float64 // 等同于[2]([2]([2]float64))

	fmt.Printf("len(b) = %d\n",len(b)) // len(b) = 10
	fmt.Printf("c[1].x = %d\n",c[1].x) // c[1].x = 22
	fmt.Printf("*d[2] = %d\n",*d[2])  // *d[2] = 8
	fmt.Println("e = ",e) // e =  [1 2 3]
	fmt.Println("f = ",f) // f =  [0 0 100 0 200]
	fmt.Println("g = ",g) // g =  [[10 11] [12 13] [14 15] [16 17]]
	fmt.Println("h = ",h) // h =  [[10 11.1] [20 21.1]]
	fmt.Println("i = ",i) // i =  [[[0 0] [-1 0]] [[0 0] [2 0]]]
}
```

![image-20210120135057814](\Pictures\从0开始学GO之基础与复合类型\D_从0开始学GO之基础与复合类型.png)



## map 

在C++/Java中， map一般都以库的方式提供，比如在C++中是STL的std::map<>。在Go中，使用map不需要引入任何库 ，是一种内置的数据结构，它是一个**无序**的key—value对的集合，比如以身份证号作为唯一键来标识一个人的信息。 在一个map里所有的键都是**唯一**的，而且必须是支持==和!=操作符的类型



```go
package main
import "fmt"

func main() {
	var m1 map[int]string  //只是声明一个map，没有初始化, 此为空(nil)map
	fmt.Println(m1 == nil) //true
	// m1[1] = "123" // err panic: assignment to entry in nil map
	fmt.Println("m1 = ", m1) // m1 =  map[]
	//m2, m3的创建方法是等价的
	m2 := map[int]string{}
	m3 := make(map[int]string)
	fmt.Println("m2 = ",m2) //m2 =  map[]
	fmt.Println("m3 = ",m3) //m3 =  map[]
	m4 := make(map[int]string, 2) //第2个参数指定容量 为2,会自动扩冲
	m4[0] = "Hello"
	m4[1] = "World"
	m4[3] = "Value"
	fmt.Println("m4 = ",m4)     //m4 =  map[0:Hello 1:World 3:Value]
	fmt.Println("len(m4) = ",len(m4)) //len(m4) =  3

	//迭代遍历，第一个返回值是key，第二个返回值是value
	for k, v := range m4 {
		fmt.Printf("%d --> %s\n", k, v)
		//0 --> Hello
		//1 --> World
		//3 --> Value
	}
	delete(m4,1) // 删除key值为1的map
	fmt.Println("m4 = ",m4) // m4 =  map[0:Hello 3:Value]

	//从map中查找一个特定的键, 第一个返回值是value(如果存在的话) 第二个 bool类型是否存在
	value, ok := m4[1] // not find key = 1 in m4
	if ok { // 找到了
		fmt.Println("find key = 1 in m4, value = ",value)
	} else {
		fmt.Println("not find key = 1 in m4")
	}
	value, ok = m4[0] // find key = 0 in m4, value =  Hello
	if ok { // 找到了
		fmt.Println("find key = 0 in m4, value = ",value)
	} else {
		fmt.Println("not find key = 0 in m4")
	}
}
```



![image-20210120142322248](\Pictures\从0开始学GO之基础与复合类型\E_从0开始学GO之基础与复合类型.png)



其他未介绍的见后续文章



个人能力有限，如有错误或者其他建议，敬请告知欢迎探讨，谢谢!