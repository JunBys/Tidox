package main

import (
	"fmt"
	"tidox/pkg/storage"
	"time"
)

func main() {
	db := storage.NewTask()
	db.AddTaskNew("n1")
	db.AddTaskNew("n2")
	db.AddTaskNow("哈哈哈")
	db.AddTaskOld("哈哈哈")

	fmt.Println(db.UpdateTaskNow(3, "淅沥沥"))

	// fmt.Println(db.RemoveTaskNew(1))

	fmt.Println(db.GetTaskNewByID(2))

	// time.Sleep(2 * time.Second)
	db.AddTaskNew("n3")
	db.AddTaskSlave(5, "你好我是自任务")
	fmt.Println(db.GetTaskNewByTime(time.Now().Add(-3*time.Second), time.Now(), true).SortByTime(true, true))

}
