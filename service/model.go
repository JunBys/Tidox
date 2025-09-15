package service

import (
	"tidox/storage"
)

// TaskData 定义了底层数据存储类型，与storage.TaskData不同之处
// 在于他的时间类型是string，方便对数据进行展示使用
type TaskData struct {
	ID         int
	Name       string
	Slave      []storage.TaskDataSlave
	CreateTime string
	ModifyTime string
}

type API struct {
	Task *storage.Task
}

func NewAPI(s *storage.Task) *API {
	return &API{
		Task: s,
	}
}
