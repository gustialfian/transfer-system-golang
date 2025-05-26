CREATE TABLE accounts (
    account_id      bigint PRIMARY KEY,
    initial_balance bigint NOT NULL,
    scale_balance   smallint NOT NULL,
    created_at      timestamp with time zone NOT NULL,
    updated_at      timestamp with time zone NOT NULL
);

CREATE TABLE transactions (
    transaction_id          bigserial PRIMARY KEY,
    source_account_id       bigint NOT NULL,
    destination_account_id  bigint NOT NULL,
    amount                  bigint NOT NULL,
    scale_amount            smallint NOT NULL,
    created_at              timestamp with time zone NOT NULL,
    updated_at              timestamp with time zone NOT NULL
);