-- +goose Up
CREATE TABLE devices (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,                         -- human friendly name
    serial_number TEXT UNIQUE,                  -- serial number of device
    ip_address TEXT,                            -- current ip of device
    firmware_version TEXT,                      -- current firmware version
    device_type TEXT DEFAULT 'raspberry_pi', 
    last_seen TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE devices;