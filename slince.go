package gohelper
// 返回 在slice 中的索引值，没有找到的话返回-1
func IndexOfSlice(item string, slice []string) int {
	for index, value := range slice {
		if item == value {
			return index
		}
	}
	return -1
}