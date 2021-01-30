package C_File

import "fmt"

func init()  {
	fmt.Println("C_File init")
}

func c_print() {
	fmt.Println("C_File c_print")
}

func Print() {  //  首字母大写就可以让标识符对外可见
	c_print()
}