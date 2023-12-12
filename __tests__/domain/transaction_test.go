package domain_test

import (
	"testing"
	"transaction/domain"
	"transaction/domain/enums"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestNewTransaction(t *testing.T) {
	transaction := domain.NewTransaction(
		100.00,
		"Starbucks",
		enums.PAYMENT,
		uuid.NewV4(),
		enums.USD,
	)


	require.NotEmpty(t, transaction.GetTransactionId())
	require.Equal(t, 100.00, transaction.GetAmount())
}