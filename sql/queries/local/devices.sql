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