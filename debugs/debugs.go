package debugs

import "fmt"

// println clearly 
func Println(args ...interface{}) {
	fmt.Println("------------")
	for _, v := range args {
		fmt.Println(v)
	}
	fmt.Println("------------")
}

// print clearly not ln
func Print(args ...interface{}) {
	fmt.Println("------------")
	for _, v := range args {
		fmt.Print(v, " ")
	}
	fmt.Println("------------")
}
