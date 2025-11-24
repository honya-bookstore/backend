-- name: UpsertReview :exec
INSERT INTO reviews (
  id,
  rating,
  vote_count,
  content,
  user_id,
  book_id,
  created_at,
  updated_at,
  deleted_at
)
VALUES (
  sqlc.arg('id'),
  sqlc.arg('rating'),
  sqlc.arg('vote_count'),
  sqlc.arg('content'),
  sqlc.arg('user_id'),
  sqlc.arg('book_id'),
  sqlc.arg('created_at'),
  sqlc.arg('updated_at'),
  sqlc.narg('deleted_at')
)
ON CONFLICT (id) DO UPDATE SET
  rating = EXCLUDED.rating,
  vote_count = EXCLUDED.vote_count,
  content = EXCLUDED.content,
  user_id = EXCLUDED.user_id,
  book_id = EXCLUDED.book_id,
  created_at = EXCLUDED.created_at,
  updated_at = EXCLUDED.updated_at,
  deleted_at = EXCLUDED.deleted_at;

-- name: ListReviews :many
SELECT
  reviews.*
FROM
  reviews
WHERE
  CASE
    WHEN sqlc.narg('ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.narg('ids')::uuid[]) = 0 THEN TRUE
    ELSE reviews.id = ANY (sqlc.narg('ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.narg('book_ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.narg('book_ids')::uuid[]) = 0 THEN TRUE
    ELSE reviews.book_id = ANY (sqlc.narg('book_ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.narg('user_ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.narg('user_ids')::uuid[]) = 0 THEN TRUE
    ELSE reviews.user_id = ANY (sqlc.narg('user_ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.arg('deleted')::text = 'exclude' THEN reviews.deleted_at IS NULL
    WHEN sqlc.arg('deleted')::text = 'only' THEN reviews.deleted_at IS NOT NULL
    WHEN sqlc.arg('deleted')::text = 'all' THEN TRUE
    ELSE reviews.deleted_at IS NULL
  END
ORDER BY
  reviews.created_at DESC
OFFSET sqlc.arg('offset')::integer
LIMIT NULLIF(sqlc.arg('limit')::integer, 0);

-- name: CountReviews :one
SELECT
  COUNT(*) AS count
FROM
  reviews
WHERE
  CASE
    WHEN sqlc.narg('ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.narg('ids')::uuid[]) = 0 THEN TRUE
    ELSE reviews.id = ANY (sqlc.narg('ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.narg('book_ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.narg('book_ids')::uuid[]) = 0 THEN TRUE
    ELSE reviews.book_id = ANY (sqlc.narg('book_ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.narg('user_ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.narg('user_ids')::uuid[]) = 0 THEN TRUE
    ELSE reviews.user_id = ANY (sqlc.narg('user_ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.arg('deleted')::text = 'exclude' THEN reviews.deleted_at IS NULL
    WHEN sqlc.arg('deleted')::text = 'only' THEN reviews.deleted_at IS NOT NULL
    WHEN sqlc.arg('deleted')::text = 'all' THEN TRUE
    ELSE reviews.deleted_at IS NULL
  END;

-- name: GetReview :one
SELECT
  *
FROM
  reviews
WHERE
  id = sqlc.arg('id')
  AND CASE
    WHEN sqlc.arg('deleted')::text = 'exclude' THEN deleted_at IS NULL
    WHEN sqlc.arg('deleted')::text = 'only' THEN deleted_at IS NOT NULL
    WHEN sqlc.arg('deleted')::text = 'all' THEN TRUE
    ELSE deleted_at IS NULL
  END;

-- name: UpsertReviewVote :exec
INSERT INTO review_votes (
  id,
  user_id,
  review_id,
  is_up,
  created_at
)
VALUES (
  sqlc.arg('id'),
  sqlc.arg('user_id'),
  sqlc.arg('review_id'),
  sqlc.arg('is_up'),
  sqlc.arg('created_at')
)
ON CONFLICT (id) DO UPDATE SET
  user_id = EXCLUDED.user_id,
  review_id = EXCLUDED.review_id,
  is_up = EXCLUDED.is_up,
  created_at = EXCLUDED.created_at;

-- name: ListReviewVotes :many
SELECT
  *
FROM
  review_votes
WHERE
  CASE
    WHEN sqlc.narg('review_ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.narg('review_ids')::uuid[]) = 0 THEN TRUE
    ELSE review_id = ANY (sqlc.narg('review_ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.narg('user_ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.narg('user_ids')::uuid[]) = 0 THEN TRUE
    ELSE user_id = ANY (sqlc.narg('user_ids')::uuid[])
  END
ORDER BY
  created_at DESC;
