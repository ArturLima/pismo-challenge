-- Write your migrate up statements here
CREATE TABLE IF NOT EXISTS operations_types (
    operation_type_id SERIAL PRIMARY KEY,
    description TEXT NOT NULL
);

INSERT INTO operations_types (operation_type_id, description) VALUES
    (1, 'Normal Purchase'),
    (2, 'Purchase with installments'),
    (3, 'Withdrawal'),
    (4, 'Credit Voucher');

---- create above / drop below ----

DROP TABLE IF EXISTS operations_types;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
