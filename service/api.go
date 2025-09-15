package service

import "tidox/storage"

// GetTask 用来获取Storage.Task, 格式化时间为人类可读
//   - t：控制返回NewTask、NowTask、OldTask
func (a *API) GetTask(t string) *[]TaskData {
	var task []storage.TaskData
	switch t {
	case "NewTask":
		task = a.Task.TaskNew
	case "NowTask":
		task = a.Task.TaskNow
	case "OldTask":
		task = a.Task.TaskOld
	}

	var data []TaskData
	for _, i := range task {
		d := TaskData{
			ID:         i.ID,
			Name:       i.Name,
			Slave:      i.Slave,
			CreateTime: i.CreateTime.Format("06/01/02 15:04"),
			ModifyTime: i.ModifyTime.Format("06/01/02 15:04"),
		}
		data = append(data, d)
	}
	return &data
}
