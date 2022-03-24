package timer

import (
	"container/list"
	"sync"
	"time"
)

type Timer struct {
	interval          time.Duration // 时间间隔
	ticker            *time.Ticker
	tasks             *list.List
	currentTime       int
	addTaskChannel    chan Task // 新增任务channel
	removeTaskChannel chan int  // 删除任务channel
	stopChannel       chan bool // 停止定时器channel
	accid             int
	acclock           sync.RWMutex
}

type Task struct {
	delay    time.Duration // 延迟时间
	key      int           // 定时器唯一标识, 用于删除定时器
	cb       func(*Timer)  // 回调函数
	Count    int           // 执行次数
	ExecTime int           // 执行时间
}

func (this *Task) addExecTime(current int) {
	this.ExecTime = current + int(this.delay.Seconds())
}

// New 创建定时器
func NewTimer(interval time.Duration) *Timer {
	if interval <= 0 {
		return nil
	}
	tw := &Timer{
		interval:          interval,
		tasks:             list.New(),
		currentTime:       int(time.Now().Unix()),
		addTaskChannel:    make(chan Task),
		removeTaskChannel: make(chan int),
		stopChannel:       make(chan bool),
		accid:             0,
	}

	return tw
}

// Start
func (tw *Timer) Start() {
	tw.ticker = time.NewTicker(tw.interval)
	go tw.start()
}

// Stop
func (tw *Timer) Stop() {
	tw.stopChannel <- true
}

// AddTimer 添加定时器 key为定时器唯一标识
func (tw *Timer) AddTimer(delay time.Duration, count int, cb func(*Timer)) int {
	if delay <= 0 || cb == nil {
		return -1
	}
	var key int
	tw.acclock.Lock()
	tw.accid++
	key = tw.accid
	tw.acclock.Unlock()

	tw.addTaskChannel <- Task{delay: delay, key: key, cb: cb, Count: count}
	return key
}

// RemoveTimer 删除定时器 key为添加定时器时传递的定时器唯一标识
func (tw *Timer) RemoveTimer(key int) {
	if key == -1 {
		return
	}
	tw.removeTaskChannel <- key
}

func (tw *Timer) start() {
	for {
		select {
		case <-tw.ticker.C:
			tw.tickHandler()
		case task := <-tw.addTaskChannel:
			tw.addTask(&task)
		case key := <-tw.removeTaskChannel:
			tw.removeTask(key)
		case <-tw.stopChannel:
			tw.ticker.Stop()
			return
		}
	}
}

func (tw *Timer) tickHandler() {
	tw.currentTime = int(time.Now().Unix())
	for e := tw.tasks.Front(); e != nil; {
		task := e.Value.(*Task)
		if task.ExecTime > tw.currentTime {
			e = e.Next()
			continue
		}

		go task.cb(tw)
		next := e.Next()
		if task.Count > 0 {
			task.Count--
		}
		if task.Count == 0 {
			tw.tasks.Remove(e)
		} else {
			task.addExecTime(tw.currentTime)
		}

		e = next
	}
}

// 新增任务到链表中
func (tw *Timer) addTask(task *Task) {
	task.addExecTime(tw.currentTime)
	tw.tasks.PushBack(task)
}

// 从链表中删除任务
func (tw *Timer) removeTask(key int) {
	for e := tw.tasks.Front(); e != nil; {
		task := e.Value.(*Task)
		if task.key == key {
			tw.tasks.Remove(e)
		}

		e = e.Next()
	}
}
