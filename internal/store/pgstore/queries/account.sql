-- name: CreateAccount :one

INSERT INTO account(
  document
)
values (
  $1
)
RETURNING *;


-- name: GetAccountById :one

SELECT * FROM account
WHERE id = $1;