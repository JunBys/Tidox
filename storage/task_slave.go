package storage

import (
	"time"
)

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
				(*taskList)[n].ModifyTime = time.Now()
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

// RemoveTaskSlave
// 从指定 masterID 的所有任务从属列表中，移除匹配 taskID 的子任务。
// - 支持多个 masterID 任务列表的情况（遍历删除）。
// - 返回总共删除的任务数量。
// 删除逻辑：
//  1. 调用 getSlaveList(masterID) 获取所有对应的从属任务切片引用。
//  2. 遍历每个切片，将不匹配 taskID 的任务保留，匹配的任务移除。
//  3. 统计所有切片中被移除的任务总数，并返回。
func (t *Task) RemoveTaskSlave(masterID, taskID int) int {
	t.Mut.Lock()
	defer t.Mut.Unlock()

	removeCount := 0
	// 获取匹配masterID的task的[]TaskDataSlave
	// 然后对其遍历，一般只有一条
	slaveList := t.getSlaveList(masterID)
	for n, slave := range slaveList {

		// 将不匹配taskID的任务添加到newSlave
		var newSlave []TaskDataSlave
		for _, slaveData := range *slave { //slave is slaveData
			if slaveData.ID != taskID {
				newSlave = append(newSlave, slaveData)
			} else {
				removeCount++
			}
		}
		*slaveList[n] = newSlave
	}
	return removeCount
}

// 根据masterID匹配符合条件的 masterTask，然后将task的slaveTask追加到切片
// 根据taskID在返回的slaveTask中匹配slaveTaskID匹配的项目
// 更新其Name字段
func (t *Task) UpdateTaskSlave(masterID, taskID int, taskName string) int {
	t.Mut.Lock()
	defer t.Mut.Unlock()

	updateCount := 0
	// 获取匹配masterID的task的[]TaskDataSlave
	// 然后对其遍历，一般只有一条
	slaveSlice := t.getSlaveList(masterID)
	for _, slavePtr := range slaveSlice {

		// 这里使用了切片的浅拷贝特性，既复制切片并非值拷贝，而是切片本身的拷贝，
		// 共同指向同一份底层数组，因为切片原本就是一个strct
		// { Data uintptr //指向底层数组的指针, Len int, Cap int}
		slave := *slavePtr
		for n := range slave {
			if slave[n].ID == taskID {
				slave[n].Name = taskName
				slave[n].ModifyTime = time.Now()
				updateCount++
			}
		}
	}
	return updateCount
}
