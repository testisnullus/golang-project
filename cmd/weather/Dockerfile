FROM golang:1.20-alpine3.17 as builder

COPY . /app

WORKDIR /app/cmd/weather

RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -o /cmd/weather main.go

FROM alpine:3.17

COPY --from=builder /cmd/weather /bin/cmd/weather

ENTRYPOINT [ "/bin/cmd/weather" ]