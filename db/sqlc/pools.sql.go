// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: pools.sql

package db

import (
	"context"
)

const createPool = `-- name: CreatePool :one
INSERT INTO pools (
  address,
  amm_id,
  token_a,
  token_b,
  reserve_a,
  reserve_b
) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING address, amm_id, token_a, token_b, reserve_a, reserve_b, total_value, last_updated
`

type CreatePoolParams struct {
	Address  string `json:"address"`
	AmmID    int64  `json:"amm_id"`
	TokenA   string `json:"token_a"`
	TokenB   string `json:"token_b"`
	ReserveA string `json:"reserve_a"`
	ReserveB string `json:"reserve_b"`
}

func (q *Queries) CreatePool(ctx context.Context, arg CreatePoolParams) (Pool, error) {
	row := q.db.QueryRowContext(ctx, createPool,
		arg.Address,
		arg.AmmID,
		arg.TokenA,
		arg.TokenB,
		arg.ReserveA,
		arg.ReserveB,
	)
	var i Pool
	err := row.Scan(
		&i.Address,
		&i.AmmID,
		&i.TokenA,
		&i.TokenB,
		&i.ReserveA,
		&i.ReserveB,
		&i.TotalValue,
		&i.LastUpdated,
	)
	return i, err
}

const deletePool = `-- name: DeletePool :exec
DELETE FROM pools
WHERE address = $1
`

func (q *Queries) DeletePool(ctx context.Context, address string) error {
	_, err := q.db.ExecContext(ctx, deletePool, address)
	return err
}

const getPoolByAddress = `-- name: GetPoolByAddress :one
SELECT address, amm_id, token_a, token_b, reserve_a, reserve_b, total_value, last_updated FROM pools
WHERE address = $1 LIMIT 1
`

func (q *Queries) GetPoolByAddress(ctx context.Context, address string) (Pool, error) {
	row := q.db.QueryRowContext(ctx, getPoolByAddress, address)
	var i Pool
	err := row.Scan(
		&i.Address,
		&i.AmmID,
		&i.TokenA,
		&i.TokenB,
		&i.ReserveA,
		&i.ReserveB,
		&i.TotalValue,
		&i.LastUpdated,
	)
	return i, err
}

const getPoolsByAmm = `-- name: GetPoolsByAmm :many
SELECT address, amm_id, token_a, token_b, reserve_a, reserve_b, total_value, last_updated FROM pools
WHERE amm_id = $1
ORDER BY address
`

func (q *Queries) GetPoolsByAmm(ctx context.Context, ammID int64) ([]Pool, error) {
	rows, err := q.db.QueryContext(ctx, getPoolsByAmm, ammID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Pool{}
	for rows.Next() {
		var i Pool
		if err := rows.Scan(
			&i.Address,
			&i.AmmID,
			&i.TokenA,
			&i.TokenB,
			&i.ReserveA,
			&i.ReserveB,
			&i.TotalValue,
			&i.LastUpdated,
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

const getPoolsByPair = `-- name: GetPoolsByPair :many
SELECT address, amm_id, token_a, token_b, reserve_a, reserve_b, total_value, last_updated FROM pools
WHERE token_a = $1 AND token_b = $2
ORDER BY amm_id
`

type GetPoolsByPairParams struct {
	TokenA string `json:"token_a"`
	TokenB string `json:"token_b"`
}

func (q *Queries) GetPoolsByPair(ctx context.Context, arg GetPoolsByPairParams) ([]Pool, error) {
	rows, err := q.db.QueryContext(ctx, getPoolsByPair, arg.TokenA, arg.TokenB)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Pool{}
	for rows.Next() {
		var i Pool
		if err := rows.Scan(
			&i.Address,
			&i.AmmID,
			&i.TokenA,
			&i.TokenB,
			&i.ReserveA,
			&i.ReserveB,
			&i.TotalValue,
			&i.LastUpdated,
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

const getPoolsByToken = `-- name: GetPoolsByToken :many
SELECT address, amm_id, token_a, token_b, reserve_a, reserve_b, total_value, last_updated FROM pools
WHERE token_a = $1 OR token_b = $1
ORDER BY amm_id
`

func (q *Queries) GetPoolsByToken(ctx context.Context, tokenA string) ([]Pool, error) {
	rows, err := q.db.QueryContext(ctx, getPoolsByToken, tokenA)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Pool{}
	for rows.Next() {
		var i Pool
		if err := rows.Scan(
			&i.Address,
			&i.AmmID,
			&i.TokenA,
			&i.TokenB,
			&i.ReserveA,
			&i.ReserveB,
			&i.TotalValue,
			&i.LastUpdated,
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
