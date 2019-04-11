FROM golang as builder
LABEL summary="App go Server" \
      name="App go Server" \
      version="0.0.1" \
      maintainer="tonkung <ton_naruto_v.2@hotmail.com>"
# ให้ใช้ go module ได้
ENV GO111MODULE=on

WORKDIR /app

# ก็อปไปเตรียมดาโหลดไฟล์
COPY go.mod .
COPY go.sum .
# ดาวโหลดไฟล์
RUN go mod download
# ก็อปทั้งหมดตามมา
COPY . .
#  ตั้งค่าและ build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

FROM golang:1.12-alpine3.9
# ก็อบไฟล์จากด้านบน
#  test-serve คือชื่อตอนกด go mod init test-serve
COPY --from=builder /app/test-serve /app/  
# เปิดพอท 8080
EXPOSE 8080
# รันไฟล์นี้
ENTRYPOINT ["/app/test-serve"]