package main

import (
	"fmt"
)

func is_Visited(compare_val int, arr []int) bool {
	for _, val := range arr {
		if val == compare_val {
			return true
		}
	}
	return false
}

func Dijkstra_path(i_vertice int, dest int, result []int, vertices [5][5]int, visited []int) [][]int {
	if i_vertice == dest {
		return [][]int{result}
	}

	if is_Visited(i_vertice, visited) {
		return [][]int{}
	}

	var result_arr [][]int
	visited = append(visited, i_vertice)
	for to_vertice, _ := range vertices[i_vertice] {
		if !is_Visited(to_vertice, visited) {
			if vertices[i_vertice][to_vertice] != 0 {
				for _, arr := range Dijkstra_path(to_vertice, dest, append(result, vertices[i_vertice][to_vertice]), vertices, visited) {
					result_arr = append(result_arr, append([]int{}, arr...))
				}
			}
		}
	}
	return result_arr
}

// func Dijkstra_path(i_vertice int, dest int, result map[int]int, vertices [5][5]int, visited []int) []map[int]int {
// 	if i_vertice == dest {
// 		fmt.Println(result)
// 		return []map[int]int{result}
// 	}

// 	if is_Visited(i_vertice, visited) {
// 		return []map[int]int{}
// 	}

// 	var result_arr []map[int]int
// 	visited = append(visited, i_vertice)
// 	for to_vertice, val := range vertices[i_vertice] {
// 		if !is_Visited(to_vertice, visited) {
// 			if vertices[i_vertice][to_vertice] != 0 {
// 				result[to_vertice] = val
// 				var sub_path []map[int]int = Dijkstra_path(to_vertice, dest, result, vertices, visited)
// 				delete(result, to_vertice)
// 				for index, _ := range sub_path {
// 					if len(sub_path[index]) != 0 {
// 						result_arr = append(result_arr, sub_path[index])
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return result_arr
// }

func main() {
	// experiment on this
	// var vertices_and_edges = map[string]map[string]int{"A": {"B": 3, "C": 8, "D": 9}}
	// fmt.Println(vertices_and_edges)

	// using map
	// method 1
	// var vertices_and_edges = make(map[string]map[string]int)
	// vertices_and_edges["A"] = map[string]int{"B": 5, "C": 3, "E": 1}
	// vertices_and_edges["A"] = make(map[string]int)
	// vertices_and_edges["A"]["B"] = 10
	// vertices_and_edges["A"]["C"] = 9
	// vertices_and_edges["A"]["D"] = 8
	// fmt.Println(vertices_and_edges)

	// method 2
	// vertices_and_edges := map[string]map[string]int{}
	// vertices_and_edges["A"] = map[string]int{}
	// vertices_and_edges["A"]["B"] = 6
	// vertices_and_edges["A"]["C"] = 0
	// vertices_and_edges["A"]["D"] = 1
	// vertices_and_edges["A"]["E"] = 0
	// vertices_and_edges["B"] = map[string]int{}
	// vertices_and_edges["B"]["A"] = 6
	// vertices_and_edges["B"]["C"] = 5
	// vertices_and_edges["B"]["D"] = 2
	// vertices_and_edges["B"]["E"] = 2
	// vertices_and_edges["C"] = map[string]int{}
	// vertices_and_edges["C"]["A"] = 0
	// vertices_and_edges["C"]["B"] = 5
	// vertices_and_edges["C"]["D"] = 0
	// vertices_and_edges["C"]["E"] = 5
	// vertices_and_edges["D"] = map[string]int{}
	// vertices_and_edges["D"]["A"] = 1
	// vertices_and_edges["D"]["B"] = 2
	// vertices_and_edges["D"]["C"] = 0
	// vertices_and_edges["D"]["E"] = 1
	// vertices_and_edges["E"] = map[string]int{}
	// vertices_and_edges["E"]["A"] = 0
	// vertices_and_edges["E"]["B"] = 2
	// vertices_and_edges["E"]["C"] = 5
	// vertices_and_edges["E"]["D"] = 1
	// fmt.Println(vertices_and_edges)

	// var vertices []string
	// static input
	// vertices = {"A", "B", "C", "D", "E"}

	// dynamic input
	// type 1
	// for i := 0; i < 5; i++ {
	// 	fmt.Println("Enter val")
	// 	reader := bufio.NewReader(os.Stdin)
	// 	val, err := reader.ReadString('\n')
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	vertices = append(vertices, val)
	// }
	// fmt.Print(vertices)

	// type 2
	// var vertices1 []string
	// for i := 0; i < 5; i++ {
	// 	var val string
	// 	fmt.Scanln(&val)
	// 	vertices1 = append(vertices1, val)
	// }

	var vertices3 = [5][5]int{
		{0, 6, 0, 1, 0},
		{6, 0, 5, 2, 2},
		{0, 5, 0, 5, 0},
		{1, 2, 0, 0, 1},
		{0, 2, 5, 1, 0},
	}
	var (
		start = 0
		dest  = 2
	)

	/* can comment and uncomment any of the
	below two type to get different output type. */
	// result := map[int]int{}
	result := []int{}

	all_possible_paths := Dijkstra_path(start, dest, result, vertices3, []int{})
	for _, val := range all_possible_paths {
		fmt.Println(val)
	}
	// find the min in this array

}
