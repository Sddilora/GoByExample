// You can immediately see that going "NORTH" and immediately "SOUTH" is not reasonable, better stay to the same place! So the task is to give to the man a simplified version of the plan. A better plan in this case is simply:package main
// a = []string{"NORTH", "SOUTH", "SOUTH", "EAST", "WEST", "NORTH", "NORTH"}
//
//	dotest(a, []string{"NORTH"})
package main

import "fmt"

func main() {
	a := []string{"NORTH", "WEST", "SOUTH", "EAST"}
	fmt.Println(DirReduc(a))
}

func DirReduc(arr []string) []string {

	// var finlDest []string
	for range arr {
		for i := 0; i < len(arr)-1; i++ {
			fmt.Println(arr[i], "i:", i, "len:", len(arr), "arr", arr)
			if arr[i] == "NORTH" && arr[i+1] == "SOUTH" || arr[i] == "SOUTH" && arr[i+1] == "NORTH" || arr[i] == "EAST" && arr[i+1] == "WEST" || arr[i] == "WEST" && arr[i+1] == "EAST" {
				copy(arr[i:], arr[i+2:])
				arr = arr[:len(arr)-2]
			}
		}
	}
	return arr
}
