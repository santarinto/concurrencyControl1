# Network Scanner Tool

A concurrent network scanning utility written in Go for discovering active hosts and services on a local network.

## Features

- **Concurrent Scanning**: Uses goroutines with a semaphore pattern (4 concurrent workers) for efficient network scanning
- **IP Range Scanning**: Scans a configurable range of IPv4 addresses
- **Service Detection**: Detects common services on discovered hosts:
  - SSH (port 22)
  - HTTP (port 80)
  - HTTPS (port 443)
  - MySQL (port 3306)
  - RabbitMQ (port 5672)
  - PostgreSQL (port 5432)
- **Host Discovery**: Performs ICMP ping checks and hostname resolution
- **Real-time Progress**: Shows scanning progress with percentage completion and current IP
- **Formatted Output**: Displays results in a clean table format

## Requirements

- Go 1.16 or higher
- Linux/Unix system (uses `ping` command)
- Network access to the target IP range

## Installation

```bash
# Clone the repository
git clone <repository-url>
cd concurrencyControl1

# Build the binary
go build -o network-scanner

# Or run directly
go run main.go
```

## Configuration

Create a `config.yaml` file in the project root:

```yaml
network:
  start_ip: "192.168.1.1"
  end_ip: "192.168.1.254"
  subnet: "192.168.1.0/24"
```

## Usage

```bash
# Run network scan
./network-scanner scan

# Show help
./network-scanner help
```

## Output

The tool displays a table with the following information for each discovered host:

| Column | Description |
|--------|-------------|
| No. | Sequential number |
| IP Address | IPv4 address of the host |
| Hostname | Resolved hostname (if available) |
| Ping | ICMP ping response (*) |
| SSH | SSH service on port 22 (*) |
| HTTP | HTTP service on port 80 (*) |
| HTTPS | HTTPS service on port 443 (*) |
| MySQL | MySQL database on port 3306 (*) |
| RabbitMQ | RabbitMQ message broker on port 5672 (*) |
| Postgres | PostgreSQL database on port 5432 (*) |

**Note**: An asterisk (*) indicates the service is available/responding.

## How It Works

1. Loads configuration from `config.yaml`
2. Generates IP range based on start and end addresses
3. Scans each IP concurrently using goroutines (limited to 4 concurrent scans)
4. For each host:
   - Performs ICMP ping test
   - Checks for open ports (22, 80, 443)
   - If SSH is available, additionally checks for database ports
   - Resolves hostname if any services are found
5. Displays results in a formatted table

## Project Structure

```
.
├── cmd/
│   └── scan.go           # Scan command implementation
├── internal/
│   └── config/
│       └── config.go     # Configuration management
├── pkg/
│   ├── display/
│   │   └── table.go      # Table formatting utilities
│   └── scanner/
│       └── network.go    # Network scanning logic
├── main.go               # Application entry point
└── config.yaml           # Configuration file
```

## Performance

- Uses concurrent scanning with a semaphore pattern to limit resource usage
- Timeout settings: 1 second for TCP connections, 1 second for ping
- Only includes hosts with SSH, HTTP, or HTTPS services in results
- Database port scanning is performed only for hosts with SSH enabled

## Security Considerations

This tool is designed for **defensive security purposes only**:
- Network discovery and inventory management
- Security auditing of owned networks
- Service availability monitoring
- Infrastructure documentation

**Important**: Only scan networks you own or have explicit permission to scan. Unauthorized network scanning may be illegal in your jurisdiction.

## Dependencies

- `gopkg.in/yaml.v3` - YAML configuration parsing

## License

[Add your license here]
