# Go Collision Anomaly Detector
Concurrent Go REST service that ingests synthetic collision events and flags unusual signatures using a transparent multivariate z-score model.

## Run
```bash
go test ./...
go run ./cmd/server
```
In another terminal:
```bash
go run ./cmd/generator -count 1000
```
Endpoints: `GET /health`, `POST /events`, `GET /stats`, `GET /anomalies`.

