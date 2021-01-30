package A_File

import "fmt"
import "B_File"

func init()  {
	fmt.Println("A_File init")
}

func a_print() {
	fmt.Println("A_File a_print")
}

func Print() {  //  首字母大写就可以让标识符对外可见
	a_print()
	B_File.Print()
}