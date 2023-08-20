FROM golang:1.19 AS builder
WORKDIR /project
ADD . .
# 编译：main.go编译成可执行的二进制文件，命名为main
RUN go env -w GO111MODULE=on && \
    go env -w GOPROXY=https://goproxy.cn,direct &&\
    GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -ldflags="-s -w" -installsuffix cgo -o main main.go

FROM golang:1.19
WORKDIR /app
ADD config config
COPY --from=builder /project/main main
CMD ["/app/main"]