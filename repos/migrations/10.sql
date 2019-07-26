CREATE OR REPLACE FUNCTION block_insert_notify() RETURNS trigger AS $$
DECLARE
  lvl bigint;
BEGIN
  PERFORM pg_notify('block_insert', json_build_object('lvl', NEW.level, 'type', TG_OP)::text);
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS blocks_notify_insert ON blocks;
CREATE TRIGGER blocks_notify_insert AFTER INSERT ON blocks FOR EACH ROW EXECUTE PROCEDURE block_insert_notify();
