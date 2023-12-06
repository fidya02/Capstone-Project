BEGIN;

CREATE TABLE "public"."transactions" (
    "id" SERIAL PRIMARY KEY,
    "name" TEXT,
    "category" TEXT,
    "price" INT,
    "stock" INTEGER,
    "image" TEXT,
    "Status" TEXT DEFAULT 'available',
    "Date" DATE,
    "Location" TEXT,
    "created_at" TIMESTAMP,
    "updated_at" TIMESTAMP,
    "deleted_at" TIMESTAMP,
    "updated_by" VARCHAR(255),
    "created_by" VARCHAR(255)
);

COMMIT;