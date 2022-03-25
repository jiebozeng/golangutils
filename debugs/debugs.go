package debugs

import "fmt"

// println clearly
func Println(args ...interface{}) {
	fmt.Println("------8888------")
	for _, v := range args {
		fmt.Println(v)
	}
	fmt.Println("------8888------")
}

// print clearly not ln
func Print(args ...interface{}) {
	fmt.Println("------8888------")
	for _, v := range args {
		fmt.Print(v, " ")
	}
	fmt.Println("------8888------")
}
