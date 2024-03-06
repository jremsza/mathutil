package mathutil

import (
	"math"
	"sort"
)

type Num interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func Average[T Num](numbers []T) T {
	//sort the numbers in ascending order
	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j]
	})

	//Perform a 5% trim
	trim := int(math.Ceil(float64(len(numbers)) * 0.05))
	numbers = numbers[trim : len(numbers)-trim]

	//calc the average
	total := 0.0
	for _, number := range numbers {
		total += float64(number)
	}

	return T(total / float64(len(numbers))) // Convert the result back to type T
}
