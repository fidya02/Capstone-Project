BEGIN;

ALTER TABLE "public"."orders" RENAME COLUMN qty TO quantity;

COMMIT;