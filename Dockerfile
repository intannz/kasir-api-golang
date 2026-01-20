# Ganti jadi versi latest (alpine saja)
FROM golang:alpine

# Set folder kerja
WORKDIR /app

# Copy semua file
COPY . .

# Build aplikasi
RUN go build -o main .

# Hugging Face Spaces WAJIB pakai port 7860
ENV PORT=7860

# Buka port 7860
EXPOSE 7860

# Jalankan aplikasi
CMD ["./main"]