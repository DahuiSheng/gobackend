-- Modify "accesses" table
ALTER TABLE "accesses" ADD COLUMN "create_time" timestamptz NOT NULL, ADD COLUMN "update_time" timestamptz NOT NULL, ADD COLUMN "name" character varying NOT NULL DEFAULT 'unknown', ADD COLUMN "check_in" timestamptz NOT NULL, ADD COLUMN "check_out" timestamptz NULL;
