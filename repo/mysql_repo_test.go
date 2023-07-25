package repo

import (
	"context"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/giovanisilqueirasantos/coinconv/domain"
	"github.com/stretchr/testify/assert"
)

func TestGetExchangeNotFound(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("error when opening a stub database conn %s", err)
	}

	rows := sqlmock.NewRows([]string{"conv_amount"})

	query := regexp.QuoteMeta("SELECT conv_amount FROM exchanges WHERE currency_from = ? AND currency_to = ? AND amount = ? AND rate = ?;")

	mock.ExpectQuery(query).WillReturnRows(rows)

	mysqlRepo := NewMysqlRepo(db)

	real := domain.Currency{}
	err = real.New("real")
	assert.Nil(t, err)

	dollar := domain.Currency{}
	err = dollar.New("dollar")
	assert.Nil(t, err)

	amount := domain.Amount{}
	err = amount.New("10")
	assert.Nil(t, err)

	rate := domain.Amount{}
	err = rate.New("4.50")
	assert.Nil(t, err)

	convAmount, err := mysqlRepo.GetExchange(context.Background(), real, dollar, amount, rate)

	assert.NoError(t, err)
	assert.Nil(t, convAmount)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}

func TestGetExchangeError(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("error when opening a stub database conn %s", err)
	}

	query := regexp.QuoteMeta("SELECT conv_amount FROM exchanges WHERE currency_from = ? AND currency_to = ? AND amount = ? AND rate = ?;")

	mock.ExpectQuery(query).WillReturnError(errors.New("error message"))

	mysqlRepo := NewMysqlRepo(db)

	real := domain.Currency{}
	err = real.New("real")
	assert.Nil(t, err)

	dollar := domain.Currency{}
	err = dollar.New("dollar")
	assert.Nil(t, err)

	amount := domain.Amount{}
	err = amount.New("10")
	assert.Nil(t, err)

	rate := domain.Amount{}
	err = rate.New("4.50")
	assert.Nil(t, err)

	_, err = mysqlRepo.GetExchange(context.Background(), real, dollar, amount, rate)

	assert.Error(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}

func TestGetExchange(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("error when opening a stub database conn %s", err)
	}

	rows := sqlmock.NewRows([]string{"conv_amount"}).AddRow("45")

	query := regexp.QuoteMeta("SELECT conv_amount FROM exchanges WHERE currency_from = ? AND currency_to = ? AND amount = ? AND rate = ?;")

	mock.ExpectQuery(query).WillReturnRows(rows)

	mysqlRepo := NewMysqlRepo(db)

	real := domain.Currency{}
	err = real.New("real")
	assert.Nil(t, err)

	dollar := domain.Currency{}
	err = dollar.New("dollar")
	assert.Nil(t, err)

	amount := domain.Amount{}
	err = amount.New("10")
	assert.Nil(t, err)

	rate := domain.Amount{}
	err = rate.New("4.50")
	assert.Nil(t, err)

	convAmount, err := mysqlRepo.GetExchange(context.Background(), real, dollar, amount, rate)

	assert.NoError(t, err)
	assert.Equal(t, float64(45), convAmount.Value)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}
