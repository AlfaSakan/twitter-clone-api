FROM golang:1.18-alpine AS builder

WORKDIR /app

COPY src ./src
COPY go.mod go.sum main.go ./

RUN go build -o my-profile

FROM golang:1.18-alpine AS runner

WORKDIR /app

COPY --from=builder /app/my-profile ./

EXPOSE 8081

CMD ./my-profile
