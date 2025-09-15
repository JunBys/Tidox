package storage

import "time"

// addTask 是 Task 结构体的内部辅助方法，负责将新任务添加到指定的任务列表中。
// 它会自动生成唯一的任务 ID，并记录任务创建时间。
// 该方法封装了任务添加的核心逻辑，供具体的 AddTaskNew、AddTaskNow 和 AddTaskOld 调用，避免代码重复。
//   - taskList: 指向任务列表切片的指针，任务将被追加到此列表中
//   - taskName: 新任务的名称
func (t *Task) addTask(taskList *[]TaskData, taskName string) {
	t.Mut.Lock()
	defer t.Mut.Unlock()

	data := TaskData{
		ID:         t.getAutoID(),
		Name:       taskName,
		CreateTime: time.Now(),
	}
	*taskList = append(*taskList, data)
}

// removeTask 是 Task 结构体的内部辅助方法，用于从指定任务列表中移除所有 ID 匹配的任务。
// 它会遍历任务列表，过滤掉所有 ID 等于给定 id 的任务，并返回实际删除的任务数量。
//   - id: 要删除的任务的唯一标识符
//   - taskList: 指向任务列表切片的指针，任务将在此列表中被删除
//   - 返回：实际删除的任务数量
func (t *Task) removeTask(id int, taskList *[]TaskData) int {
	t.Mut.Lock()
	defer t.Mut.Unlock()

	var filtered []TaskData
	for _, task := range *taskList {
		if task.ID != id {
			filtered = append(filtered, task)
		}
	}

	removedCount := len(*taskList) - len(filtered)
	*taskList = filtered
	return removedCount
}

// updateTask 是 Task 结构体的内部辅助方法，用于从指定任务列表中更新所有 ID 匹配的Name字段。
// 它会遍历任务列表， 更新ID等与给定id的Name字段，最终返回更新的任务总数。
//   - id: 要更新的任务的唯一标识符
//   - taskList: 指向任务列表切片的指针，任务将在此列表中更新Name字段
//   - taskName：将指定id的Name字段更新成此值
//   - 返回值：实际删除的任务数量
func (t *Task) updateTask(id int, taskName string, taskList *[]TaskData) int {
	t.Mut.Lock()
	defer t.Mut.Unlock()

	replaceSum := 0
	for n := range *taskList {
		if (*taskList)[n].ID == id {
			replaceSum++
			(*taskList)[n].Name = taskName
			(*taskList)[n].ModifyTime = time.Now()
		}
	}

	return replaceSum
}

// getTaskByID是Task结构的的辅助内部方法，用于通过ID获取task，实际上就是为
// AddTaskNew、AddTaskNow 和 AddTaskOld 做的功能实现封装
//   - id: 要更新的任务的唯一标识符
//   - taskList: 指向任务列表切片的指针，任务将在此列表中更新Name字段
//   - 返回值：匹配ID的Task list
func (t *Task) getTaskByID(id int, taskList *[]TaskData) []TaskData {
	t.Mut.Lock()
	defer t.Mut.Unlock()

	var newTask []TaskData
	for _, task := range *taskList {
		if task.ID == id {
			newTask = append(newTask, task)
		}
	}
	return newTask
}

// GetTaskByTime 根据时间范围筛选任务
//   - useCreateTime: 为 true 时使用任务的 CreateTime 进行筛选；为 false 时使用 ModifyTime
//   - s: 起始时间（包含）
//   - d: 结束时间（包含）
//   - taskList: 要筛选的任务列表
//   - 返回值： 返回满足时间范围条件的任务切片
func (t *Task) getTaskByTime(s, d time.Time, useCreatTime bool, taskList *[]TaskData) []TaskData {
	t.Mut.Lock()
	defer t.Mut.Unlock()

	var newTaskList []TaskData
	for _, task := range *taskList {
		// 确定要筛选的时间类型：CreateTime 或 ModifyTime
		var taskTime time.Time
		if useCreatTime {
			taskTime = task.CreateTime
		} else {
			taskTime = task.ModifyTime
		}

		// 时间比较，taskTime 不在s之前而且不再d之后则成立
		if !taskTime.Before(s) && !taskTime.After(d) {
			newTaskList = append(newTaskList, task)
		}
	}
	return newTaskList
}
