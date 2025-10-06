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

