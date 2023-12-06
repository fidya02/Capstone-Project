BEGIN;

CREATE TABLE "public"."transactions" (
    "id" SERIAL PRIMARY KEY,
    "ticket_id" INT,
    "user_id" INT,
    "qty" INT,
    "total" INTEGER,
    "Status" VARCHAR(255),
    "updated_at" TIMESTAMP,
    "updated_at" TIMESTAMP,
    "deleted_at" TIMESTAMP,
    "updated_by" VARCHAR(255),
    "deleted_by" VARCHAR(255)
    FOREIGN KEY (ticket_id) REFERENCES tickets(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE
);

COMMIT;