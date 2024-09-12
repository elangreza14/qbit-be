BEGIN
;

-- delete orders after cart is purchased
CREATE TABLE IF NOT EXISTS "orders" (
    "id" UUID PRIMARY KEY,
    "user_id" UUID REFERENCES users(id),
    "status" VARCHAR(50),
    "total" INT,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMPTZ NULL
);

CREATE TRIGGER "log_cart_update" BEFORE
UPDATE
    ON "orders" FOR EACH ROW EXECUTE PROCEDURE log_update_master();

COMMIT;