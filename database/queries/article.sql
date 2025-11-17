-- name: CreateArticle :one
INSERT INTO articles (
  slug,
  title,
  thumbnail,
  content,
  tags,
  user_id,
  created_at
) VALUES (
  sql.arg('slug'),
  sql.arg('title'),
  sql.arg('thumbnail'),
  sql.arg('content'),
  sql.arg('tags'),
  sql.arg('user_id'),
  COALESCE(sqlc.narg('created_at')::timestamp, NOW())
);

-- name: ListArticles :many
SELECT
  *,
  COUNT(*) OVER() AS current_count,
  COUNT(*) AS total_count
FROM
  articles
WHERE
  CASE
    WHEN sqlc.narg('search')::text IS NULL THEN TRUE
    ELSE title ||| (sqlc.narg('search')::text)::pdb.fuzzy(2)
  END
  AND CASE
    WHEN sqlc.narg('ids')::integer[] IS NULL THEN TRUE
    ELSE id = ANY (sqlc.narg('ids')::integer[])
  END
  AND CASE
    WHEN sqlc.narg('slugs')::text[] IS NULL THEN TRUE
    ELSE slug = ANY (sqlc.narg('slugs')::text[])
  END
  AND CASE
    WHEN sqlc.narg('user_ids')::uuid[] IS NULL THEN TRUE
    ELSE user_id = ANY (sqlc.narg('user_ids')::uuid[])
  END
  AND CASE
    WHEN sql.arg('deleted')::text = 'exclude' THEN deleted_at IS NOT NULL
    WHEN sql.arg('deleted')::text = 'only' THEN deleted_at IS NULL
    WHEN sql.arg('deleted')::text = 'all' THEN TRUE
    ELSE FALSE
  END
ORDER BY
  CASE WHEN
    sqlc.narg('search') IS NOT NULL THEN pdb.score(id)
  END DESC,
  id DESC
OFFSET COALESCE(sqlc.narg('offset')::integer, 0)
LIMIT COALESCE(sqlc.narg('limit')::integer, 20);

-- name: GetArticle :one
SELECT
  *
FROM
  articles
WHERE
  CASE
    WHEN sqlc.narg('id')::integer IS NULL THEN TRUE
    ELSE id = sqlc.narg('id')::integer
  END
  AND CASE
    WHEN sqlc.narg('slug')::text IS NULL THEN TRUE
    ELSE slug = sqlc.narg('slug')::text
  END
  AND CASE
    WHEN sql.arg('deleted')::text = 'exclude' THEN deleted_at IS NOT NULL
    WHEN sql.arg('deleted')::text = 'only' THEN deleted_at IS NULL
    WHEN sql.arg('deleted')::text = 'all' THEN TRUE
    ELSE FALSE
  END;

-- name: UpdateArticle :one
UPDATE
  articles
SET
  slug = COALESCE(sqlc.narg('slug')::text, slug),
  title = COALESCE(sqlc.narg('title')::text, title),
  thumbnail = COALESCE(sqlc.narg('thumbnail')::text, thumbnail),
  content = COALESCE(sqlc.narg('content')::text, content),
  tags = COALESCE(sqlc.narg('tags')::text[], tags),
  updated_at = COALESCE(sqlc.narg('updated_at')::timestamp, NOW())
WHERE
  deleted_at IS NULL
  AND id = sql.arg('id')
RETURNING
  *;

-- name: DeleteArticles :execrows
UPDATE
  articles
SET
  deleted_at = COALESCE(sqlc.narg('deleted_at')::timestamp, NOW())
WHERE
  deleted_at IS NULL
  AND id = ANY (sql.arg('ids')::integer[]);
