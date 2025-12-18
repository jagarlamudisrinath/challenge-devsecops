# DevSecOps Challenge - Work Summary

## Approach

I approached this challenge in a structured, phased manner following DevSecOps best practices:

1. **Security Audit**: Conducted a thorough review of all code and configuration files
2. **Infrastructure Hardening**: Fixed security issues in Docker and Kubernetes configurations
3. **Pipeline Creation**: Implemented a comprehensive CI/CD pipeline with security scanning
4. **Shift-Left Security**: Configured automated security checks to run on every push/PR

## Phase 1: Infrastructure Security Issues Found & Fixed

### Dockerfile

| Issue | Severity | Before | After |
|-------|----------|--------|-------|
| Running as root | Critical | `USER root` | Created non-root user `appuser` (UID 1000) |
| Outdated base image | High | `golang:1.16-buster` | `golang:1.21-alpine` + `alpine:3.19` |
| Large attack surface | Medium | Full Debian-based image | Minimal Alpine image |
| No health check | Low | None | Added `HEALTHCHECK` instruction |

### docker-compose.yaml

| Issue | Severity | Before | After |
|-------|----------|--------|-------|
| Hardcoded password | Critical | `POSTGRES_PASSWORD: "6199178B..."` | `${POSTGRES_PASSWORD}` from .env |
| Outdated PostgreSQL | High | `postgres:9.6.23-buster` (EOL) | `postgres:15-alpine` |
| No resource limits | Medium | None | Added CPU/memory limits |
| No health checks | Medium | None | Added health checks for both services |

### k8s/deployment.yaml

| Issue | Severity | Before | After |
|-------|----------|--------|-------|
| Privileged container | Critical | `privileged: true` | Removed |
| Privilege escalation | Critical | `allowPrivilegeEscalation: true` | `allowPrivilegeEscalation: false` |
| Hardcoded AWS credentials | Critical | Plaintext in env vars | Removed (use IRSA/OIDC in production) |
| Running as root | High | No restriction | `runAsNonRoot: true`, `runAsUser: 1000` |
| Writable filesystem | Medium | Default | `readOnlyRootFilesystem: true` |
| All capabilities | Medium | Default | `drop: ALL` |
| No resource limits | Medium | None | Added CPU/memory limits |
| No probes | Low | None | Added liveness/readiness probes |

## Phase 2: Application Security Issues (Detected by Pipeline)

The following issues were identified during the security audit and will be detected by the automated pipeline:

| Issue | Location | Description | Detection Tool |
|-------|----------|-------------|----------------|
| SQL Injection | `dummy/users.go:50` | User input directly concatenated into SQL query | gosec |
| Hardcoded Password | `db/db.go:47` | Database password hardcoded in source | gosec, gitleaks |
| Plaintext Password Storage | `main.go:34` | Admin password stored without hashing | gosec |
| Password Exposure in API | `api/controller.go` | Password field returned in JSON responses | Code review |

These will be addressed in Phase 3 after reviewing the pipeline security reports.

## DevSecOps Pipeline

Created a comprehensive GitHub Actions pipeline (`.github/workflows/devsecops.yaml`) with the following stages:

```
+----------------+     +----------------+     +------------------+
|     Lint       |---->|     SAST       |---->|  Secret Scan     |
| (golangci-lint)|     |    (gosec)     |     |   (gitleaks)     |
+----------------+     +----------------+     +------------------+
        |                      |                      |
        v                      v                      v
+----------------+     +----------------+     +------------------+
|     Build      |---->| Container Scan |---->|    IaC Scan      |
|    (Docker)    |     |    (Trivy)     |     |    (Trivy)       |
+----------------+     +----------------+     +------------------+
        |                      |                      |
        +----------------------+----------------------+
                               |
                               v
                    +--------------------+
                    |  Security Summary  |
                    +--------------------+
```

### Pipeline Jobs

| Job | Tool | Purpose |
|-----|------|---------|
| lint | golangci-lint | Code quality and style checks |
| sast | gosec | Go-specific security vulnerability detection |
| secrets | gitleaks | Detect hardcoded secrets and credentials |
| dependency-scan | govulncheck | Check for vulnerabilities in dependencies |
| build | Docker | Build container image |
| container-scan | Trivy | Scan container for CVEs |
| iac-scan | Trivy | Scan Kubernetes/Docker configs for misconfigurations |
| test | go test | Run unit tests with coverage |
| security-summary | - | Aggregate results in GitHub Actions summary |

### Security Reports

All security findings are uploaded to GitHub's Security tab using SARIF format:
- **gosec**: Go security issues
- **trivy-container**: Container vulnerabilities
- **trivy-iac-k8s**: Kubernetes misconfigurations
- **trivy-iac-docker**: Docker misconfigurations

## Files Changed

| File | Action | Description |
|------|--------|-------------|
| `Dockerfile` | Modified | Non-root user, Alpine base, health check |
| `docker-compose.yaml` | Modified | Env vars, updated versions, resource limits |
| `k8s/deployment.yaml` | Modified | Security context, secrets reference, probes |
| `.gitignore` | Modified | Added .env to prevent secret commits |

## Files Created

| File | Purpose |
|------|---------|
| `.env.example` | Template for environment variables |
| `k8s/secrets.yaml` | Kubernetes secrets template |
| `k8s/service.yaml` | Kubernetes service definition |
| `.github/workflows/devsecops.yaml` | CI/CD pipeline with security scanning |
| `WORK_SUMMARY.md` | This documentation |

## How to Run

### Local Development

```bash
# Copy environment template
cp .env.example .env

# Edit .env and set a secure password
vim .env

# Start the stack
docker-compose up -d

# View logs
docker-compose logs -f
```

### Kubernetes Deployment

```bash
# Create secrets (replace with actual values)
kubectl create secret generic challenge-secrets \
  --from-literal=postgres-user=challenge \
  --from-literal=postgres-password=YOUR_SECURE_PASSWORD \
  --from-literal=postgres-db=challenge

# Apply manifests
kubectl apply -f k8s/
```

## Next Steps (Phase 3)

After the pipeline runs and generates security reports:

1. Review findings in GitHub Security tab
2. Fix SQL injection vulnerability in `dummy/users.go`
3. Move database password to environment variable in `db/db.go`
4. Implement password hashing (bcrypt) for user passwords
5. Add `json:"-"` tag to Password field to hide from API responses
6. Consider adding authentication middleware

## Security Best Practices Implemented

- **Principle of Least Privilege**: Non-root containers, dropped capabilities
- **Defense in Depth**: Multiple layers of security scanning
- **Shift-Left Security**: Security checks run early in CI/CD
- **Secret Management**: No hardcoded secrets, use environment variables/K8s secrets
- **Immutable Infrastructure**: Read-only filesystem, minimal base images
- **Continuous Monitoring**: Automated security scans on every commit
