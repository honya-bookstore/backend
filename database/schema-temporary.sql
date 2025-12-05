-- temp_books_medium
CREATE TEMPORARY TABLE temp_books_medium (
  book_id UUID NOT NULL,
  media_id UUID NOT NULL,
  "order" INTEGER NOT NULL,
  is_cover BOOLEAN NOT NULL,
  PRIMARY KEY (book_id, media_id)
);

-- temp_cart_items
CREATE TEMPORARY TABLE temp_cart_items (
  id UUID PRIMARY KEY,
  quantity INTEGER NOT NULL,
  cart_id UUID NOT NULL,
  book_id UUID NOT NULL
);

-- temp_order_items
CREATE TEMPORARY TABLE temp_order_items (
  id UUID PRIMARY KEY,
  quantity INTEGER NOT NULL,
  order_id UUID NOT NULL,
  price DECIMAL(12, 0) NOT NULL,
  book_id UUID NOT NULL
);

-- temp_books_categories
CREATE TEMPORARY TABLE temp_books_categories (
  book_id UUID NOT NULL,
  category_id UUID NOT NULL,
  PRIMARY KEY (book_id, category_id)
);
