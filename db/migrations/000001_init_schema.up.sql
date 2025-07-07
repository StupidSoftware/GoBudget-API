CREATE TYPE "transaction_type" AS ENUM (
  'income',
  'expense'
);

CREATE TABLE "users" (
  "id" UUID PRIMARY KEY NOT NULL,
  "username" VARCHAR(20) UNIQUE NOT NULL,
  "password" VARCHAR(255) NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
  "balance" BIGINT NOT NULL DEFAULT 0
);

CREATE TABLE "categories" (
  "id" UUID PRIMARY KEY NOT NULL,
  "name" VARCHAR(46) NOT NULL,
  "user_id" UUID,
  FOREIGN KEY ("user_id") REFERENCES "users" ("id")
);

CREATE TABLE "transactions" (
  "id" UUID PRIMARY KEY NOT NULL,
  "user_id" UUID NOT NULL,
  "category_id" UUID NOT NULL,
  "description" VARCHAR(255) NOT NULL,
  "amount" BIGINT NOT NULL,
  "type" transaction_type NOT NULL,
  "date" TIMESTAMP NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
  FOREIGN KEY ("category_id") REFERENCES "categories" ("id"),
  CONSTRAINT "user_transactions" FOREIGN KEY ("user_id") REFERENCES "users" ("id")
);
