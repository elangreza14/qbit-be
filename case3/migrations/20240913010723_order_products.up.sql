BEGIN
;

CREATE TABLE IF NOT EXISTS "order_products" (
    "id" SERIAL PRIMARY KEY,
    "order_id" UUID REFERENCES orders(id),
    "product_id" INT REFERENCES products(id),
    "quantity" INT,
    "price" INT,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMPTZ NULL
);

CREATE TRIGGER "log_order_product_update" BEFORE
UPDATE
    ON "order_products" FOR EACH ROW EXECUTE PROCEDURE log_update_master();

COMMIT;