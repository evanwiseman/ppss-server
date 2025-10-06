-- name: CreateDevice :one
INSERT INTO devices (
    name,
    serial_number,
    ip_address,
    firmware_version,
    device_type
) VALUES (
    $1,
    $2,
    $3,
    $4,
    COALESCE($5, DEFAULT)
)
RETURNING *;