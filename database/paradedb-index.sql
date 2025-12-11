CREATE INDEX IF NOT EXISTS categories_search_idx ON categories
USING bm25 (id, name, description)
WITH (key_field = 'id');

CREATE INDEX IF NOT EXISTS books_search_idx ON books
USING bm25 (id, title, author, description, publisher)
WITH (key_field = 'id');

CREATE INDEX IF NOT EXISTS medium_search_idx ON medium
USING bm25 (id, url, alt_text)
WITH (key_field = 'id');
