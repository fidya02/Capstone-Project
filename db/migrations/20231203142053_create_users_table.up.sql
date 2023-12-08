BEGIN;
CREATE TABLE IF NOT EXISTS "public"."users" (
    "id" SERIAL NOT NULL PRIMARY KEY,
    "name" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
    "number" varchar(255) NOT NULL,
    "email" varchar(255) NOT NULL,
    "role" varchar(255) NOT NULL,
    "password" varchar(255) NOT NULL,
    "wallet" integer NOT NULL DEFAULT 0,
    "created_at" timestamptz (6),
    "updated_at" timestamptz (6),
    "deleted_at" timestamptz (6)
);
COMMIT;