package main

import (
	"github.com/jiebozeng/golangutils/debugs"
	"runtime"
)

func main() {
	debugs.Print("test")
	runtime.Goexit()
}
