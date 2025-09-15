package storage

import (
	"sync"
	"time"
)

type TaskData struct {
	ID         int
	Name       string
	Slave      []TaskDataSlave
	CreateTime time.Time
	ModifyTime time.Time
}

type TaskDataSlave struct {
	ID         int
	Name       string
	CreateTime time.Time
	ModifyTime time.Time
}

// TaskNew 代办任务列表
// TaskNow 在办任务列表
// TaskOld 完成任务列表
type Task struct {
	AutoID  int // 自增id
	Mut     sync.Mutex
	TaskNew []TaskData
	TaskNow []TaskData
	TaskOld []TaskData
}

type TaskDataList []TaskData
type TaskDataSlaveList []TaskDataSlave

// NewTask 创建一个任务数据库实例
func NewTask() *Task {
	return &Task{
		TaskNew: []TaskData{},
		TaskNow: []TaskData{},
		TaskOld: []TaskData{},
	}
}
