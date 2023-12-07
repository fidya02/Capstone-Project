BEGIN;
CREATE TABLE IF NOT EXISTS "public"."tickets" (
    "id" SERIAL NOT NULL PRIMARY KEY,
    "name" varchar(100),
    "description" text,
    "price" int,
    "status" text default 'available',
    "image" varchar(255),
    "location" varchar(255),
    "quantity" int,
    "category" varchar(255),
    "date" date,
    "sold" int,
    "created_at" timestamptz (6),
    "updated_at" timestamptz (6),
    "deleted_at" timestamptz (6)

);
COMMIT;