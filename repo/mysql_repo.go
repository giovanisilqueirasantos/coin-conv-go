package repo

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/giovanisilqueirasantos/coin-conv-go/domain"
)

type mysqlRepo struct {
	Conn *sql.DB
}

func NewMysqlRepo(conn *sql.DB) domain.Repo {
	return &mysqlRepo{Conn: conn}
}

func (r mysqlRepo) GetExchange(ctx context.Context, currencyFrom, currencyTo domain.Currency, amountQuant, rate domain.Amount) (*domain.Amount, error) {
	query := `SELECT conv_amount FROM exchanges WHERE currency_from = ? AND currency_to = ? AND amount = ? AND rate = ?;`

	row := r.Conn.QueryRowContext(ctx, query, currencyFrom.Text, currencyTo.Text, fmt.Sprintf("%f", amountQuant.Value), fmt.Sprintf("%f", rate.Value))

	var convAmount string

	if err := row.Scan(&convAmount); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	res := domain.Amount{}

	err := res.New(convAmount)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (r mysqlRepo) StoreExchange(ctx context.Context, currencyFrom, currencyTo domain.Currency, amountQuant, rate, convAmount domain.Amount) error {
	query := `INSERT INTO exchanges (currency_from, currency_to, amount, rate, conv_amount) VALUES (?, ?, ?, ?, ?);`

	stmt, err := r.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	exec, err := stmt.ExecContext(ctx, currencyFrom.Text, currencyTo.Text, fmt.Sprintf("%f", amountQuant.Value), fmt.Sprintf("%f", rate.Value), fmt.Sprintf("%f", convAmount.Value))
	if err != nil {
		return err
	}

	affect, err := exec.RowsAffected()
	if err != nil {
		return err
	}

	if affect != 1 {
		return fmt.Errorf("error trying to store exchange with total rows affcted: %d", affect)
	}

	return nil
}
