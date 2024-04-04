FROM golang:1.22 AS stage1
WORKDIR /app

COPY . ./
RUN go mod download
RUN GO111MODULE="on" CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ./cmd/app/main.go

FROM alpine:latest as stage2
WORKDIR /app
COPY --from=stage1 /app/app .
COPY --from=stage1 /app/.env.local .

CMD ["./app"]