-- name: UpsertArticle :exec
INSERT INTO articles (
  id,
  slug,
  title,
  thumbnail_id,
  content,
  tags,
  user_id,
  created_at,
  updated_at,
  deleted_at
)
VALUES (
  sqlc.arg('id'),
  sqlc.arg('slug'),
  sqlc.arg('title'),
  sqlc.narg('thumbnail_id'),
  sqlc.arg('content'),
  sqlc.arg('tags'),
  sqlc.arg('user_id'),
  sqlc.arg('created_at'),
  sqlc.arg('updated_at'),
  sqlc.narg('deleted_at')
)
ON CONFLICT (id) DO UPDATE SET
  slug = EXCLUDED.slug,
  title = EXCLUDED.title,
  thumbnail_id = EXCLUDED.thumbnail_id,
  content = EXCLUDED.content,
  tags = EXCLUDED.tags,
  user_id = EXCLUDED.user_id,
  created_at = EXCLUDED.created_at,
  updated_at = EXCLUDED.updated_at,
  deleted_at = EXCLUDED.deleted_at;

-- name: ListArticles :many
SELECT
  *
FROM
  articles
WHERE
  CASE
    WHEN sqlc.narg('ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.narg('ids')::uuid[]) = 0 THEN TRUE
    ELSE id = ANY (sqlc.narg('ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.narg('user_ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.narg('user_ids')::uuid[]) = 0 THEN TRUE
    ELSE user_id = ANY (sqlc.narg('user_ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.narg('tags')::text[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.narg('tags')::text[]) = 0 THEN TRUE
    ELSE tags && sqlc.narg('tags')::text[]
  END
  AND CASE
    WHEN sqlc.arg('deleted')::text = 'exclude' THEN deleted_at IS NULL
    WHEN sqlc.arg('deleted')::text = 'only' THEN deleted_at IS NOT NULL
    WHEN sqlc.arg('deleted')::text = 'all' THEN TRUE
    ELSE deleted_at IS NULL
  END
ORDER BY
  created_at DESC
OFFSET sqlc.arg('offset')::integer
LIMIT NULLIF(sqlc.arg('limit')::integer, 0);

-- name: CountArticles :one
SELECT
  COUNT(*) AS count
FROM
  articles
WHERE
  CASE
    WHEN sqlc.narg('ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.narg('ids')::uuid[]) = 0 THEN TRUE
    ELSE id = ANY (sqlc.narg('ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.narg('user_ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.narg('user_ids')::uuid[]) = 0 THEN TRUE
    ELSE user_id = ANY (sqlc.narg('user_ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.narg('tags')::text[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.narg('tags')::text[]) = 0 THEN TRUE
    ELSE tags && sqlc.narg('tags')::text[]
  END
  AND CASE
    WHEN sqlc.arg('deleted')::text = 'exclude' THEN deleted_at IS NULL
    WHEN sqlc.arg('deleted')::text = 'only' THEN deleted_at IS NOT NULL
    WHEN sqlc.arg('deleted')::text = 'all' THEN TRUE
    ELSE deleted_at IS NULL
  END;

-- name: GetArticle :one
SELECT
  *
FROM
  articles
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