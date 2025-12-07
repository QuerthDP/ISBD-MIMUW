# Public Interface Tests (PIT)

Test suite for validating the database public interface.

## Running Tests

By default, tests connect to an existing database running on `localhost:8080` (Docker is not started).

### Basic Usage

```bash
# Default behavior - connects to localhost:8080
go test ./tests -v
```

### Configuration Options

Tests can be configured via command-line flags or environment variables. Flags take precedence over environment variables.

#### Command-line Flags

- `--db-image`: Docker image name to use (env: `DB_IMAGE`, default: `isbd-mimuw-db:latest`)
- `--db-hostname`: Hostname of running database (env: `DB_HOSTNAME`, default: `localhost`)
- `--db-port`: Port on which database listens (env: `DB_PORT`, default: `8080`)
- `--db-run-docker`: Skip Docker container and use existing database (env: `DB_RUN_DOCKER`, default: `true`)

#### Environment Variables

Set these to configure test behavior:

```bash
DB_HOSTNAME=db.example.com          # Database hostname
DB_PORT=5432                        # Database port
DB_RUN_DOCKER=true                  # Start Docker container instead of connecting to existing
DB_IMAGE=my-custom-db:latest        # Custom Docker image (makes sense only with DB_RUN_DOCKER=true)
```

### Usage Examples

```bash
# Connect to custom hostname and port
go test ./tests -db-hostname 192.168.1.100 -db-port 5432 -v

# Start Docker with a custom image
go test ./tests -db-run-docker true -db-image my-custom-db:v1 -v

# Use environment variables
DB_HOSTNAME=mydb.local DB_PORT=8080 go test ./tests -v

# Run only specific tests by name (standard Go -run flag)
go test ./tests -run "SystemInfo" -v

# Combine test filtering with custom database config
go test ./tests -run "TableCreation/TableEmpty" -db-hostname mydb.local -db-port 8080 -v
```

### Docker Container

To start Docker automatically, set `--db-run-docker true` or `DB_RUN_DOCKER=true`:

```bash
# This will start the Docker container defined by DB_IMAGE
go test ./tests -db-run-docker true -v
```