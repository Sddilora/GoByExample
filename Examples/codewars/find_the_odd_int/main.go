// Given an array of integers, find the one that appears an odd number of times.
// There will always be only one integer that appears an odd number of times.
package main

// func main () {
// 	var _ = Describe("Basic tests",func() {
// 		arr := [...][]int{
// 			{20,1,-1,2,-2,3,3,5,5,1,2,4,20,4,-1,-2,5},
// 			{1,1,2,-2,5,2,4,4,-1,-2,5},
// 			{20,1,1,2,2,3,3,5,5,4,20,4,5},
// 			{10},
// 			{1,1,1,1,1,1,10,1,1,1,1},
// 			{5,4,3,2,1,5,4,3,2,10,10},
// 		}
// 		sol := [...]int{5,-1,5,10,10,1}
// 		for i,v := range arr {
// 			It(fmt.Sprintf("Testing input %v",v),func() {Expect(FindOdd(v)).To(Equal(sol[i]))})}
// 	})
// }

func main() {
	a := FindOdd([]int{5, -1, -1, 55, 5, 10, 10, 1, 1})
	print(a)
}

func FindOdd(arr []int) int {

	count := make(map[int]int)
	for _, v := range arr {
		count[v]++
	}

	for k, v := range count {
		if v%2 == 1 {
			return k
		}
	}
	return 0
}
