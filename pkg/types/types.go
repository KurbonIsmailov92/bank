package types

type Money int64
type Currency string

const (
	StatusOk Status = "OK"
	StatusFail Status = "FAIL"
	StatusInProgress Status = "INPROGRESS"	

)

type PAN string
type Status string
type Card struct {
	ID         int
	PAN        PAN
	Balance    Money
	Currency   Currency
	Color      string
	Name       string
	Active     bool
}

type Payment struct {
	ID       	int
	Amount   	Money
	Category 	Category
	Status		Status
}

// type of sourses of payment for choosing
type PaymentSource struct {
	Type    string // 'card'
	Number  string // PAN number
	Balance Money  // balance in dirams
}

//Категории платежей
type Category string