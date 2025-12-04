-- name: UpsertMedia :exec
INSERT INTO medium (
  id,
  url,
  alt_text,
  created_at,
  deleted_at
)
VALUES (
  sqlc.arg('id'),
  sqlc.arg('url'),
  NULLIF(sqlc.arg('alt_text')::text, ''),
  sqlc.arg('created_at'),
  NULLIF(sqlc.arg('deleted_at')::timestamptz, '0001-01-01T00:00:00Z'::timestamptz)
)
ON CONFLICT (id) DO UPDATE SET
  url = EXCLUDED.url,
  alt_text = EXCLUDED.alt_text,
  created_at = EXCLUDED.created_at,
  deleted_at = COALESCE(EXCLUDED.deleted_at, media.deleted_at);

-- name: ListMedium :many
SELECT
  *
FROM
  medium
WHERE
  CASE
    WHEN sqlc.arg('ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.arg('ids')::uuid[]) = 0 THEN TRUE
    ELSE id = ANY (sqlc.arg('ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.arg('deleted')::text = 'exclude' THEN deleted_at IS NULL
    WHEN sqlc.arg('deleted')::text = 'only' THEN deleted_at IS NOT NULL
    WHEN sqlc.arg('deleted')::text = 'all' THEN TRUE
    ELSE deleted_at IS NULL
  END
ORDER BY
  id ASC
OFFSET sqlc.arg('offset')::integer
LIMIT NULLIF(sqlc.arg('limit')::integer, 0);

-- name: CountMedium :one
SELECT
  COUNT(*)
FROM
  medium
WHERE
  CASE
    WHEN sqlc.arg('ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.arg('ids')::uuid[]) = 0 THEN TRUE
    ELSE id = ANY (sqlc.arg('ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.arg('deleted')::text = 'exclude' THEN deleted_at IS NULL
    WHEN sqlc.arg('deleted')::text = 'only' THEN deleted_at IS NOT NULL
    WHEN sqlc.arg('deleted')::text = 'all' THEN TRUE
    ELSE deleted_at IS NULL
  END;

-- name: GetMedia :one
SELECT
  *
FROM
  medium
WHERE
  id = sqlc.arg('id')
  AND CASE
    WHEN sqlc.arg('deleted')::text = 'exclude' THEN deleted_at IS NULL
    WHEN sqlc.arg('deleted')::text = 'only' THEN deleted_at IS NOT NULL
    WHEN sqlc.arg('deleted')::text = 'all' THEN TRUE
    ELSE deleted_at IS NULL
  END;
