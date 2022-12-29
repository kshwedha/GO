package main

import "fmt"

func reverse(word string) (reversed_word string) {
	reversed_word = ""
	for i := len(word) - 1; i >= 0; i-- {
		reversed_word = reversed_word + string(word[i])
	}
	return reversed_word
}

func isPalin(arr []string, result [][]int, a int, b int) [][]int {
	if arr[a]+arr[b] == reverse(arr[a]+arr[b]) {
		result = append(result, []int{a, b})
	}
	return result
}

// func isPalin(a string, b string) bool {
// 	return a+b == reverse(a+b)
// 	return false
// }

func main() {
	arr := []string{"code", "edoc", "da", "d"}
	var result [][]int
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			// if isPalin(arr[i], arr[j]) {
			// 	fmt.Println(i, j)
			// }
			// if isPalin(arr[j], arr[i]) {
			// 	fmt.Println(j, i)
			// }
			result = isPalin(arr, result, i, j)
			result = isPalin(arr, result, j, i)
		}
	}
	fmt.Println(result)
}
