# secret-decoder
Automatically decode base64-encoded fields from Kubernetes secrets

## Install
```
go install github.com/mikesmitty/secret-decoder@latest
```

## Usage
```
kubectl get secret your-secret -o yaml | secret-decoder
```
