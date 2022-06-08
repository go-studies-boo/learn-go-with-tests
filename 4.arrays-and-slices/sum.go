package arraysandslices

func Sum(numbers []int) int {
	sum := 0
	// https://go.dev/doc/effective_go#blank
	for _, number := range numbers {
		sum += number
	}

	return sum
}

// As mentioned, slices have a capacity. 
// If you have a slice with a capacity of 2 and try to do mySlice[10] = 1 
// you will get a runtime error.
//
// However, you can use the append function which takes a slice and a new value, 
// then returns a new slice with all the items in it.
//
// In this implementation, we are worrying less about capacity.
func SumAll(numbersToSum... []int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers)) // here we doesn't have problem with capacity
	}

	return sums
}

func SumAllTails(numbersToSum... []int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			// Slices can be sliced! The syntax is slice[low:high]. 
			// If you omit the value on one of the sides of the : it captures everything to that side of it. 
			// In our case, we are saying "take from 1 to the end" with numbers[1:]
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}