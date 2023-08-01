package defer_statement

import "fmt"

func A() {
	defer B() // first defered last run
	defer C()
	fmt.Println("a fonksiyonu çalıştı")
}

func B() {
	fmt.Println("b fonksiyonu çalıştı")
}

func C() {
	fmt.Println("c fonksiyonu çalıştı")
}
