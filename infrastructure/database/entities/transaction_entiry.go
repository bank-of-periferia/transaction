package entities

type TransactionEntity struct {
	TransactionId string `gorm:"column:transaction_id;primary_key"`
	Amount float64 `gorm:"column:amount"`
	Merchant string `gorm:"column:merchant"`
	Status string `gorm:"column:status"`
	TransactionType string `gorm:"column:transaction_type"`
	Currency string `gorm:"column:currency"`
	CustomerId string `gorm:"column:customer_id"`
}