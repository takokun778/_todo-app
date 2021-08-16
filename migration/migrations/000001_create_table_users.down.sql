-- migrate down

BEGIN;

DROP TRIGGER IF EXISTS update_trigger_on_users ON users;

DROP FUNCTION IF EXISTS set_update_time_on_users();

DROP TABLE IF EXISTS users;

COMMIT;
