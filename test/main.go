package main

import (
	"fmt"
	"tidox/pkg/storage"
	"time"
)

func main() {
	db := storage.NewTask()
	db.AddTaskNew("哈哈哈")
	db.AddTaskNew("哈哈哈")
	db.AddTaskNow("哈哈哈")
	db.AddTaskOld("哈哈哈")

	fmt.Println(db.UpdateTaskNow(3, "淅沥沥"))

	fmt.Println(db.RemoveTaskNew(1))

	fmt.Println(db.GetTaskNewByID(2))

	time.Sleep(2 * time.Second)
	db.AddTaskNew("dongdong")
	fmt.Println(db.GetTaskNewByTime(time.Now().Add(-3*time.Second), time.Now(), true))

}
