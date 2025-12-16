package nge

func NextGreaterElement(ar1, ar2 []int) []int {
	var stack []int
	nges := make(map[int]int)

	for _, v := range ar2 {
		for len(stack) > 0 && v > stack[len(stack)-1] {
			nges[stack[len(stack)-1]] = v
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, v)
	}

	var answer []int
	for _, v := range ar1 {
		if nge, ok := nges[v]; ok {
			answer = append(answer, nge)
			continue
		}
		answer = append(answer, -1)
	}

	return answer
}
