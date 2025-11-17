-- name: CreateCategory :one
INSERT INTO categories (
  slug,
  name,
  description,
  created_at
) VALUES (
  @slug,
  @name,
  @description,
  COALESCE(sqlc.narg('created_at')::timestamp, NOW())
);

-- name: ListCategories :many
SELECT
  *,
  COUNT(*) OVER() AS current_count,
  COUNT(*) AS total_count
FROM
  categories
WHERE
  CASE
    WHEN sqlc.narg('search')::text IS NULL THEN TRUE
    ELSE name ||| (sqlc.narg('search')::text)::pdb.fuzzy(2)
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
    WHEN @deleted::text = 'exclude' THEN deleted_at IS NOT NULL
    WHEN @deleted::text = 'only' THEN deleted_at IS NULL
    WHEN @deleted::text = 'all' THEN TRUE
    ELSE FALSE
  END
ORDER BY
  CASE WHEN
    sqlc.narg('search') IS NOT NULL THEN pdb.score(id)
  END DESC,
  id DESC
OFFSET COALESCE(sqlc.narg('offset')::integer, 0)
LIMIT COALESCE(sqlc.narg('limit')::integer, 20);

-- name: UpdateCategory :one
UPDATE
  categories
SET
  slug = COALESCE(sqlc.narg('slug')::text, slug),
  name = COALESCE(sqlc.narg('name')::text, name),
  description = COALESCE(sqlc.narg('description')::text, description),
  updated_at = COALESCE(sqlc.narg('updated_at')::timestamp, NOW())
WHERE
  deleted_at IS NULL
  AND id = @id
RETURNING
  *;

-- name: DeleteCategory :execrows
UPDATE
  categories
SET
  deleted_at = COALESCE(sqlc.narg('deleted_at')::timestamp, NOW())
WHERE
  deleted_at IS NULL
  AND id = ANY (@ids::integer[]);
