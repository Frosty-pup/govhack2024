package structs

type User struct {
	Id           string `json:"id"`
	WeeklyIncome int    `json:"incomeweekly"`
}

type DatabaseCounter struct {
	transactionCounter int
}

type Transaction struct {
	TransactionID int    `json:"transactionid"`
	Description   string `json:"description"`
	Datetime      string `json:"datetime"`
	Isdeduction   bool   `json:"Isdeduction"`
}
