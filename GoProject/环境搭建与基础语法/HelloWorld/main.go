package main

import (
	"A_File"
	"fmt"
)

var x = 1024

func init()  {
	fmt.Println(x)
}
func main()  {
	A_File.Print() //  首字母大写就可以让标识符对外可见
	fmt.Println("Hello World！")
}