package storage

// 按照时间赠序排序[]TaskData
func (d TaskDataList) SortByTime(useCreateTime bool) TaskDataList {
	// sort.Slice(d, func(i, j int) bool {
	// 	d[i].CreateTime
	// })
	return nil
}

// 按照时间降序排序 []TaskData
func (d TaskDataList) SortByTimeDesc() TaskDataList {
	return nil
}

// 按照时间赠序排序[]TaskData
func (d TaskDataSlaveList) SortByTime() TaskDataSlaveList {
	return nil
}

// 按照时间降序排序 []TaskData
func (d TaskDataSlaveList) SortByTimeDesc() TaskDataSlaveList {
	return nil
}
