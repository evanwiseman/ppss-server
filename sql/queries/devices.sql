-- name: CreateDevice :one
INSERT INTO devices (
    serial_number,
    name,
    ip_address,
    device_type
) VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING *;

-- name: DeleteDevice :exec
DELETE FROM devices
WHERE serial_number = $1;