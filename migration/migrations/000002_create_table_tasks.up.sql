-- migrate up

BEGIN;

CREATE TYPE status AS ENUM (
    'INCOMPLETE',
    'COMPLETE'
);

CREATE TABLE IF NOT EXISTS tasks (
    id          uuid primary key,
    creator_id  uuid not null,
    title       text not null,
    detail      text,
    status      status not null default 'INCOMPLETE',
    created_at  timestamp(6) with time zone not null default current_timestamp,
    updated_at  timestamp(6) with time zone not null default current_timestamp,
    deadline    timestamp(6) with time zone
);

CREATE OR REPLACE FUNCTION set_update_time_on_tasks()
RETURNS trigger AS
$BODY$
BEGIN
    new.updated_at := current_timestamp;
    return new;
END;
$BODY$
LANGUAGE plpgsql;

CREATE TRIGGER update_trigger_on_tasks
BEFORE UPDATE ON tasks 
FOR EACH ROW
EXECUTE PROCEDURE set_update_time_on_tasks();

COMMIT;
