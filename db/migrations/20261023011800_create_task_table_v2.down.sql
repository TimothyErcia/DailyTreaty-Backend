DROP TABLE IF EXISTS Task
DROP TRIGGER IF EXISTS trg_set_last_update_date ON Task;
DROP FUNCTION IF EXISTS set_last_update_date();
