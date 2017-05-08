package main

/*
create by ppf 294881866@qq.com at 2017.4.18
*/
import (
	"flag"
	"fmt"
	"task"
	"os"
)

func main() {
	p := fmt.Println
	var syn chan string = make(chan string)
	// 参数
	arg_num := len(os.Args)
	fmt.Printf("the num of input is %d\n", arg_num)
	fmt.Printf("they are :\n")
	for i := 0; i < arg_num; i++ {
		fmt.Println(os.Args[i])
	}

	// 数据路径
	if len(os.Args) > 1 {
		path := *flag.String("path", "", "The input parameter for the data to save directory")
		if len(path) != 0 {
			task.SetPath(path)
		} else {
			path = os.Args[1]
			task.SetPath(path)
		}
	} else {
		task.SetPath("")
	}

	// 执行任务
	producer := task.Producer{}
	producer.Exec()

	syn <- "blocking"
	p("blocking  is over :", <-syn)
}
