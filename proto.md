## Microservice Proto Update and Regeneration

Use this when proto contracts change and you need fresh Go gRPC files.

Prerequisite: Go must be installed (`go` available in `PATH`).

### 1. Update proto submodule

From microservice root:

```bash
git submodule update --init --remote --merge proto
```


### 2. Regenerate Go files

From microservice root:

```bash
make proto
```

This generates:
- `internal/pb/common/common.pb.go`
- `internal/pb/symmetric/symmetric.pb.go`
- `internal/pb/symmetric/symmetric_grpc.pb.go`

### 3. Verify

```bash
git diff -- internal/pb
```

