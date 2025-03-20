# Dung Golang lam moi truong build
FROM golang:1.23 AS builder

# Dat thuc muc lam viec trong container
WORKDIR /app

# Copy module files và tai dependency
COPY go.mod go.sum ./
RUN go mod download

# Copy tat ca code vao container
COPY . .

# Lay het cac dependency
RUN go mod tidy

# Build ung dung Golang
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ip_management main.go

# Giai doan chay ung dung (dung image nhe hon)
FROM alpine:latest

WORKDIR /root/

# Cai dat MySQL client de kiem tra connection neu can
RUN apk add --no-cache mysql-client

# Copy binary tu giai doan build
COPY --from=builder /app/ip_management /root/ip_management
# Copy file .env neu co
COPY .env .env

# Chay ung dung
CMD ["/app/ip_management"]
