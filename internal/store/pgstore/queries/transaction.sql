-- name: CreateTransaction :one
INSERT INTO transaction(
  account_id,
  operation_id,
  amount
)
values(
  $1, -- account_id
  $2, -- operation_id
  $3  -- amount
)
RETURNING *;