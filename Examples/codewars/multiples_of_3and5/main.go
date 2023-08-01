package main

func main() {
	var a = Multiple3And5(10)
	print("a:", a, "\n")
}

func Multiple3And5(number int) int {

	var multiples []int
	var sum int

	if number < 0 {
		return 0
	}

	for i := 0; i < number; i++ {
		if i >= 15 && i%15 == 0 {
			multiples = append(multiples, i)
			print("%15:", multiples, "i:", i, "\n")

		} else if i%5 == 0 {
			multiples = append(multiples, i)
			print("%5:", multiples, "i:", i, "\n")

		} else if i%3 == 0 {
			multiples = append(multiples, i)
			print("%3:", multiples, "i:", i, "\n")

		} else {
			continue
		}
	}
	for _, v := range multiples {
		sum += v
	}
	print(sum, "\n")
	return sum
}
