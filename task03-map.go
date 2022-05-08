package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	key := make([]int, 0)
	for k, _ := range input {
		key = append(key, k)
	}
	sort.Ints(key)
	value := make([]string, 0)
	for _, v := range key {
		value = append(value, input[v])
	}
	return value
}
