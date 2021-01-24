[TOC]

更多参见：[从0开始学GO之目录](https://blog.csdn.net/leacock1991/article/details/112853343)

指针是一个代表着某个内存地址的值。一个指针变量指向了一个值的内存地址。



## Go语言中的指针

Go语言中指针**默认值为 nil**，操作符 “&” 取变量地址， “*” 通过指针访问目标对象，**不⽀持指针运算**，不⽀持 “->” 运算符，直接⽤ “.” 访问目标成员。



## 指针使用

指针使用流程：

- 定义指针变量。
- 为指针变量赋值。
- 访问指针变量中指向地址的值。



```go
package main
import "fmt"

func main() {
	var a = 20     /* 声明实际变量 */
	var point *int /* 声明指针变量 */

	fmt.Println("指针默认值： ", point)
	point = &a  /* 指针变量的存储地址 */
	fmt.Printf("a 变量的地址是: %p\n", &a  ) //操作符 "&" 取变量地址

	/* 指针变量的存储地址 */
	fmt.Printf("point 变量储存的指针地址: %p\n", point)
	/* 使用指针访问值 */
	fmt.Printf("修改前 *point 变量的值: %d\n", *point)
	fmt.Printf("修改前 a 变量的值: %d\n", a)
	/* *p操作指针所指向的内存 */
	*point = -20 //*p，即为a
	fmt.Printf("修改后 *point 变量的值: %d\n", *point)
	fmt.Printf("修改后 a 变量的值: %d\n", a)
}
```

![image-20210124215650115](\Pictures\从0开始学GO之指针\A_从0开始学GO之指针.png)



## 指针数组



指针数组：


> var ptr [Num]*T  // T为类型




```go
package main
import "fmt"

func main() {
	a := []int{10,100,110}
	var ptr [5]*int
	for i,value := range ptr {
		fmt.Printf("ptr[%d] = ", i )
		fmt.Print(value,"\t")
	}
	// var ptr1 *int = ptr+1 // erro invalid operation: ptr + 1 (mismatched types [5]*int and int) 不⽀持指针运算
	fmt.Printf("\n")
	for i,value:= range a {
		fmt.Printf(" &value = %d\t", &value )
		ptr[i] = &a[i] // 整数地址赋值给指针数组   不能赋值 &value
		fmt.Printf("a[%d] = %d\t", i, a[i] )
		fmt.Printf("\n")
	}
	fmt.Println("==========")
	for  i := 0; i < 3; i++ {
		fmt.Printf("a[%d] = %d\t", i,a[i] )
		fmt.Printf("*ptr[%d] = %d\t", i,*ptr[i] )
		fmt.Printf("\n")
	}
	fmt.Println("==========")
	fmt.Print("ptr[3] = ", ptr[3],"\t")
	fmt.Print("ptr[4] = ", ptr[4],"\t")

}
```

![image-20210124222807210](\Pictures\从0开始学GO之指针\B_从0开始学GO之指针.png)



## 指向指针的指针

如果一个指针变量存放的又是另一个指针变量的地址，则称这个指针变量为指向指针的指针变量。

当定义一个指向指针的指针变量时，第一个指针存放第二个指针的地址，第二个指针存放变量的地址：

![img](\Pictures\从0开始学GO之指针\C_从0开始学GO之指针.png)



指向指针的指针:


> var ptr **int




```go

package main

import "fmt"

func main() {

	var a float32
	var ptr *float32
	var pptr **float32
	a = 3.14
	/* 指针 ptr 地址 */
	ptr = &a
	/* 指向指针 ptr 地址 */
	pptr = &ptr

	/* 获取 pptr 的值 */
	fmt.Printf("变量 a = %f\n", a )
	fmt.Printf("指针变量 *ptr = %f\n", *ptr )
	fmt.Printf("指向指针的指针变量 **pptr = %f\n", **pptr)

	var b = [2]int{10, 20}
	var pptrarr [2]**int
	for i := 0; i < 2; i++ {
		//pptrarr[i] = &(&b[i]) // erro  cannot take the address of &b[i]
		tmp := &b[i]
		pptrarr[i] = &tmp
	}

	for i, _ := range pptrarr{
		fmt.Print("i = ",i," \tpptrarr[i] = ", pptrarr[i] ,"\t*pptrarr[i] = ", *pptrarr[i],"\t**pptrarr[i] = ", **pptrarr[i],"\n")
	}
}
```

![image-20210124225057338](\Pictures\从0开始学GO之指针\D_从0开始学GO之指针.png)



## 指针作为函数参数

Go 语言允许向函数传递指针，只需要在函数定义的参数上设置为指针类型即可。

```go
package main

import "fmt"

func swap(x *interface{}, y *interface{}){
	*x, *y = *y, *x
}

func main() {
	/* 定义局部变量 */
	var a interface{}= 100
	var b interface{}= 200
	fmt.Print("a地址 : ", &a ,"\t b地址 : ", &b ,"\n")
	fmt.Print("交换前 a 的值 : ", a ,"\n")
	fmt.Print("交换前 b 的值 : ", b ,"\n")
	swap(&a, &b)
	fmt.Print("交换后 a 的值 : ", a ,"\n")
	fmt.Print("交换后 b 的值 : ", b ,"\n")
	fmt.Println("============")
	a = "Hello"
	b = "World"
	fmt.Print("a地址 : ", &a ,"\t b地址 : ", &b ,"\n")
	fmt.Print("交换前 a 的值 : ", a ,"\n")
	fmt.Print("交换前 b 的值 : ", b ,"\n")
	swap(&a, &b)
	fmt.Print("交换后 a 的值 : ", a ,"\n")
	fmt.Print("交换后 b 的值 : ", b ,"\n")
}

```



![image-20210124225645597](\Pictures\从0开始学GO之指针\E_从0开始学GO之指针.png)

个人能力有限，如有错误或者其他建议，敬请告知欢迎探讨，谢谢!