ALTER TABLE "users" ADD COLUMN "gender" int NOT NULL;
ALTER TABLE "users" ADD COLUMN "image" varchar;
ALTER TABLE "users" ADD COLUMN "phone" varchar UNIQUE NOT NULL;
