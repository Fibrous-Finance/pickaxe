-- name: CreatePool :one
INSERT INTO pools_v2 (
  address,
  amm_id,
  token_a,
  token_b
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: GetPoolByAddress :one
SELECT * FROM pools_v2
WHERE address = $1 LIMIT 1;

-- name: GetPoolByAddressExtra :one
SELECT * FROM pools_v2
WHERE address = $1 AND extra_data=$2 LIMIT 1;

-- name: GetPoolsByPair :many
SELECT * FROM pools_v2
WHERE token_a = $1 AND token_b = $2
ORDER BY amm_id;

-- name: GetPoolsByToken :many
SELECT * FROM pools_v2
WHERE token_a = $1 OR token_b = $1
ORDER BY amm_id;

-- name: GetPoolsByAmm :many
SELECT * FROM pools_v2
WHERE amm_id = $1
ORDER BY address;

-- name: GetAllPools :many
SELECT * FROM pools_v2
ORDER BY address;

-- name: GetAllPoolsWithoutKeys :many
SELECT * FROM pools_v2
WHERE amm_id IN 
(SELECT amm_id FROM amms WHERE key = '')
ORDER BY address;


-- name: UpdatePoolReserves :one
UPDATE pools_v2
SET reserve_a = $2, reserve_b = $3, last_block = $4, last_updated = NOW()
WHERE pool_id = $1
RETURNING *;

-- name: UpdatePoolTV :one
UPDATE pools_v2
SET total_value = $2
WHERE pool_id = $1
RETURNING *;

-- name: UpdatePoolFee :one
UPDATE pools_v2
SET fee = $2
WHERE pool_id = $1
RETURNING *;

-- name: UpdatePoolExtraData :one
UPDATE pools_v2
SET extra_data = $2
WHERE pool_id = $1
RETURNING *;
 
-- name: DeletePool :exec
DELETE FROM pools_v2
WHERE pool_id = $1;