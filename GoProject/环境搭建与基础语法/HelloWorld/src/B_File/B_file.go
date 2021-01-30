package B_File

import "fmt"
import "C_File"

func init()  {
	fmt.Println("B_File init")
}

func b_print() {
	fmt.Println("B_File b_print")
}

func Print() {  //  首字母大写就可以让标识符对外可见
	b_print()
	C_File.Print()
}