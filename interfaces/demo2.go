package interfaces

import "fmt"

type Pen struct {
	size          int //(1:small 2:medium 3:large)
	penColor      string
	brandFixPrice int //(stabilo:20, faberCastell:40)
}

type Notebook struct {
	size          int    //(1:small 2:medium 3:large)
	notebookType  string //(checkered, striped)
	brandFixPrice int    // (matt:15 kraft:50)
}

type PaymentCalculator interface {
	Calculate() int
}

func (p Pen) Calculate() int {
	return p.size * p.brandFixPrice
}

func (n Notebook) Calculate() int {
	return n.size * n.brandFixPrice
}

func CalculateTotalPayment(payments []PaymentCalculator) int {
	paymentTotal := 0
	for i := 0; i < len(payments); i++ {
		paymentTotal = paymentTotal + payments[i].Calculate()
	}
	return paymentTotal
}

func Demo2() {

	payment1 := Pen{1, "grey", 20}
	payment2 := Notebook{size: 2, brandFixPrice: 15}

	payments := []PaymentCalculator{payment1, payment2}
	total := CalculateTotalPayment(payments)

	fmt.Println("Kırtasiyeye borcunuz:", total)

}
