package repo

import (
	"context"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetExchangeNotFound(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("error when opening a stub database conn %s", err)
	}

	rows := sqlmock.NewRows([]string{"currency", "amount", "rate", "conv_amount", "currency_symbol"})

	query := regexp.QuoteMeta("SELECT currency, amount, rate, conv_amount, currency_symbol FROM exchanges WHERE currency = ? AND amount = ? AND rate = ?;")

	mock.ExpectQuery(query).WillReturnRows(rows)

	mysqlRepo := NewMysqlRepo(db)

	exchange, err := mysqlRepo.GetExchange(context.Background(), "dollar", "10", "4.50")

	assert.NoError(t, err)
	assert.Nil(t, exchange)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}
