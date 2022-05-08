package homework

func average(input [15]float32) (result float32) {
	count := 0
	var sum float32 = 0
	for i := 0; i < len(input); i++ {
		if input[i] != 0 {
			sum = sum + input[i]
			count++
		}
	}
	return sum / float32(count)
}
