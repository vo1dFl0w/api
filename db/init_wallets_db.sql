CREATE TABLE IF NOT EXISTS wallets (
    uuid UUID NOT NULL DEFAULT gen_random_uuid(),
    account BIGINT NOT NULL DEFAULT 0
);

INSERT INTO wallets (uuid, account) VALUES ('f81d4fae-7dec-11d0-a765-00a0c91e6bf6', 10000) RETURNING uuid;