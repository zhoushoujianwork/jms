# AGENTS.md

## Cursor Cloud specific instructions

### Project overview
JMS is a Go-based SSH bastion host / jump server with three runtime modes: `sshd` (SSH proxy), `api` (REST management), `scheduler` (background tasks). See `README.md` for user-facing docs.

### Running services locally
- **Config**: The app reads from `/opt/jms/config.yaml` by default. A sample is at `/workspace/config.yaml`. Copy it and adjust (set `withDB.enable: false` for minimal mode with default `jms/jms` auth).
- **Directories**: Ensure `/opt/jms/`, `/opt/jms/.ssh/`, `/opt/jms/logs/`, `/opt/jms/audit/` exist with write permission.
- **SSHD**: `./jms sshd --port 22222 --config /opt/jms/config.yaml --debug`
- **API**: `./jms api --port 8013 --config /opt/jms/config.yaml --debug` — Swagger at `http://localhost:8013/swagger/index.html`
- **Scheduler**: `./jms scheduler --config /opt/jms/config.yaml --debug` — requires `withDB.enable: true`

### Testing
- `go vet ./...` for lint.
- `go test ./model/...` passes without external dependencies.
- `go test ./core/... -run TestAuditArch` passes without DB.
- Most other tests (`app`, `core/db`, `core/dingtalk`, `io`) are integration tests that require a real PostgreSQL or DingTalk API. They will fail in environments without those services.
- Tests require `/opt/jms/config.yaml` and `/opt/jms/.ssh/authorized_keys` to exist.

### Build
- `go build -o jms -ldflags "-X main.version=$(date +%Y%m%d)" ./main.go`
- Binary supports subcommands: `sshd`, `api`, `scheduler`, `db`.

### Non-obvious notes
- The default config uses `withDB.enable: false`, which means the SSH server works in standalone mode with `jms/jms` as the only credential. No PostgreSQL required.
- pprof server starts on `:6060` in all modes — when running multiple modes locally, only the first will bind successfully (non-fatal, just a warning in logs).
- The `core/sshd/sshd.go` SSH client config disables host key verification (`HostKeyCallback` returns nil) — expected for a bastion host connecting to many servers.
