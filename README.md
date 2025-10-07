# PPSS Server
The PPSS server allows for local connections via a network switch to many raspberry pi's and to local GUIs. While allowing for public connections over https where users can interact with the raspberry pi's.

## Installation
TODO

## Setup Repository
### Go v1.25.0
Official Installer: [https://go.dev/dl/](https://go.dev/dl/)

| OS      | Command / Notes |
|---------|----------------|
| macOS   | `brew install go` |
| Linux   | `sudo snap install --classic go` |
| Windows | Download and install `.msi` from Go website. Add Go to your PATH. |

Verify installation:

```bash
go version
```

### Postgres 15.14
Offical Installer: [(https://www.postgresql.org/download/)](https://www.postgresql.org/download/)

| OS        | Command / Notes |
|-----------|-----------------|
| macOS     | `brew install postgresql` <br> `brew services start postgresql` |
| Linux     | `sudo apt update && sudo apt install postgresql postgresql-contrib` <br> `sudo systemctl enable --now postgresql` |
| Windows   | Use the installer from the website. Use pgAdmin or psql CLI. |

Verify installation:
```bash
psql --version
```

### Goose (Database Migrations)
| OS        | Command / Notes |
|-----------|-----------------|
| macOS     | `brew install goose` |
| Linux     | `go install github.com/pressly/goose/v3/cmd/goose@latest` |
| Windows   | `go install github.com/pressly/goose/v3/cmd/goose@latest` |

Verify installation:
```bash
goose -v
```

**Ensure Goose is in your $PATH**

### SQLC (Generate Type-Safe Go Queries)
Install on any OS:
```
# Go install
go install github.com/kyleconroy/sqlc/cmd/sqlc@latest

# check version
sqlc version
```

**Make sure $GOPATH/bin (or %GOPATH%\bin on Windows) is in your PATH.**

## Server
There exists 2 databases and 2 places to serve endpoints from, locally and publicly.

### Local
Serves endpoints and queries from/for the raspberry pi's. This will likely predominately hold raspberry pi configuration and settings, such as devices, their name, ip address, and sensors associated. This will not hold stateful data such as a sensor value.

### Public
Serves endpoints and queries from/for users accessing the system from a webpage or nonlocally i.e. from home or a device not connected via a network switch to the host.

Goal: Users are validated and logged in via their work email and password and validated on the work network. A user would likely need to use the work VPN.

## API

### Local

#### Template
**Behavior:**

**Request Body:**

**Response (code):**

#### PostDeviceHandler
Creates a device in the devices table

**Behavior:**
- Unmarshals device parameters from the request body
- Attempts to create a record in the database
- Returns the record created

**Request Body:**
```json
{
    "serial_number": "00000000abcdef01",
    "name": "Test Pi",
    "ip_address": "192.168.1.10",
    "device_type": "raspberry_pi"
}
```

**Response (201 Created):**
```json
{
    "serial_number": "00000000abcdef01",
    "name": "Test Pi",
    "ip_address": "192.168.1.10",
    "device_type": "raspberry_pi",
    "created_at": "2025-10-06T11:30:45Z",
    "updated_at": "2025-10-06T11:30:45Z",
    "last_seen": null
}
```

**Response (400 Bad Request):**
- JSON is invalid

**Response (409 Conflict):**
- Device already exists in the database

**Response (500 Internal Server Error):**
- Unable to create record in database
- Fail to send a response of the created device

#### DeleteDeviceByIDHandler
Deletes a device with a given ID

**Behavior:**
- Parses the device ID in the URL `/devices/{deviceID}`
- Checks if the device is in the database
- Deletes the device with the ID
- Returns no content

**Request Body:**
```
<empty>
```

**Response (204 No Content):**
```
<empty>
```

**Response (400 Bad Request):**
- Invalid ID is provided in URL

**Response (404 Not Found):**
- No device with ID is found in the database

**Response (500 Internal Server Error):**
- Unable to delete the device in the database

#### GetDevicesHandler
Gets all devices in the database

**Behavior:**
- Queries for all devices in the database
- Formats a response as a slice of devices in json format
- Returns a list of all devices

**Request Body:**
```
<empty>
```

**Response (200 Ok):**
```json
{
    {
        "serial_number": "00000000abcdef01",
        "name": "Test Pi 1",
        "ip_address": "192.168.1.10",
        "device_type": "raspberry_pi",
        "created_at": "2025-10-06T11:30:45Z",
        "updated_at": "2025-10-06T11:30:45Z",
        "last_seen": null
    },
    {
        "serial_number": "00000000abcdef02",
        "name": "Test Pi 2",
        "ip_address": "192.168.1.11",
        "device_type": "raspberry_pi",
        "created_at": "2025-10-06T11:30:45Z",
        "updated_at": "2025-10-06T11:30:45Z",
        "last_seen": null
    }
}
```

**Response (404 Not Found):**
- The devices database is not found

#### GetDeviceByIDHandler
Gets the device matching an ID

**Behavior:**
- Parses the device ID in the URL `/devices/{deviceID}`
- Find the device in the database
- Return the device as a JSON

**Request Body:**
```
<empty>
```

**Response (200 Ok):**
```json
{
    "serial_number": "00000000abcdef01",
    "name": "Test Pi 1",
    "ip_address": "192.168.1.10",
    "device_type": "raspberry_pi",
    "created_at": "2025-10-06T11:30:45Z",
    "updated_at": "2025-10-06T11:30:45Z",
    "last_seen": null
}
```

**Response (400 Bad Request):**
- Invalid ID is provided in URL

**Response (404 Not Found):**
- The devices database is not found

#### ResetDevicesHandler
Clears the entire devices table

**Behavior:**
- Check the platform before reseting
- Reset the device table

**Request Body:**
```
<empty>
```

**Response (401 Unauthorized):**
- When the platform is not = "dev" in the config

**Response (500 Internal Server Error):**
- Unable to reset the devices table