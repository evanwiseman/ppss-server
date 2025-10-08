-- name: CreateWdlm :one
INSERT INTO wdlms (
    name
) VALUES (
    $1
)
RETURNING *;

-- name: UpdateWdlm :one
UPDATE wdlms
SET name = $2, updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: GetWdlms :many
SELECT * FROM wdlms;