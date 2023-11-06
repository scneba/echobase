
-- +migrate Up
CREATE SCHEMA  main;

-- +migrate Down
DROP SCHEMA main CASCADE;
