package main

import (
	"fmt"
	"sort"
)

type arr_list struct {
	arr1 []float64
	arr2 []float64
}

func Sum(arr ...float64) float64 {
	var sum float64 = 0
	for i := range arr {
		sum += arr[i]
	}
	return sum
}

func medianByEquation(arr1 []float64, arr2 []float64) float64 {
	arr1 = append(arr2, arr1...)
	sort.Float64s(arr1)
	list_len := len(arr1)
	return Sum(arr1[(list_len/2)-(list_len+1)%2:(list_len/2)-(list_len+1)%2+(list_len%2)+(2*((list_len-1)%2))]...) / (float64((list_len-1)%2) + 1)
}

func (arr arr_list) findMedian() float64 {
	if len(arr.arr1) > 1 && len(arr.arr2) > 1 {
		if len(arr.arr2) > len(arr.arr1) {
			arr.arr1, arr.arr2 = arr.arr2, arr.arr1
			return arr.findMedian()
		}
		return medianByEquation(arr.arr1, arr.arr2)
	} else if len(arr.arr1) > 1 {
		return float64(1)
	} else if len(arr.arr2) > 1 {
		return float64(2)
	} else {
		return float64(0)
	}
}

func main() {
	arrays := &arr_list{}
	arrays.arr1 = []float64{3.00, 2.00, 1.11}
	arrays.arr2 = []float64{5, 6, 7, 9}
	fmt.Println(arrays.findMedian())

}
