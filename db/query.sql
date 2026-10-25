-- name: GetAllTask :many
SELECT * FROM Task;

-- name: GetTaskByCategory :many
SELECT * FROM Task
GROUP BY CATEGORY_TYPE;

-- name: CreateTask :one
INSERT INTO Task (
  taskId,
  category_type,
  last_update_by,
  price
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;
