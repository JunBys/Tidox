package storage

import "time"

// getSlaveList 返回所有主任务 ID 为 masterID 的任务的子任务列表指针切片。
// 该方法不会加锁，调用者必须自行保证在持有 Task.Mut 锁的情况下调用，
// 否则并发访问可能导致数据竞态和不安全的读写行为。
// 返回值为指向各任务 Slave 切片的指针，允许调用者直接修改子任务列表。
func (t *Task) getSlaveList(masterID int) []*[]TaskDataSlave {
	var taskSlave []*[]TaskDataSlave
	tasks := []*[]TaskData{&t.TaskNew, &t.TaskNow, &t.TaskOld}
	for _, taskList := range tasks {
		for n, task := range *taskList {
			if task.ID == masterID {
				taskSlave = append(taskSlave, &(*taskList)[n].Slave)
			}
		}
	}
	return taskSlave
}

// AddTaskSlave 向匹配masterID的所有主任务中添加子任务
// 如果成功添加了子任务，则返回被添加了子任务的master任务个数
//   - masterID：添加子任务的master ID
//   - taskNmae：添加的任务内容
//   - 返回：被添加了子任务的master任务个数
func (t *Task) AddTaskSlave(masterID int, taskName string) int {
	t.Mut.Lock()
	defer t.Mut.Unlock()

	data := TaskDataSlave{
		ID:         t.getAutoID(),
		Name:       taskName,
		CreateTime: time.Now(),
	}

	slaveList := t.getSlaveList(masterID)
	for n := range slaveList {
		*slaveList[n] = append(*slaveList[n], data)
	}

	return len(slaveList)
}
