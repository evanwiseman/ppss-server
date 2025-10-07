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

-- name: DeleteDeviceByID :exec
DELETE FROM devices
WHERE id = $1;

-- name: GetDevices :many
SELECT * FROM devices;

-- name: GetDeviceByID :one
SELECT * FROM devices
WHERE id = $1;