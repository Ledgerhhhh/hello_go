package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"testing"
	"time"
)

func TestCorn(t *testing.T) {
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
