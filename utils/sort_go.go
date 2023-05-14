package utils

import (
	"sort"
)

/*
* 对map的key值进行排序
 */
func SortMap(m map[string]string) []string {
	var arr []string
	for k := range m {
		arr = append(arr, k)
	}
	sort.Strings(arr)
	return arr
}
