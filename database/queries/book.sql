-- name: UpsertBook :exec
INSERT INTO books (
  id,
  title,
  description,
  author,
  price,
  pages_count,
  year,
  publisher,
  weight,
  stock_quantity,
  purchase_count,
  rating,
  created_at,
  updated_at,
  deleted_at
)
VALUES (
  sqlc.arg('id'),
  sqlc.arg('title'),
  sqlc.arg('description'),
  sqlc.arg('author'),
  sqlc.arg('price'),
  sqlc.arg('pages_count'),
  sqlc.arg('year'),
  sqlc.arg('publisher'),
  sqlc.arg('weight'),
  sqlc.arg('stock_quantity'),
  sqlc.arg('purchase_count'),
  sqlc.arg('rating'),
  sqlc.arg('created_at'),
  sqlc.arg('updated_at'),
  NULLIF(sqlc.arg('deleted_at')::timestamptz, '0001-01-01T00:00:00Z'::timestamptz)
)
ON CONFLICT (id) DO UPDATE SET
  title = EXCLUDED.title,
  description = EXCLUDED.description,
  author = EXCLUDED.author,
  price = EXCLUDED.price,
  pages_count = EXCLUDED.pages_count,
  year = EXCLUDED.year,
  publisher = EXCLUDED.publisher,
  weight = EXCLUDED.weight,
  stock_quantity = EXCLUDED.stock_quantity,
  purchase_count = EXCLUDED.purchase_count,
  rating = EXCLUDED.rating,
  created_at = EXCLUDED.created_at,
  updated_at = EXCLUDED.updated_at,
  deleted_at = COALESCE(EXCLUDED.deleted_at, books.deleted_at);

-- name: ListBooks :many
SELECT
  books.*
FROM
  books
LEFT JOIN (
  SELECT
    categories.id AS category_id,
    pdb.score(books.id) AS category_score
  FROM books
  INNER JOIN books_categories
    ON books.id = books_categories.book_id
  INNER JOIN categories
    ON books_categories.category_id = categories.id
  WHERE
    CASE
      WHEN sqlc.arg('search')::text = '' THEN TRUE
      ELSE (
        categories.name ||| sqlc.arg('search')::text
        AND categories.deleted_at IS NULL
      )
    END
) AS category_scores
  ON books.id = (SELECT bc.book_id FROM books_categories bc WHERE bc.category_id = category_scores.category_id LIMIT 1)
LEFT JOIN (
  SELECT
    book_id
  FROM
    books_categories
  WHERE
    CASE
      WHEN sqlc.arg('category_ids')::uuid[] IS NULL THEN TRUE
      WHEN cardinality(sqlc.arg('category_ids')::uuid[]) = 0 THEN TRUE
      ELSE category_id = ANY (sqlc.arg('category_ids')::uuid[])
    END
  GROUP BY
    book_id
  HAVING
    COUNT(DISTINCT category_id) >= CASE
      WHEN sqlc.arg('category_ids')::uuid[] IS NULL THEN 0
      ELSE cardinality(sqlc.arg('category_ids')::uuid[])
    END
) AS category_filter
  ON books.id = category_filter.book_id
WHERE
  CASE
    WHEN sqlc.arg('ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.arg('ids')::uuid[]) = 0 THEN TRUE
    ELSE books.id = ANY (sqlc.arg('ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.arg('search')::text = '' THEN TRUE
    ELSE books.title ||| sqlc.arg('search')::text
      OR books.author ||| sqlc.arg('search')::text
      OR books.description ||| sqlc.arg('search')::text
  END
  AND CASE
    WHEN sqlc.arg('category_ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.arg('category_ids')::uuid[]) = 0 THEN TRUE
    ELSE category_filter.book_id IS NOT NULL
  END
  AND CASE
    WHEN sqlc.arg('min_price')::decimal IS NULL THEN TRUE
    WHEN sqlc.arg('min_price')::decimal = 0 THEN TRUE
    ELSE books.price >= sqlc.arg('min_price')::decimal
  END
  AND CASE
    WHEN sqlc.arg('max_price')::decimal IS NULL THEN TRUE
    WHEN sqlc.arg('max_price')::decimal = 0 THEN TRUE
    ELSE books.price <= sqlc.arg('max_price')::decimal
  END
  AND CASE
    WHEN sqlc.arg('rating')::real IS NULL THEN TRUE
    WHEN sqlc.arg('rating')::real = 0 THEN TRUE
    ELSE books.rating >= sqlc.arg('rating')::real
  END
  AND CASE
    WHEN sqlc.arg('publisher')::text = '' THEN TRUE
    ELSE books.publisher ||| sqlc.arg('publisher')::text
  END
  AND CASE
    WHEN sqlc.arg('year')::integer IS NULL THEN TRUE
    WHEN sqlc.arg('year')::integer = 0 THEN TRUE
    ELSE books.year = sqlc.arg('year')::integer
  END
  AND CASE
    WHEN sqlc.arg('deleted')::text = 'exclude' THEN books.deleted_at IS NULL
    WHEN sqlc.arg('deleted')::text = 'only' THEN books.deleted_at IS NOT NULL
    WHEN sqlc.arg('deleted')::text = 'all' THEN TRUE
    ELSE books.deleted_at IS NULL
  END
ORDER BY
  CASE WHEN
    sqlc.arg('search')::text <> '' THEN pdb.score(books.id) + COALESCE(category_scores.category_score, 0)
  END DESC,
  CASE WHEN
    sqlc.arg('sort_rating')::text = 'asc' THEN books.rating
  END ASC,
  CASE WHEN
    sqlc.arg('sort_rating')::text = 'desc' THEN books.rating
  END DESC,
  CASE WHEN
    sqlc.arg('sort_price')::text = 'asc' THEN books.price
  END ASC,
  CASE WHEN
    sqlc.arg('sort_price')::text = 'desc' THEN books.price
  END DESC,
  CASE WHEN
    sqlc.arg('sort_recent')::text = 'asc' THEN books.created_at
  END ASC,
  CASE WHEN
    sqlc.arg('sort_recent')::text = 'desc' THEN books.created_at
  END DESC,
  books.id DESC
OFFSET sqlc.arg('offset')::integer
LIMIT NULLIF(sqlc.arg('limit')::integer, 0);

-- name: CountBooks :one
SELECT
  COUNT(*) AS count
FROM
  books
LEFT JOIN (
  SELECT
    book_id
  FROM
    books_categories
  WHERE
    CASE
      WHEN sqlc.arg('category_ids')::uuid[] IS NULL THEN TRUE
      WHEN cardinality(sqlc.arg('category_ids')::uuid[]) = 0 THEN TRUE
      ELSE category_id = ANY (sqlc.arg('category_ids')::uuid[])
    END
  GROUP BY
    book_id
  HAVING
    COUNT(DISTINCT category_id) >= CASE
      WHEN sqlc.arg('category_ids')::uuid[] IS NULL THEN 0
      WHEN cardinality(sqlc.arg('category_ids')::uuid[]) = 0 THEN 0
      ELSE cardinality(sqlc.arg('category_ids')::uuid[])
    END
) AS category_filter
  ON books.id = category_filter.book_id
WHERE
  CASE
    WHEN sqlc.arg('ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.arg('ids')::uuid[]) = 0 THEN TRUE
    ELSE books.id = ANY (sqlc.arg('ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.arg('category_ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.arg('category_ids')::uuid[]) = 0 THEN TRUE
    ELSE category_filter.book_id IS NOT NULL
  END
  AND CASE
    WHEN sqlc.arg('min_price')::decimal IS NULL THEN TRUE
    WHEN sqlc.arg('min_price')::decimal = 0 THEN TRUE
    ELSE books.price >= sqlc.arg('min_price')::decimal
  END
  AND CASE
    WHEN sqlc.arg('max_price')::decimal IS NULL THEN TRUE
    WHEN sqlc.arg('max_price')::decimal = 0 THEN TRUE
    ELSE books.price <= sqlc.arg('max_price')::decimal
  END
  AND CASE
    WHEN sqlc.arg('rating')::real IS NULL THEN TRUE
    WHEN sqlc.arg('rating')::real = 0 THEN TRUE
    ELSE books.rating >= sqlc.arg('rating')::real
  END
  AND CASE
    WHEN sqlc.arg('deleted')::text = 'exclude' THEN books.deleted_at IS NULL
    WHEN sqlc.arg('deleted')::text = 'only' THEN books.deleted_at IS NOT NULL
    WHEN sqlc.arg('deleted')::text = 'all' THEN TRUE
    ELSE books.deleted_at IS NULL
  END;

-- name: ListBooksMedium :many
SELECT
  *
FROM
  books_medium
WHERE
  CASE
    WHEN sqlc.arg('book_ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.arg('book_ids')::uuid[]) = 0 THEN TRUE
    ELSE book_id = ANY (sqlc.arg('book_ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.arg('media_ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.arg('media_ids')::uuid[]) = 0 THEN TRUE
    ELSE media_id = ANY (sqlc.arg('media_ids')::uuid[])
  END
ORDER BY
  book_id ASC,
  media_id ASC;

-- name: CreateTempTableBooksMedium :exec
CREATE TEMPORARY TABLE temp_books_medium (
  book_id UUID NOT NULL,
  media_id UUID NOT NULL,
  "order" INTEGER NOT NULL,
  is_cover BOOLEAN NOT NULL,
  PRIMARY KEY (book_id, media_id)
) ON COMMIT DROP;

-- name: InsertTempTableBooksMedium :copyfrom
INSERT INTO temp_books_medium (
  book_id,
  media_id,
  "order",
  is_cover
) VALUES (
  $1,
  $2,
  $3,
  $4
);

-- name: MergeBooksMediumFromTemp :exec
MERGE INTO books_medium AS target
USING temp_books_medium AS source
  ON target.book_id = source.book_id
  AND target.media_id = source.media_id
WHEN MATCHED THEN
  UPDATE SET
    "order" = source."order",
    is_cover = source.is_cover
WHEN NOT MATCHED THEN
  INSERT (
    book_id,
    media_id,
    "order",
    is_cover
  )
  VALUES (
    source.book_id,
    source.media_id,
    source."order",
    source.is_cover
  )
WHEN NOT MATCHED BY SOURCE
  AND target.book_id = (SELECT DISTINCT book_id FROM temp_books_medium) THEN
  DELETE;

-- name: GetBook :one
SELECT
  *
FROM
  books
WHERE
  books.id = sqlc.arg('id')
  AND CASE
    WHEN sqlc.arg('deleted')::text = 'exclude' THEN deleted_at IS NULL
    WHEN sqlc.arg('deleted')::text = 'only' THEN deleted_at IS NOT NULL
    WHEN sqlc.arg('deleted')::text = 'all' THEN TRUE
    ELSE deleted_at IS NULL
  END;

-- name: ListBooksCategories :many
SELECT
  *
FROM
  books_categories
WHERE
  CASE
    WHEN sqlc.arg('book_ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.arg('book_ids')::uuid[]) = 0 THEN TRUE
    ELSE book_id = ANY (sqlc.arg('book_ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.arg('category_ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.arg('category_ids')::uuid[]) = 0 THEN TRUE
    ELSE category_id = ANY (sqlc.arg('category_ids')::uuid[])
  END
ORDER BY
  book_id ASC,
  category_id ASC;

-- name: CreateTempTableBooksCategories :exec
CREATE TEMPORARY TABLE temp_books_categories (
  book_id UUID NOT NULL,
  category_id UUID NOT NULL,
  PRIMARY KEY (book_id, category_id)
) ON COMMIT DROP;

-- name: InsertTempTableBooksCategories :copyfrom
INSERT INTO temp_books_categories (
  book_id,
  category_id
) VALUES (
  sqlc.arg('book_id'),
  sqlc.arg('category_id')
);

-- name: MergeBooksCategoriesFromTemp :exec
MERGE INTO books_categories AS target
USING temp_books_categories AS source
  ON target.book_id = source.book_id
  AND target.category_id = source.category_id
WHEN NOT MATCHED THEN
  INSERT (
    book_id,
    category_id
  )
  VALUES (
    source.book_id,
    source.category_id
  )
WHEN NOT MATCHED BY SOURCE
  AND target.book_id = (SELECT DISTINCT book_id FROM temp_books_categories) THEN
  DELETE;