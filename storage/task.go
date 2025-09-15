package storage

import "time"

// AddTaskNew 添加任务到TaskNew
func (t *Task) AddTaskNew(taskName string) {
	t.addTask(&t.TaskNew, taskName)
}

// AddTaskNow 添加任务到TaskNow
func (t *Task) AddTaskNow(taskName string) {
	t.addTask(&t.TaskNow, taskName)
}

// AddTaskOld 添加任务到TaskOld
func (t *Task) AddTaskOld(taskName string) {
	t.addTask(&t.TaskOld, taskName)
}

// 获取最新的自增id，并维护关系
func (t *Task) getAutoID() int {
	t.AutoID++
	return t.AutoID
}

// RemoveTaskNew 删除任务从TaskNew，通过TaskID，返回删除的数据总数
func (t *Task) RemoveTaskNew(id int) int {
	return t.removeTask(id, &t.TaskNew)
}

// RemoveTaskNow 删除任务从TaskNow，通过TaskID，返回删除的数据总数
func (t *Task) RemoveTaskNow(id int) int {
	return t.removeTask(id, &t.TaskNow)
}

// RemoveTaskOld 删除任务从TaskOld，通过TaskID，返回删除的数据总数
func (t *Task) RemoveTaskOld(id int) int {
	return t.removeTask(id, &t.TaskOld)
}

// UpdateTaskNew 更新TaskNew中的Name字段，根据ID字段匹配，并返回更新的次数
func (t *Task) UpdateTaskNew(id int, taskName string) int {
	return t.updateTask(id, taskName, &t.TaskNew)
}

// UpdateTaskNow 更新TaskNow中的Name字段，根据ID字段匹配，并返回更新的次数
func (t *Task) UpdateTaskNow(id int, taskName string) int {
	return t.updateTask(id, taskName, &t.TaskNow)
}

// UpdateTaskOld 更新TaskOld中的Name字段，根据ID字段匹配，并返回更新的次数
func (t *Task) UpdateTaskOld(id int, taskName string) int {
	return t.updateTask(id, taskName, &t.TaskOld)
}

// GetTaskNewByID 按照ID获取TaskNew列表, 他返回TaskDataList类型（[]TaskData 的别名）
// 方便调用TaskDataList类的方法，实现排序等功能
func (t *Task) GetTaskNewByID(id int) TaskDataList {
	return t.getTaskByID(id, &t.TaskNew)
}

// GetTaskNowByID 按照 ID 获取 TaskNow 列表，返回 TaskDataList 类型（[]TaskData 的别名）
// 方便调用 TaskDataList 的方法，例如排序等
func (t *Task) GetTaskNowByID(id int) TaskDataList {
	return t.getTaskByID(id, &t.TaskNow)
}

// GetTaskOldByID 按照 ID 获取 TaskOld 列表，返回 TaskDataList 类型（[]TaskData 的别名）
// 方便调用 TaskDataList 的方法，例如排序等
func (t *Task) GetTaskOldByID(id int) TaskDataList {
	return t.getTaskByID(id, &t.TaskOld)
}

// GetTaskNewByTime 根据时间范围查询TaskNew, useCreateTime false时使用modify时间匹配
func (t *Task) GetTaskNewByTime(s, d time.Time, useCreateTime bool) TaskDataList {
	return t.getTaskByTime(s, d, useCreateTime, &t.TaskNew)
}

// GetTaskNowByTime 根据时间范围查询TaskNow, useCreateTime false时使用modify时间匹配
func (t *Task) GetTaskNowByTime(s, d time.Time, useCreateTime bool) TaskDataList {
	return t.getTaskByTime(s, d, useCreateTime, &t.TaskNow)
}

// GetTaskOldByTime 根据时间范围查询TaskOld, useCreateTime false时使用modify时间匹配
func (t *Task) GetTaskOldByTime(s, d time.Time, useCreateTime bool) TaskDataList {
	return t.getTaskByTime(s, d, useCreateTime, &t.TaskOld)
}
