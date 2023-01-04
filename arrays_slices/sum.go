package main

func Sum(numbers []int) (result int) {
	for _, number := range numbers {
		result += number
	}
	return
}

func SumAll(numbers ...[]int) (sums []int) {
	for _, slice := range numbers {
		sums = append(sums, Sum(slice))
	}
	return
}

func SumAllTails(numbers ...[]int) (sums []int) {
	for _, slice := range numbers {
		if len(slice) == 0 {
			sums = append(sums, 0)
		} else {
			tail := slice[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return
}

func main() {

}
