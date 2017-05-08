package task

import (
	"fmt"
	"time"
)

// 随机生产的已经待机时间
var RestTime float64 = 0

// 开始时间
var BeginTime time.Time = time.Date(2017, 4, 23, 0, 0, 0, 0, time.Local)

type Task interface {
	Exec()
}

type Producer struct {
}

// 使用多少个协程并行执行数据生成
// 一个协程生成一个文件
// 推荐300个，根据操作系统查看cpu使用率
const ThreadNumber int = 300

func (p *Producer) Exec() {
	fmt.Println("The machines will running!")
	fmt.Println(time.Now())
	readAbort()
	//	不重命名 直接生成目标文件，原因 linux文件共享，读写时候也可以重命名导致文件关闭不了，fd只增不减
	//	go WritedAndRename()
	for i := 0; i < ThreadNumber; i++ {
		sleep := time.Duration(3000000000) //0.3
		time.Sleep(sleep)
		fmt.Println("The thread-" + fmt.Sprint(i) + " is running")
		go func() {
			for true {
				MiTask()
			}
		}()
	}
}

// 工作am8:00-pm9:00
// 改为任何时间都工作 代码注释掉了
func WorkTime() {
	//	now := time.Now()
	//	hour := now.Hour()
	//	//第二天早上八点的距离
	//	sleep := time.Date(2017, 1, 2, 8, 0, 0, 0, time.Local).UnixNano() - time.Date(2017, 1, 1, now.Hour(), now.Minute(), now.Second(), now.Nanosecond(), time.Local).UnixNano()
	//	if hour > 21 {
	//		fmt.Println(sleep)
	//		time.Sleep(time.Duration(sleep))
	//	}
}
