# Climinal Server

A secure messaging server supporting end-to-end encryption and centralized modes.

## Features

- End-to-end encryption support
- Centralized server mode
- File attachment handling
- Configurable message retention
- Secure storage

## Installation

```bash
go install github.com/climinal/server@latest
```

## Configuration

Configuration is stored in `~/.config/climinal/config.yml`. Example configuration:

```yaml
port: 8080
encryption_mode: "e2e"
log_retention_days: 30
storage_path: "./storage"
```

## Usage

Run the server:

```bash
climinal-server
```

## License

MIT License 