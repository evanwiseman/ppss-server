-- name: CreateDevice :one
INSERT INTO devices (name,type)
VALUES ($1,$2)
RETURNING *;

-- name: UpdateDevice :one
UPDATE devices
SET name = $2, type = $3, updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: GetDevices :many
SELECT * FROM devices;

-- name: GetDeviceByID :one
SELECT * FROM devices
WHERE id = $1;

-- name: DeleteDeviceByID :exec
DELETE FROM devices
WHERE id = $1;

-- name: ResetDevices :exec
DELETE FROM devices;
