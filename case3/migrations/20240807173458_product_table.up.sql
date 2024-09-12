BEGIN
;

CREATE TABLE IF NOT EXISTS "products" (
    "id" SERIAL PRIMARY KEY,
    "name" VARCHAR NOT NULL,
    "description" VARCHAR,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMPTZ NULL
);

CREATE TRIGGER "log_product_update" BEFORE
UPDATE
    ON "products" FOR EACH ROW EXECUTE PROCEDURE log_update_master();

COMMIT;