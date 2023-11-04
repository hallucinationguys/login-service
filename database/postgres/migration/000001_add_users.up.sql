CREATE TABLE "users" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "password" varchar(255) NOT NULL,
  "last_name" varchar NOT NULL,
  "first_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "phone" varchar NOT NULL,
  "role" int NOT NULL,
  "status" int NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z'
);
