
-- +migrate Up
CREATE TABLE main.users(
    id uuid NOT NULL,
    first_name character varying(100) NOT NULL,
    last_name character varying(100) NOT NULL,
    user_name character varying(100) NOT NULL,
    email character varying(100) NOT NULL,
    identification_no character varying(20),
    address character varying(50),
    phone_number character varying(15) NOT NULL,
    token  character varying(50),
    date_created_utc timestamp without time zone DEFAULT
    timezone('utc'::text, now()),
    date_updated_utc timestamp without time zone
);
-- +migrate Down
DROP TABLE main.users;
