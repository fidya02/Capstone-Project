BEGIN;

ALTER TABLE "public"."orders" RENAME COLUMN created_at TO order_at;

COMMIT;