# Gunakan image Go resmi sebagai image dasar
FROM golang:alpine AS builder

# Set direktori kerja di dalam container
WORKDIR /app

# Salin file kode sumber Go ke dalam container
COPY . .

# Build aplikasi Go
RUN go build -o main .

# Gunakan alpine terbaru untuk ukuran image yang kecil
FROM alpine:latest  
WORKDIR /root/

# Salin executable dari builder ke direktori root
COPY --from=builder /app/main .

# Ekspose port 8080
EXPOSE 8081

# Jalankan executable
CMD ["./main"]
