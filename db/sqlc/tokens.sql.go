// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: tokens.sql

package db

import (
	"context"
)

const createToken = `-- name: CreateToken :one
INSERT INTO tokens (
  address,
  name,
  symbol,
  decimals,
  ticker
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING address, name, symbol, decimals, base, native, ticker, price, created_at
`

type CreateTokenParams struct {
	Address  string `json:"address"`
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Decimals int32  `json:"decimals"`
	Ticker   string `json:"ticker"`
}

func (q *Queries) CreateToken(ctx context.Context, arg CreateTokenParams) (Token, error) {
	row := q.db.QueryRowContext(ctx, createToken,
		arg.Address,
		arg.Name,
		arg.Symbol,
		arg.Decimals,
		arg.Ticker,
	)
	var i Token
	err := row.Scan(
		&i.Address,
		&i.Name,
		&i.Symbol,
		&i.Decimals,
		&i.Base,
		&i.Native,
		&i.Ticker,
		&i.Price,
		&i.CreatedAt,
	)
	return i, err
}

const deleteToken = `-- name: DeleteToken :exec
DELETE FROM tokens
WHERE address = $1
`

func (q *Queries) DeleteToken(ctx context.Context, address string) error {
	_, err := q.db.ExecContext(ctx, deleteToken, address)
	return err
}

const getAllTokens = `-- name: GetAllTokens :many
SELECT address, name, symbol, decimals, base, native, ticker, price, created_at FROM tokens
ORDER BY address
`

func (q *Queries) GetAllTokens(ctx context.Context) ([]Token, error) {
	rows, err := q.db.QueryContext(ctx, getAllTokens)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Token{}
	for rows.Next() {
		var i Token
		if err := rows.Scan(
			&i.Address,
			&i.Name,
			&i.Symbol,
			&i.Decimals,
			&i.Base,
			&i.Native,
			&i.Ticker,
			&i.Price,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllTokensWithTickers = `-- name: GetAllTokensWithTickers :many
SELECT address, name, symbol, decimals, base, native, ticker, price, created_at FROM tokens
WHERE price IS NOT NULL
ORDER BY address
`

func (q *Queries) GetAllTokensWithTickers(ctx context.Context) ([]Token, error) {
	rows, err := q.db.QueryContext(ctx, getAllTokensWithTickers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Token{}
	for rows.Next() {
		var i Token
		if err := rows.Scan(
			&i.Address,
			&i.Name,
			&i.Symbol,
			&i.Decimals,
			&i.Base,
			&i.Native,
			&i.Ticker,
			&i.Price,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getBaseTokens = `-- name: GetBaseTokens :many
SELECT address, name, symbol, decimals, base, native, ticker, price, created_at FROM tokens
WHERE base = true
ORDER BY name
`

func (q *Queries) GetBaseTokens(ctx context.Context) ([]Token, error) {
	rows, err := q.db.QueryContext(ctx, getBaseTokens)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Token{}
	for rows.Next() {
		var i Token
		if err := rows.Scan(
			&i.Address,
			&i.Name,
			&i.Symbol,
			&i.Decimals,
			&i.Base,
			&i.Native,
			&i.Ticker,
			&i.Price,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getNativeTokens = `-- name: GetNativeTokens :many
SELECT address, name, symbol, decimals, base, native, ticker, price, created_at FROM tokens
WHERE native = true
ORDER BY name
`

func (q *Queries) GetNativeTokens(ctx context.Context) ([]Token, error) {
	rows, err := q.db.QueryContext(ctx, getNativeTokens)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Token{}
	for rows.Next() {
		var i Token
		if err := rows.Scan(
			&i.Address,
			&i.Name,
			&i.Symbol,
			&i.Decimals,
			&i.Base,
			&i.Native,
			&i.Ticker,
			&i.Price,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTokenAPriceByPool = `-- name: GetTokenAPriceByPool :one
SELECT price FROM tokens
WHERE address = (SELECT token_a FROM pools_v2 WHERE pool_id = $1)
`

func (q *Queries) GetTokenAPriceByPool(ctx context.Context, poolID int64) (string, error) {
	row := q.db.QueryRowContext(ctx, getTokenAPriceByPool, poolID)
	var price string
	err := row.Scan(&price)
	return price, err
}

const getTokenBPriceByPool = `-- name: GetTokenBPriceByPool :one
SELECT price FROM tokens
WHERE address = (SELECT token_b FROM pools_v2 WHERE pool_id = $1)
`

func (q *Queries) GetTokenBPriceByPool(ctx context.Context, poolID int64) (string, error) {
	row := q.db.QueryRowContext(ctx, getTokenBPriceByPool, poolID)
	var price string
	err := row.Scan(&price)
	return price, err
}

const getTokenByAddress = `-- name: GetTokenByAddress :one
SELECT address, name, symbol, decimals, base, native, ticker, price, created_at FROM tokens
WHERE address = $1 LIMIT 1
`

func (q *Queries) GetTokenByAddress(ctx context.Context, address string) (Token, error) {
	row := q.db.QueryRowContext(ctx, getTokenByAddress, address)
	var i Token
	err := row.Scan(
		&i.Address,
		&i.Name,
		&i.Symbol,
		&i.Decimals,
		&i.Base,
		&i.Native,
		&i.Ticker,
		&i.Price,
		&i.CreatedAt,
	)
	return i, err
}

const getTokenBySymbol = `-- name: GetTokenBySymbol :one
SELECT address, name, symbol, decimals, base, native, ticker, price, created_at FROM tokens
WHERE symbol = $1 LIMIT 1
`

func (q *Queries) GetTokenBySymbol(ctx context.Context, symbol string) (Token, error) {
	row := q.db.QueryRowContext(ctx, getTokenBySymbol, symbol)
	var i Token
	err := row.Scan(
		&i.Address,
		&i.Name,
		&i.Symbol,
		&i.Decimals,
		&i.Base,
		&i.Native,
		&i.Ticker,
		&i.Price,
		&i.CreatedAt,
	)
	return i, err
}

const updateBaseNativeStatus = `-- name: UpdateBaseNativeStatus :one
UPDATE tokens
SET base = $2, native = $3
WHERE address = $1
RETURNING address, name, symbol, decimals, base, native, ticker, price, created_at
`

type UpdateBaseNativeStatusParams struct {
	Address string `json:"address"`
	Base    bool   `json:"base"`
	Native  bool   `json:"native"`
}

func (q *Queries) UpdateBaseNativeStatus(ctx context.Context, arg UpdateBaseNativeStatusParams) (Token, error) {
	row := q.db.QueryRowContext(ctx, updateBaseNativeStatus, arg.Address, arg.Base, arg.Native)
	var i Token
	err := row.Scan(
		&i.Address,
		&i.Name,
		&i.Symbol,
		&i.Decimals,
		&i.Base,
		&i.Native,
		&i.Ticker,
		&i.Price,
		&i.CreatedAt,
	)
	return i, err
}

const updatePrice = `-- name: UpdatePrice :one
UPDATE tokens
SET price = $2
WHERE address = $1
RETURNING address, name, symbol, decimals, base, native, ticker, price, created_at
`

type UpdatePriceParams struct {
	Address string `json:"address"`
	Price   string `json:"price"`
}

func (q *Queries) UpdatePrice(ctx context.Context, arg UpdatePriceParams) (Token, error) {
	row := q.db.QueryRowContext(ctx, updatePrice, arg.Address, arg.Price)
	var i Token
	err := row.Scan(
		&i.Address,
		&i.Name,
		&i.Symbol,
		&i.Decimals,
		&i.Base,
		&i.Native,
		&i.Ticker,
		&i.Price,
		&i.CreatedAt,
	)
	return i, err
}

const updateTicker = `-- name: UpdateTicker :one
UPDATE tokens
SET ticker = $2
WHERE address = $1
RETURNING address, name, symbol, decimals, base, native, ticker, price, created_at
`

type UpdateTickerParams struct {
	Address string `json:"address"`
	Ticker  string `json:"ticker"`
}

func (q *Queries) UpdateTicker(ctx context.Context, arg UpdateTickerParams) (Token, error) {
	row := q.db.QueryRowContext(ctx, updateTicker, arg.Address, arg.Ticker)
	var i Token
	err := row.Scan(
		&i.Address,
		&i.Name,
		&i.Symbol,
		&i.Decimals,
		&i.Base,
		&i.Native,
		&i.Ticker,
		&i.Price,
		&i.CreatedAt,
	)
	return i, err
}
