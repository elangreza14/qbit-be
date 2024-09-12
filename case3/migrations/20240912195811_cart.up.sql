BEGIN
;

CREATE TABLE IF NOT EXISTS "carts" (
    "id" SERIAL PRIMARY KEY,
    "user_id" UUID,
    "product_id" INT REFERENCES products(id),
    "quantity" INT,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMPTZ NULL
);

CREATE TRIGGER "log_cart_update" BEFORE
UPDATE
    ON "carts" FOR EACH ROW EXECUTE PROCEDURE log_update_master();

COMMIT;