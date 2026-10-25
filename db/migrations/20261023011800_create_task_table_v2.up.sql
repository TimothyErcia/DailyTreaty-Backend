CREATE TABLE Task (
  taskId uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  category_type VARCHAR NOT NULL,
  created_date TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  last_update_date TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  last_update_by VARCHAR NOT NULL,
  price INTEGER NOT NULL DEFAULT 0
);

CREATE OR REPLACE FUNCTION set_last_update_date()
RETURNS trigger AS $$
BEGIN
  NEW.last_update_date = now();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_set_last_update_date
BEFORE UPDATE ON Task
FOR EACH ROW
WHEN (OLD.* IS DISTINCT FROM NEW.*)
EXECUTE FUNCTION set_last_update_date();
