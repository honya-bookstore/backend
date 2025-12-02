-- temp_media
CREATE TABLE temp_media (
  id UUID PRIMARY KEY,
  url TEXT NOT NULL,
  alt_text TEXT,
  "order" INTEGER NOT NULL,
  book_id UUID,
  created_at TIMESTAMPTZ NOT NULL,
  deleted_at TIMESTAMPTZ
);

-- temp_cart_items
CREATE TABLE temp_cart_items (
  id UUID PRIMARY KEY,
  quantity INTEGER NOT NULL,
  cart_id UUID NOT NULL,
  book_id UUID NOT NULL
);

-- temp_order_items
CREATE TABLE temp_order_items (
  id UUID PRIMARY KEY,
  quantity INTEGER NOT NULL,
  order_id UUID NOT NULL,
  price DECIMAL(12, 0) NOT NULL,
  book_id UUID NOT NULL
);
