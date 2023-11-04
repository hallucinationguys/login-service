FROM golang:1.20-alpine as builder 

WORKDIR /build

COPY go.mod ./
RUN go mod download && go mod verify

COPY . .
ENV GOOS=linux CGO_ENABLED=0
RUN set -ex && \
    apk add --no-progress --no-cache \
    gcc \
    musl-dev

RUN go build -o app ./main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates bash 
WORKDIR /
COPY --from=builder /build/app ./
COPY --from=builder /build/production.env ./
COPY --from=builder /build/database/postgres/migration ./database/postgres/migration

CMD ["./app"]