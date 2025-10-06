-- +goose Up
CREATE TABLE devices (
    serial_number TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    ip_address TEXT NOT NULL,
    device_type TEXT NOT NULL, 
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    last_seen TIMESTAMP
);

-- +goose Down
DROP TABLE devices;