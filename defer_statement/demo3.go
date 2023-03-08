package defer_statement

import "fmt"

type product struct {
	productName string
	unitPrice   int
}

func (p product) save() {
	fmt.Println("ürün kaydedildi:", p.productName)
	defer Log() // üstteki kod hata verse de log çalışsın istiyorum onun iin log gibi altyapıları defer ifadesiyle kullanırız.
}

func Log() {
	fmt.Println("Loglandı")
}

func Demo3() {
	p := product{productName: "Laptop", unitPrice: 50000}
	//p.save()
	defer p.save() //deferi sıraya eklediğim zaman o anki değeriyle sıraya alırım yani product olarak Laptobu alıp gidip sırasına çekilir en son laptobu verir
	p = product{productName: "Mouse", unitPrice: 100}
	fmt.Println("İşlem başarılı")
}
