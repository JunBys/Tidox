package main

import (
	"fmt"
	"tidox/api"
	"tidox/storage"
)

// 入口函数
func main() {
	st := storage.NewTask()
	apiServer := api.NewServer(st)

	st.AddTaskNew("新任务1")
	st.AddTaskNew("新任务2")
	st.AddTaskNew("新任务3")
	st.AddTaskNow("当前任务")

	err := apiServer.Run(":8000")
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
