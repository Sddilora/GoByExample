package interfaces

type Mortgage struct { //mortgage: ipotek
	creditPaymentTotal float64
	address            string
	rate               float64
}

type Car struct {
	creditPaymentTotal float64
	carInfo            string
	rate               float64
}

type CreditCalculator interface {
	Calculate() float64
}

//structlar için interfaceimizi implamente edeceğiz

// Şimdi bir metod yazalım
func (m Mortgage) Calculate() float64 {
	return m.creditPaymentTotal * m.rate / 12 //12'ye yıllık ödeme varsayıp aylıködemeyi hesaplayalım diye böldük
}

// Bir tane de araba kredisi için yazalım
func (c Car) Calculate() float64 {
	return c.creditPaymentTotal * c.rate / 12
}

func CalculateMonthlyPayment(credits []CreditCalculator) { //birden fazla kredii olabillir diye array olarak interfacei gönderdim
	paymentTotal := 0.0

	for i := 0; i < len(credits); i++ {
		paymentTotal = paymentTotal + credits[i].Calculate()
	}
}

func Demo1() {
	credit1 := Mortgage{creditPaymentTotal: 100000, rate: 10, address: "Ankara"}
	credit2 := Mortgage{creditPaymentTotal: 50000, rate: 5, address: "İstanbul"}
	credit3 := Car{creditPaymentTotal: 700000, rate: 15, carInfo: "vosvos"}

	credits := []CreditCalculator{credit1, credit2, credit3}
	total := CalculateMonthlyPayment(credits)
}
