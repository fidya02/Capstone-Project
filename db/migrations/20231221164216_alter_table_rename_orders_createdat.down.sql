BEGIN;

ALTER TABLE "public"."orders" DROP COLUMN IF EXISTS created_at;

COMMIT;