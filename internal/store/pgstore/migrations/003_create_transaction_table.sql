-- Write your migrate up statements here

CREATE TABLE IF NOT EXISTS transaction(
  id SERIAL PRIMARY KEY,
  account_id INTEGER NOT NULL REFERENCES account(id),
  operation_id INTEGER NOT NULL REFERENCES operations_types(operation_type_id),
  amount NUMERIC(15,2) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

---- create above / drop below ----

DROP TABLE IF EXISTS transaction;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
