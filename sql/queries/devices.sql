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

-- name: UpdateDevice :one
UPDATE devices
SET name = $2, ip_address = $3, device_type = $4, updated_at = NOW(), last_seen = NOW()
WHERE id = $1
RETURNING *;


-- name: DeleteDeviceByID :exec
DELETE FROM devices
WHERE id = $1;

-- name: GetDevices :many
SELECT * FROM devices;

-- name: GetDeviceByID :one
SELECT * FROM devices
WHERE id = $1;

-- name: ResetDevices :exec
DELETE FROM devices;
