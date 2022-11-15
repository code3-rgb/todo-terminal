package dummy

import (
	"fmt"
	"time"
)

var Name string = "John"

var Array []string = []string{"John", "Peter", "Mark"}

func Print() int {
	fmt.Println("Hello how do you do!")
	return 1
}

type constraint interface {
	int | float64
}

func Display[T constraint](x T, y T) T {
	return x + y
}

func BubbleSort(arr []int) []int {

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
	return arr

}

func Recur(count int) {
	if count == 0 {
		return
	}
	time.Sleep(1 * time.Second)
	fmt.Println("Hello, ", count)
	count -= 1
	Recur(count)

}

func BinarySearch(arr []int64, start int64, end int64, target int64) string {
	if start > end {
		return "Not found "
	}

	midIndx := (start + end) / 2

	if arr[midIndx] == target {
		return fmt.Sprintf("Found at %d", midIndx)

	}

	if arr[midIndx] > target {
		return BinarySearch(arr, start, midIndx-1, target)
	}
	return BinarySearch(arr, midIndx+1, end, target)

}
