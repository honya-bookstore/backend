-- name: UpsertOrder :exec
INSERT INTO orders (
  id,
  user_id,
  address,
  total_amount,
  is_paid,
  provider_id,
  status_id,
  created_at,
  updated_at
) VALUES (
  sqlc.arg('id'),
  sqlc.arg('user_id'),
  sqlc.arg('address'),
  sqlc.arg('total_amount'),
  sqlc.arg('is_paid'),
  sqlc.arg('provider_id'),
  sqlc.arg('status_id'),
  sqlc.arg('created_at'),
  sqlc.arg('updated_at')
)
ON CONFLICT (id) DO UPDATE SET
  user_id = EXCLUDED.user_id,
  address = EXCLUDED.address,
  total_amount = EXCLUDED.total_amount,
  is_paid = EXCLUDED.is_paid,
  provider_id = EXCLUDED.provider_id,
  status_id = EXCLUDED.status_id,
  created_at = EXCLUDED.created_at,
  updated_at = EXCLUDED.updated_at;

-- name: ListOrders :many
SELECT
  *
FROM
  orders
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
    WHEN sqlc.narg('status_ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.narg('status_ids')::uuid[]) = 0 THEN TRUE
    ELSE status_id = ANY (sqlc.narg('status_ids')::uuid[])
  END
ORDER BY
  id ASC
OFFSET sqlc.arg('offset')::integer
LIMIT NULLIF(sqlc.arg('limit')::integer, 0);

-- name: CountOrders :one
SELECT
  COUNT(*) AS count
FROM
  orders
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
    WHEN sqlc.narg('status_ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.narg('status_ids')::uuid[]) = 0 THEN TRUE
    ELSE status_id = ANY (sqlc.narg('status_ids')::uuid[])
  END;

-- name: GetOrder :one
SELECT
  *
FROM
  orders
WHERE
  id = sqlc.arg('id');

-- name: ListOrderItems :many
SELECT
  *
FROM
  order_items
WHERE
  CASE
    WHEN sqlc.narg('ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.narg('ids')::uuid[]) = 0 THEN TRUE
    ELSE id = ANY (sqlc.narg('ids')::uuid[])
  END
  AND CASE
    WHEN sqlc.narg('order_ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.narg('order_ids')::uuid[]) = 0 THEN TRUE
    ELSE order_id = ANY (sqlc.narg('order_ids')::uuid[])
  END
ORDER BY
  id;

-- name: GetOrderItem :one
SELECT
  *
FROM
  order_items
WHERE
  id = sqlc.arg('id');

-- name: ListOrderStatuses :many
SELECT
  *
FROM
  order_statuses
WHERE
  CASE
    WHEN sqlc.narg('ids')::uuid[] IS NULL THEN TRUE
    WHEN cardinality(sqlc.narg('ids')::uuid[]) = 0 THEN TRUE
    ELSE id = ANY (sqlc.narg('ids')::uuid[])
  END
ORDER BY
  id ASC;

-- name: GetOrderStatus :one
SELECT
  *
FROM
  order_statuses
WHERE
  CASE
    WHEN sqlc.narg('id')::uuid IS NULL THEN TRUE
    ELSE id = sqlc.narg('id')::uuid
  END
  AND CASE
    WHEN sqlc.narg('name')::text IS NULL THEN TRUE
    ELSE name = sqlc.narg('name')::text
  END;

-- name: GetOrderProvider :one
SELECT
  *
FROM
  order_providers
WHERE
  CASE
    WHEN sqlc.narg('id')::uuid IS NULL THEN TRUE
    ELSE id = sqlc.narg('id')::uuid
  END
  AND CASE
    WHEN sqlc.narg('name')::text IS NULL THEN TRUE
    ELSE name = sqlc.narg('name')::text
  END;

-- name: CreateTempTableOrderItems :exec
CREATE TEMPORARY TABLE temp_order_items (
  id UUID PRIMARY KEY,
  quantity INTEGER NOT NULL,
  order_id UUID NOT NULL,
  price DECIMAL(12, 0) NOT NULL,
  book_id UUID NOT NULL
) ON COMMIT DROP;

-- name: InsertTempTableOrderItems :copyfrom
INSERT INTO temp_order_items (
  id,
  quantity,
  order_id,
  price,
  book_id
) VALUES (
  @id,
  @quantity,
  @order_id,
  @price,
  @book_id
);

-- name: MergeOrderItemsFromTemp :exec
MERGE INTO order_items AS target
USING temp_order_items AS source
  ON target.id = source.id
WHEN MATCHED THEN
  UPDATE SET
    quantity = source.quantity,
    order_id = source.order_id,
    price = source.price,
    book_id = source.book_id
WHEN NOT MATCHED THEN
  INSERT (
    id,
    quantity,
    order_id,
    price,
    book_id
  )
  VALUES (
    source.id,
    source.quantity,
    source.order_id,
    source.price,
    source.book_id
  )
WHEN NOT MATCHED BY SOURCE
  AND target.order_id IN (SELECT DISTINCT order_id FROM source) THEN
  DELETE;
