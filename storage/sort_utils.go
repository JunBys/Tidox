package storage

import (
	"sort"
	"time"
)

// 泛型排序函数 sortByTime 使用 Go 1.18+ 引入的泛型特性，支持对任意类型的切片进行时间排序。
// 它通过函数参数 getTime 提供从泛型元素中提取时间字段的能力，从而保持通用性。
//   - T any: 表示泛型类型参数 T，可以是任何类型（any 是 interface{} 的类型别名）
//     我们在此处依赖调用者提供 getTime 函数来告诉 sortByTime 如何从 T 中获取 time.Time 字段。
//   - d []T: 需要排序的数据切片，类型为 T 的切片
//   - revers bool: 是否反转排序（true 表示最新时间在前）
//   - getTime func(T) time.Time: 用于从元素中提取时间字段的函数（回调函数：由调用方提供）
//   - 返回值： - 排序后的切片（原切片被就地排序）
func sortByTime[T any](d []T, revers bool, getTime func(T) time.Time) []T {
	sort.Slice(d, func(i, j int) bool {
		itime := getTime(d[i])
		jtime := getTime(d[j])

		if revers {
			// 比较时间，jtime是否早于itime
			return jtime.Before(itime)
		}
		return itime.Before(jtime)
	})
	return d
}

// SortByTime 是 []TaskData 的类型别名，我们为它添加了排序方法。
// 使用 sortByTime 泛型函数实现按时间排序。
//   - useCreateTime bool: 指定使用 CreateTime（true）或 ModifyTime（false）作为排序字段
//   - revers bool: 是否反转排序（true 表示时间新 → 旧）
//   - 返回值：排序后的 TaskDataList 副本（就地排序）
func (d TaskDataList) SortByTime(useCreateTime, revers bool) TaskDataList {
	return sortByTime(d, revers, func(task TaskData) time.Time {
		if useCreateTime {
			return task.CreateTime
		}
		return task.ModifyTime
	})
}

// SortByTime 是 []TaskDataSlave 的类型别名，同样支持按时间字段排序。
// 使用 sortByTime 泛型函数实现按时间排序。
//   - useCreateTime bool: 指定使用 CreateTime（true）或 ModifyTime（false）作为排序字段
//   - revers bool: 是否反转排序（true 表示时间新 → 旧）
//   - 返回值：排序后的 TaskDataList 副本（就地排序）
func (d TaskDataSlaveList) SortByTime(useCreateTime, revers bool) TaskDataSlaveList {
	return sortByTime(d, revers, func(task TaskDataSlave) time.Time {
		if useCreateTime {
			return task.CreateTime
		}
		return task.ModifyTime

	})
}
