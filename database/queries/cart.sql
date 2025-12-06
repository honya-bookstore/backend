-- name: UpsertCart :exec
INSERT INTO carts (
  id,
  user_id,
  updated_at
) VALUES (
  sqlc.arg('id'),
  sqlc.arg('user_id'),
  sqlc.arg('updated_at')
)
ON CONFLICT (id) DO UPDATE SET
  user_id = EXCLUDED.user_id,
  updated_at = EXCLUDED.updated_at;

-- name: GetCart :one
SELECT
  *
FROM
  carts
WHERE
  CASE
    WHEN sqlc.arg('id')::uuid IS NULL THEN TRUE
    WHEN sqlc.arg('id')::uuid = '00000000-0000-0000-0000-000000000000'::uuid THEN TRUE
    ELSE id = sqlc.arg('id')::uuid
  END
  AND CASE
    WHEN sqlc.arg('user_id')::uuid IS NULL THEN TRUE
    WHEN sqlc.arg('user_id')::uuid = '00000000-0000-0000-0000-000000000000'::uuid THEN TRUE
    ELSE user_id = sqlc.arg('user_id')::uuid
  END;

-- name: ListCartItems :many
SELECT
  *
FROM
  cart_items
WHERE
  CASE
    WHEN sqlc.arg('cart_id')::uuid IS NULL THEN TRUE
    WHEN sqlc.arg('cart_id')::uuid = '00000000-0000-0000-0000-000000000000'::uuid THEN TRUE
    ELSE cart_id = sqlc.arg('cart_id')::uuid
  END;

-- name: CreateTempTableCartItems :exec
CREATE TEMPORARY TABLE temp_cart_items (
  id UUID PRIMARY KEY,
  quantity INTEGER NOT NULL,
  cart_id UUID NOT NULL,
  book_id UUID NOT NULL
) ON COMMIT DROP;

-- name: InsertTempTableCartItems :copyfrom
INSERT INTO temp_cart_items (
  id,
  quantity,
  cart_id,
  book_id
) VALUES (
  @id,
  @quantity,
  @cart_id,
  @book_id
);

-- name: MergeCartItemsFromTemp :exec
MERGE INTO cart_items AS target
USING temp_cart_items AS source
  ON target.id = source.id
WHEN MATCHED THEN
  UPDATE SET
    quantity = source.quantity,
    cart_id = source.cart_id,
    book_id = source.book_id
WHEN NOT MATCHED THEN
  INSERT (
    id,
    quantity,
    cart_id,
    book_id
  )
  VALUES (
    source.id,
    source.quantity,
    source.cart_id,
    source.book_id
  )
WHEN NOT MATCHED BY SOURCE
  AND target.cart_id IN (SELECT DISTINCT cart_id FROM temp_cart_items) THEN
  DELETE;

-- name: DeleteCartItems :exec
DELETE FROM
  cart_items
WHERE
  cart_id = sqlc.arg('cart_id');