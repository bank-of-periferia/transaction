package domain

import (
	"transaction/domain/enums"

	uuid "github.com/satori/go.uuid"
)

type Transaction struct {
	transactionId uuid.UUID
	amount float64
	merchant string
	status enums.Status
	transactionType enums.Type
	currency enums.Currency
	customerId uuid.UUID
}

func NewTransaction(amount float64, merchant string, transactionType enums.Type, customerId uuid.UUID, currency enums.Currency) *Transaction {
	return &Transaction{
		transactionId: uuid.NewV4(),
		amount: amount,
		merchant: merchant,
		status: enums.CREATED,
		transactionType: transactionType,
		currency: currency,
		customerId: customerId,
	}
}

func (t *Transaction) GetTransactionId() uuid.UUID {
	return t.transactionId
}

func (t *Transaction) GetAmount() float64 {
	return t.amount
}

