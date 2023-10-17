
-- +migrate Up
CREATE SCHEMA  accounts;

-- +migrate Down
DROP SCHEMA accounts CASCADE;
