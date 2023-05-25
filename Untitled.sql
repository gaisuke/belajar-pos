CREATE SCHEMA "ecommerce";

CREATE TYPE "ecommerce"."products_status" AS ENUM (
  'out_of_stock',
  'in_stock',
  'running_low'
);

CREATE TABLE "ecommerce"."order_items" (
  "order_id" int,
  "product_id" int,
  "quantity" int DEFAULT 1
);

CREATE TABLE "ecommerce"."orders" (
  "id" int PRIMARY KEY,
  "status" varchar,
  "created_at" varchar
);

CREATE TABLE "ecommerce"."products" (
  "id" int PRIMARY KEY,
  "name" varchar,
  "price" int,
  "status" ecommerce.products_status,
  "created_at" datetime DEFAULT (now())
);

COMMENT ON COLUMN "ecommerce"."orders"."created_at" IS 'When order created';

ALTER TABLE "ecommerce"."order_items" ADD FOREIGN KEY ("order_id") REFERENCES "ecommerce"."orders" ("id");

ALTER TABLE "ecommerce"."order_items" ADD FOREIGN KEY ("product_id") REFERENCES "ecommerce"."products" ("id");
