package types

type Money int64

type Currency string

type Category string

type Status string

const (
	TJS Currency = "TJS"
	USD Currency = "USD"
	RUB Currency = "RUB"
	EUR Currency = "EUR"
)

const (
	StatusOk 			Status = "OK"
	StatusFail 			Status = "FAIL"
	StatusInProgress 	Status = "INPROGRESS"
)

type PAN string

type Payment struct {
	ID 			int
	Amount 		Money
	Category	Category
	Status		Status
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