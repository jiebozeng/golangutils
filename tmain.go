package main

import (
	"github.com/jiebozeng/golangutils/debugs"
	"github.com/jiebozeng/golangutils/log"
	"github.com/nats-io/nats.go"
	"runtime"
)

/**
没啥用 不要看
*/
func main() {
	debugs.Print("test")
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Error("connect nats error", err)
	}

	defer nc.Close()
	runtime.Goexit()
}


