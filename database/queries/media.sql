-- name: UpsertMedia :exec
INSERT INTO media (
  id,
  url,
  alt_text,
  "order",
  book_id,
  created_at,
  deleted_at
)
VALUES (
  sqlc.arg('id'),
  sqlc.arg('url'),
  sqlc.arg('alt_text'),
  sqlc.arg('order'),
  sqlc.arg('book_id'),
  sqlc.arg('created_at'),
  NULLIF(sqlc.arg('deleted_at')::timestamptz, '0001-01-01T00:00:00Z'::timestamptz)
)
ON CONFLICT (id) DO UPDATE SET
  url = EXCLUDED.url,
  alt_text = EXCLUDED.alt_text,
  "order" = EXCLUDED."order",
  book_id = EXCLUDED.book_id,
  created_at = EXCLUDED.created_at,
  deleted_at = COALESCE(EXCLUDED.deleted_at, media.deleted_at);

-- name: ListMedia :many
SELECT
  *
FROM
  media
WHERE
  CASE
    WHEN sqlc.arg('ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.arg('ids')::uuid[]) = 0 THEN TRUE
    ELSE id = ANY (sqlc.arg('ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.arg('book_ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.arg('book_ids')::uuid[]) = 0 THEN TRUE
    ELSE book_id = ANY (sqlc.arg('book_ids')::uuid[])
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
