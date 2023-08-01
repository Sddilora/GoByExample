package error_handling

import (
	"fmt"
	"os"
)

func Demo1() {
	f, err := os.Open("error_handling/demo1.txt")
	//Dosya bulunursa err nil, f/osFile
	if err != nil {
		fmt.Println("Dosya bulunamadı", "; os.Open tarafından gönderilen hata:", err)
	} else {
		fmt.Println(f.Name())
	}

	d, err := os.Open("error_handling/demo2.txt") //Fonksiyonu mainde çağırdığımzı için main kendiyle aynı konumda olan dosyaları dosya yolu belirtmeden kullanabilir.
	//Dosya bulunursa err nil, f/osFile
	if err != nil {
		fmt.Println("Dosya bulunamadı", "; os.Open tarafından gönderilen hata:", err)
	} else {
		fmt.Println(d.Name())
	}
}
