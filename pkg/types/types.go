package types

type Money int64

type Currency string

type Category string

const (
	TJS Currency = "TJS"
	USD Currency = "USD"
	RUB Currency = "RUB"
	EUR Currency = "EUR"
)

type PAN string

type Payment struct {
	ID 			int
	Amount 		Money
	Category	Category
}

type Card struct {
	ID			int
	PAN			PAN
	Balance		Money
	Currency	Currency
	Color		string
	Name		string
	Active		bool
	MinBalance 	Money
}