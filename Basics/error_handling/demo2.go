package error_handling

import (
	"fmt"
	"os"
)

func Demo2() {
	f, err := os.Open("demo1.txt") // Dosya bulunamadı demo1.txt ( dosyası main.go dosyasının bulunduğu konumda olmalıdır. )

	if err != nil {
		if pErr, ok := err.(*os.PathError); ok {
			fmt.Println("Dosya bulunamadı", pErr.Path)
			return
		} else {
			fmt.Println("Dosya bulunamadı")
			return
		}
	} else {
		fmt.Println(f.Name())
		return
	}
}
