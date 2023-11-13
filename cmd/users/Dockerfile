FROM golang:1.20-alpine3.17 as builder

COPY . /app

WORKDIR /app/cmd/users

RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -o /cmd/users main.go

FROM alpine:3.17

COPY --from=builder /cmd/users /bin/cmd/users

ENTRYPOINT [ "/bin/cmd/users" ]