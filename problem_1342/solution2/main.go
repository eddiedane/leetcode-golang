package main

func main() {
	println(numberOfSteps(14))
}

// Using bit operation
func numberOfSteps(num int) int {
	steps := 0

	for num != 0 {
		if (num & 1) == 0 {
			num >>= 1
		} else {
			num -= 1
		}

		steps += 1
	}

	return steps
}
