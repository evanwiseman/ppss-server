-- name: CreateWdlm :one
INSERT INTO wdlms (
    name
) VALUES (
    $1
)
RETURNING *;