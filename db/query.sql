-- name: GetAllTask :many
SELECT * FROM Task;

-- name: GetTaskByCategory :many
SELECT * FROM Task
WHERE category_type = $1;

-- name: CreateTask :one
INSERT INTO Task (
  category_type,
  last_update_by,
  price
) VALUES (
  $1, $2, $3
)
RETURNING *;
