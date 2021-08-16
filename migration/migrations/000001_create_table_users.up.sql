-- migrate up

BEGIN;

CREATE TABLE IF NOT EXISTS users(
    id          uuid primary key,
    first_name  text not null,
    last_name   text not null,
    birthday    date not null,
    created_at  timestamp(6) with time zone not null default current_timestamp,
    updated_at  timestamp(6) with time zone not null default current_timestamp
);

CREATE OR REPLACE FUNCTION set_update_time_on_users()
RETURNS trigger AS
$BODY$
BEGIN
    new.updated_at := current_timestamp;
    return new;
END;
$BODY$
LANGUAGE plpgsql;

CREATE TRIGGER update_trigger_on_users
BEFORE UPDATE ON users 
FOR EACH ROW
EXECUTE PROCEDURE set_update_time_on_users();

COMMIT;
