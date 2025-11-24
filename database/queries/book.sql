-- name: UpsertBook :exec
INSERT INTO books (
  id,
  title,
  description,
  author,
  price,
  pages_count,
  year_published,
  publisher,
  weight,
  stock_quantity,
  purchase_count,
  rating,
  category_id,
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
  sqlc.arg('year_published'),
  sqlc.arg('publisher'),
  sqlc.arg('weight'),
  sqlc.arg('stock_quantity'),
  sqlc.arg('purchase_count'),
  sqlc.arg('rating'),
  sqlc.arg('category_id'),
  sqlc.arg('created_at'),
  sqlc.arg('updated_at'),
  sqlc.narg('deleted_at')
)
ON CONFLICT (id) DO UPDATE SET
  title = EXCLUDED.title,
  description = EXCLUDED.description,
  author = EXCLUDED.author,
  price = EXCLUDED.price,
  pages_count = EXCLUDED.pages_count,
  year_published = EXCLUDED.year_published,
  publisher = EXCLUDED.publisher,
  weight = EXCLUDED.weight,
  stock_quantity = EXCLUDED.stock_quantity,
  purchase_count = EXCLUDED.purchase_count,
  rating = EXCLUDED.rating,
  category_id = EXCLUDED.category_id,
  created_at = EXCLUDED.created_at,
  updated_at = EXCLUDED.updated_at,
  deleted_at = EXCLUDED.deleted_at;

-- name: ListBooks :many
SELECT
  books.*
FROM
  books
WHERE
  CASE
    WHEN sqlc.narg('ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.narg('ids')::uuid[]) = 0 THEN TRUE
    ELSE books.id = ANY (sqlc.narg('ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.narg('category_ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.narg('category_ids')::uuid[]) = 0 THEN TRUE
    ELSE books.category_id = ANY (sqlc.narg('category_ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.narg('min_price')::decimal IS NULL THEN TRUE
    ELSE books.price >= sqlc.narg('min_price')::decimal
  END
  AND CASE
    WHEN sqlc.narg('max_price')::decimal IS NULL THEN TRUE
    ELSE books.price <= sqlc.narg('max_price')::decimal
  END
  AND CASE
    WHEN sqlc.narg('rating')::real IS NULL THEN TRUE
    ELSE books.rating >= sqlc.narg('rating')::real
  END
  AND CASE
    WHEN sqlc.arg('deleted')::text = 'exclude' THEN books.deleted_at IS NULL
    WHEN sqlc.arg('deleted')::text = 'only' THEN books.deleted_at IS NOT NULL
    WHEN sqlc.arg('deleted')::text = 'all' THEN TRUE
    ELSE books.deleted_at IS NULL
  END
ORDER BY
  CASE WHEN
    sqlc.narg('sort_rating')::text = 'asc' THEN books.rating
  END ASC,
  CASE WHEN
    sqlc.narg('sort_rating')::text = 'desc' THEN books.rating
  END DESC,
  CASE WHEN
    sqlc.narg('sort_price')::text = 'asc' THEN books.price
  END ASC,
  CASE WHEN
    sqlc.narg('sort_price')::text = 'desc' THEN books.price
  END DESC,
  books.id DESC
OFFSET sqlc.arg('offset')::integer
LIMIT NULLIF(sqlc.arg('limit')::integer, 0);

-- name: CountBooks :one
SELECT
  COUNT(*) AS count
FROM
  books
WHERE
  CASE
    WHEN sqlc.narg('ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.narg('ids')::uuid[]) = 0 THEN TRUE
    ELSE books.id = ANY (sqlc.narg('ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.narg('category_ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.narg('category_ids')::uuid[]) = 0 THEN TRUE
    ELSE books.category_id = ANY (sqlc.narg('category_ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.narg('min_price')::decimal IS NULL THEN TRUE
    ELSE books.price >= sqlc.narg('min_price')::decimal
  END
  AND CASE
    WHEN sqlc.narg('max_price')::decimal IS NULL THEN TRUE
    ELSE books.price <= sqlc.narg('max_price')::decimal
  END
  AND CASE
    WHEN sqlc.narg('rating')::real IS NULL THEN TRUE
    ELSE books.rating >= sqlc.narg('rating')::real
  END
  AND CASE
    WHEN sqlc.arg('deleted')::text = 'exclude' THEN books.deleted_at IS NULL
    WHEN sqlc.arg('deleted')::text = 'only' THEN books.deleted_at IS NOT NULL
    WHEN sqlc.arg('deleted')::text = 'all' THEN TRUE
    ELSE books.deleted_at IS NULL
  END;

-- name: ListMedia :many
SELECT
  *
FROM
  media
WHERE
  CASE
    WHEN sqlc.narg('ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.narg('ids')::uuid[]) = 0 THEN TRUE
    ELSE id = ANY (sqlc.narg('ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.narg('book_ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.narg('book_ids')::uuid[]) = 0 THEN TRUE
    ELSE book_id = ANY (sqlc.narg('book_ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.arg('deleted')::text = 'exclude' THEN deleted_at IS NULL
    WHEN sqlc.arg('deleted')::text = 'only' THEN deleted_at IS NOT NULL
    WHEN sqlc.arg('deleted')::text = 'all' THEN TRUE
    ELSE deleted_at IS NULL
  END
ORDER BY
  book_id ASC,
  "order" ASC,
  id ASC;

-- name: GetMedia :one
SELECT
  *
FROM
  media
WHERE
  id = sqlc.arg('id')
  AND CASE
    WHEN sqlc.arg('deleted')::text = 'exclude' THEN deleted_at IS NULL
    WHEN sqlc.arg('deleted')::text = 'only' THEN deleted_at IS NOT NULL
    WHEN sqlc.arg('deleted')::text = 'all' THEN TRUE
    ELSE deleted_at IS NULL
  END;

-- name: CreateTempTableMedia :exec
CREATE TEMPORARY TABLE temp_media (
  id UUID PRIMARY KEY,
  url TEXT NOT NULL,
  alt_text TEXT,
  "order" INTEGER NOT NULL,
  book_id UUID,
  created_at TIMESTAMPTZ NOT NULL,
  deleted_at TIMESTAMPTZ
) ON COMMIT DROP;

-- name: InsertTempTableMedia :copyfrom
INSERT INTO temp_media (
  id,
  url,
  alt_text,
  "order",
  book_id,
  created_at,
  deleted_at
) VALUES (
  @id,
  @url,
  @alt_text,
  @order,
  @book_id,
  @created_at,
  @deleted_at
);

-- name: MergeMediaFromTemp :exec
MERGE INTO media AS target
USING temp_media AS source
  ON target.id = source.id
WHEN MATCHED THEN
  UPDATE SET
    url = source.url,
    alt_text = source.alt_text,
    "order" = source."order",
    book_id = source.book_id,
    created_at = source.created_at,
    deleted_at = source.deleted_at
WHEN NOT MATCHED THEN
  INSERT (
    id,
    url,
    alt_text,
    "order",
    book_id,
    created_at,
    deleted_at
  )
  VALUES (
    source.id,
    source.url,
    source.alt_text,
    source."order",
    source.book_id,
    source.created_at,
    source.deleted_at
  )
WHEN NOT MATCHED BY SOURCE
  AND target.book_id = (SELECT DISTINCT book_id FROM source) THEN
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