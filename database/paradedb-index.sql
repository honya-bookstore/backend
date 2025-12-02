CREATE INDEX IF NOT EXISTS categories_search_idx ON categories
USING bm25 (id, name, description)
WITH (key_field = 'id');

CREATE INDEX IF NOT EXISTS books_search_idx ON books
USING bm25 (id, title, author, description)
WITH (key_field = 'id');
