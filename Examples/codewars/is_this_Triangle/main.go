// Implement a function that accepts 3 integer values a, b, c. The function should return true if a triangle can be built with the sides of given length and false in any other case.
// (In this case, all triangles must have surface greater than 0 to be accepted).
package main

func main() {
	print("This sentence is 'Those lines can be triangle':", IsTriangle(13, 4, 5))
}

func IsTriangle(a, b, c int) bool {
	if a == 0 || b == 0 || c == 0 {
		return false
	}
	if a+b > c && a+c > b && b+c > a {
		return true
	}
	return false
}
