package timer

import (
	"fmt"
	"testing"
	"time"
)

func TestCreate(t *testing.T) {
	timer := NewTimer(time.Second)
	timer.Start()

	// 1秒的定时器,执行10次
	timer.AddTimer(time.Minute, -1, func(*Timer) {
		fmt.Println("timer 1 second")
	})

	// 10秒的定时器,执行1次
	timeid2 := timer.AddTimer(10*time.Second, 1, func(*Timer) {
		fmt.Println("timer 10 second")
	})

	timer.RemoveTimer(timeid2)

	// 永久执行的1秒定时器任务
	timer.AddTimer(time.Second, -1, func(*Timer) {
		fmt.Println("timer 1 second forever")
	})
}
