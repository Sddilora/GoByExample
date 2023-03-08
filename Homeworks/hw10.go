package hw

import "fmt"

func hw10() {

	var a []int
	var v []int
	var b int

	var h int
	var n int
	var m int

	fmt.Println("Gireceğiniz Dizinin Eleman Sayısını Giriniz")
	fmt.Scanf("%d\n", &h)
	fmt.Println("Dizinizi arasına virgül koyarak ya da her sayıdan sonra entera basarak Giriniz")
	for i := 0; i < h; i++ {

		fmt.Scanf("%d\n", &n)
		a = append(a, n)
	}

	for c := 0; c < len(a); c++ {

		for i := 0; i < len(v); i++ {

			if a[c] == v[i] {
				a[c] = 0
			}
		}

		if a[c] != 0 {

			for i := 0; i < len(a); i++ {

				if a[c] == a[i] {
					b++
					m = a[c]
					v = append(v, m)
				}

			}
			fmt.Println(b, "adet", a[c])
			b = 0
		}
	}
}

/*	for c := 0; c < len(a); c++ {

	v = append(v, a[c])

	for i := 0; i < len(a); i++ {

		if a[c] == a[i] {
			b++

		}

		fmt.Println(b, "tane", v[c])
		b = 0
	}

}*/
