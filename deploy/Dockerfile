# Этот образ собирается заранее, чтобы не подтягивать зависимости каждый раз,
# а использовать уже собранные в этом контейнере
FROM golang:1.22-alpine3.20
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download