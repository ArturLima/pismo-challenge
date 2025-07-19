-- Write your migrate up statements here

CREATE TABLE IF NOT EXISTS account (
    id SERIAL PRIMARY KEY,
    document VARCHAR(250) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

---- create above / drop below ----

DROP TABLE IF EXISTS account;

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
