-- migrate down

BEGIN;

DROP TRIGGER IF EXISTS update_trigger_on_tasks ON tasks;

DROP FUNCTION IF EXISTS set_update_time_on_tasks();

DROP TABLE IF EXISTS tasks;

DROP TYPE status;

COMMIT;
