-- Create "medium" table
CREATE TABLE "public"."medium" (
  "id" uuid NOT NULL,
  "url" text NOT NULL,
  "alt_text" text NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "deleted_at" timestamptz NULL,
  PRIMARY KEY ("id")
);
-- Create "articles" table
CREATE TABLE "public"."articles" (
  "id" uuid NOT NULL,
  "slug" text NOT NULL,
  "title" text NOT NULL,
  "thumbnail_id" uuid NULL,
  "content" text NOT NULL,
  "tags" text[] NULL,
  "user_id" uuid NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now(),
  "deleted_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "articles_slug_key" UNIQUE ("slug"),
  CONSTRAINT "articles_thumbnail_id_fkey" FOREIGN KEY ("thumbnail_id") REFERENCES "public"."medium" ("id") ON UPDATE CASCADE ON DELETE NO ACTION
);
-- Create "books" table
CREATE TABLE "public"."books" (
  "id" uuid NOT NULL,
  "title" text NOT NULL,
  "description" text NULL,
  "author" text NOT NULL,
  "price" numeric(12) NOT NULL,
  "pages_count" integer NOT NULL,
  "year" integer NOT NULL,
  "publisher" text NOT NULL,
  "weight" numeric NULL,
  "stock_quantity" integer NOT NULL DEFAULT 0,
  "purchase_count" integer NOT NULL DEFAULT 0,
  "rating" real NOT NULL DEFAULT 0,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now(),
  "deleted_at" timestamptz NULL,
  PRIMARY KEY ("id")
);
-- Create "categories" table
CREATE TABLE "public"."categories" (
  "id" uuid NOT NULL,
  "slug" text NOT NULL,
  "name" text NOT NULL,
  "description" text NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now(),
  "deleted_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "categories_slug_key" UNIQUE ("slug")
);
-- Create "books_categories" table
CREATE TABLE "public"."books_categories" (
  "book_id" uuid NOT NULL,
  "category_id" uuid NOT NULL,
  PRIMARY KEY ("book_id", "category_id"),
  CONSTRAINT "books_categories_book_id_fkey" FOREIGN KEY ("book_id") REFERENCES "public"."books" ("id") ON UPDATE CASCADE ON DELETE NO ACTION,
  CONSTRAINT "books_categories_category_id_fkey" FOREIGN KEY ("category_id") REFERENCES "public"."categories" ("id") ON UPDATE CASCADE ON DELETE NO ACTION
);
-- Create "books_medium" table
CREATE TABLE "public"."books_medium" (
  "book_id" uuid NOT NULL,
  "media_id" uuid NOT NULL,
  "order" integer NOT NULL DEFAULT 0,
  "is_cover" boolean NOT NULL,
  PRIMARY KEY ("book_id", "media_id"),
  CONSTRAINT "books_medium_book_id_fkey" FOREIGN KEY ("book_id") REFERENCES "public"."books" ("id") ON UPDATE CASCADE ON DELETE NO ACTION,
  CONSTRAINT "books_medium_media_id_fkey" FOREIGN KEY ("media_id") REFERENCES "public"."medium" ("id") ON UPDATE CASCADE ON DELETE NO ACTION
);
-- Create "carts" table
CREATE TABLE "public"."carts" (
  "id" uuid NOT NULL,
  "user_id" uuid NOT NULL,
  "updated_at" timestamptz NOT NULL DEFAULT now(),
  PRIMARY KEY ("id"),
  CONSTRAINT "carts_user_id_key" UNIQUE ("user_id")
);
-- Create "cart_items" table
CREATE TABLE "public"."cart_items" (
  "id" uuid NOT NULL,
  "quantity" integer NOT NULL DEFAULT 1,
  "cart_id" uuid NOT NULL,
  "book_id" uuid NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "cart_items_book_id_fkey" FOREIGN KEY ("book_id") REFERENCES "public"."books" ("id") ON UPDATE CASCADE ON DELETE NO ACTION,
  CONSTRAINT "cart_items_cart_id_fkey" FOREIGN KEY ("cart_id") REFERENCES "public"."carts" ("id") ON UPDATE CASCADE ON DELETE NO ACTION
);
-- Create "order_providers" table
CREATE TABLE "public"."order_providers" (
  "id" uuid NOT NULL,
  "name" text NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "order_providers_name_key" UNIQUE ("name")
);
-- Create "order_statuses" table
CREATE TABLE "public"."order_statuses" (
  "id" uuid NOT NULL,
  "name" text NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "order_statuses_name_key" UNIQUE ("name")
);
-- Create "orders" table
CREATE TABLE "public"."orders" (
  "id" uuid NOT NULL,
  "email" text NOT NULL,
  "first_name" text NOT NULL,
  "last_name" text NOT NULL,
  "address" text NOT NULL,
  "city" text NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now(),
  "total_amount" numeric(12) NOT NULL,
  "is_paid" boolean NOT NULL DEFAULT false,
  "user_id" uuid NOT NULL,
  "status_id" uuid NOT NULL,
  "provider_id" uuid NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "orders_provider_id_fkey" FOREIGN KEY ("provider_id") REFERENCES "public"."order_providers" ("id") ON UPDATE CASCADE ON DELETE NO ACTION,
  CONSTRAINT "orders_status_id_fkey" FOREIGN KEY ("status_id") REFERENCES "public"."order_statuses" ("id") ON UPDATE CASCADE ON DELETE NO ACTION
);
-- Create "order_items" table
CREATE TABLE "public"."order_items" (
  "id" uuid NOT NULL,
  "quantity" integer NOT NULL,
  "order_id" uuid NOT NULL,
  "price" numeric(12) NOT NULL,
  "book_id" uuid NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "order_items_book_id_fkey" FOREIGN KEY ("book_id") REFERENCES "public"."books" ("id") ON UPDATE CASCADE ON DELETE NO ACTION,
  CONSTRAINT "order_items_order_id_fkey" FOREIGN KEY ("order_id") REFERENCES "public"."orders" ("id") ON UPDATE CASCADE ON DELETE NO ACTION
);
-- Create "reviews" table
CREATE TABLE "public"."reviews" (
  "id" uuid NOT NULL,
  "rating" smallint NOT NULL,
  "vote_count" integer NOT NULL DEFAULT 0,
  "content" text NULL,
  "user_id" uuid NOT NULL,
  "book_id" uuid NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  "updated_at" timestamptz NOT NULL DEFAULT now(),
  "deleted_at" timestamptz NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "reviews_book_id_fkey" FOREIGN KEY ("book_id") REFERENCES "public"."books" ("id") ON UPDATE CASCADE ON DELETE NO ACTION,
  CONSTRAINT "reviews_rating_check" CHECK ((rating >= 1) AND (rating <= 5))
);
-- Create "review_votes" table
CREATE TABLE "public"."review_votes" (
  "id" uuid NOT NULL,
  "user_id" uuid NOT NULL,
  "review_id" uuid NOT NULL,
  "is_up" boolean NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT now(),
  PRIMARY KEY ("id"),
  CONSTRAINT "review_votes_review_id_fkey" FOREIGN KEY ("review_id") REFERENCES "public"."reviews" ("id") ON UPDATE CASCADE ON DELETE NO ACTION
);
