-- name: CreateDevice :one
INSERT INTO devices (
    name,
    ip_address,
    device_type
) VALUES (
    $1,
    $2,
    $3
)
RETURNING *;

-- name: DeleteDevice :exec
DELETE FROM devices
WHERE id = $1;