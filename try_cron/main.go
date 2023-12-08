package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

func main() {
	spec := "* * * * * *"

	c := cron.New(cron.WithSeconds())
	entryID, _ := c.AddFunc(spec, func() {
		// Your task logic
		fmt.Println("Hello, world!")
	})

	// 获取 Entry 对象
	entry := c.Entry(entryID)

	// 在需要的时候进行更精细的控制
	entry.Job.Run()
	time.Sleep(10 * time.Second)

}

type MyJobWrapper struct {
	AdditionalMessage string
}

func Wrap(j cron.Job) cron.Job {
	return cron.FuncJob(func() {
		fmt.Println("Before job:")
		// 在这里可以使用 w 中的数据进行修改
		if job, ok := j.(*MyJob); ok {
			job.Message += " - Modified"
		}
		j.Run()
		fmt.Println("After job:")
	})
}

type MyJob struct {
	Message string
}

func (n *MyJob) Run() {
	fmt.Println("执行任务...")
}
