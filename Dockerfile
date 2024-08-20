FROM golang:1.22-alpine

WORKDIR /app

# Menyalin go.mod dan go.sum untuk dependency
COPY go.mod go.sum ./

# Mengunduh dependensi Go
RUN go mod download
RUN go install -tags "postgres" github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Menyalin seluruh kode sumber ke dalam container
COPY . .

# Mengkompilasi kode Go
RUN go clean
RUN go build -o main ./cmd

# Mengekspos port yang digunakan oleh aplikasi
EXPOSE 8081