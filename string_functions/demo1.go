package string_functions

import (
	"fmt"
	"strings"
	// if I wrote (s  "strings") The `s` would be used as strings. ex. s.count instead strings.count
	//and s called as "alias" (tr:takma ad)
)

func Demo1() {
	name := "Dilara"
	fmt.Println(strings.Count(name, "a"))    //->2
	fmt.Println(strings.Count(name, ""))     //->7
	fmt.Println(strings.Count(name, "d"))    //->0   //case sensitive so prints result will be zero
	fmt.Println(strings.Contains(name, "i")) //->true //turns bool
	fmt.Println(strings.Index(name, "a"))    //->3

	result := strings.Index(name, "a") //-> the string contains `a`
	if result != -1 {
		fmt.Println("the string contains `a` ")
	} else {
		fmt.Println("the string doesn't contain `a` ")
	}

	fmt.Println(strings.ToLower(name)) //->dilara
	fmt.Println(strings.ToUpper(name)) //->DILARA
}
