-- name: UpsertCategory :exec
INSERT INTO categories (
  id,
  slug,
  name,
  description,
  created_at,
  updated_at,
  deleted_at
)
VALUES (
  sqlc.arg('id'),
  sqlc.arg('slug'),
  sqlc.arg('name'),
  sqlc.arg('description'),
  sqlc.arg('created_at'),
  sqlc.arg('updated_at'),
  sqlc.narg('deleted_at')
)
ON CONFLICT (id) DO UPDATE SET
  slug = EXCLUDED.slug,
  name = EXCLUDED.name,
  description = EXCLUDED.description,
  created_at = EXCLUDED.created_at,
  updated_at = EXCLUDED.updated_at,
  deleted_at = EXCLUDED.deleted_at;

-- name: ListCategories :many
SELECT
  *
FROM
  categories
WHERE
  CASE
    WHEN sqlc.narg('ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.narg('ids')::uuid[]) = 0 THEN TRUE
    ELSE id = ANY (sqlc.narg('ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.arg('deleted')::text = 'exclude' THEN deleted_at IS NULL
    WHEN sqlc.arg('deleted')::text = 'only' THEN deleted_at IS NOT NULL
    WHEN sqlc.arg('deleted')::text = 'all' THEN TRUE
    ELSE deleted_at IS NULL
  END
ORDER BY
  id DESC
OFFSET sqlc.arg('offset')::integer
LIMIT NULLIF(sqlc.arg('limit')::integer, 0);

-- name: CountCategories :one
SELECT
  COUNT(*) AS count
FROM
  categories
WHERE
  CASE
    WHEN sqlc.arg('deleted')::text = 'exclude' THEN deleted_at IS NULL
    WHEN sqlc.arg('deleted')::text = 'only' THEN deleted_at IS NOT NULL
    WHEN sqlc.arg('deleted')::text = 'all' THEN TRUE
    ELSE deleted_at IS NULL
  END;

-- name: GetCategory :one
SELECT
  *
FROM
  categories
WHERE
  CASE
    WHEN sqlc.narg('id')::uuid IS NULL THEN TRUE
    ELSE id = sqlc.narg('id')::uuid
  END
  AND CASE
    WHEN sqlc.narg('slug')::text IS NULL THEN TRUE
    ELSE slug = sqlc.narg('slug')::text
  END
  AND CASE
    WHEN sqlc.arg('deleted')::text = 'exclude' THEN deleted_at IS NULL
    WHEN sqlc.arg('deleted')::text = 'only' THEN deleted_at IS NOT NULL
    WHEN sqlc.arg('deleted')::text = 'all' THEN TRUE
    ELSE deleted_at IS NULL
  END;
