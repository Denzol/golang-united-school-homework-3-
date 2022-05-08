package homework

func reverse(input []int64) (result []int64) {
	for i := 0; i < len(input)/2; i++ {
		var a int64 = 0
		a = input[i]
		input[i] = input[(len(input)-1)-i]
		input[(len(input)-1)-i] = a
	}
	result = input
	return result
}
