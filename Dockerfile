# 使用官方的 Golang 镜像作为基础镜像
FROM golang:1.16 AS builder

# 设置工作目录
WORKDIR /app

# 将 go.mod 和 go.sum 文件复制到容器中
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 将项目源代码复制到容器中
COPY . .

# 编译项目
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o my_binary .

# 使用一个较小的基础镜像来减小最终镜像的大小
FROM alpine:latest

# 将编译好的二进制文件从 builder 镜像复制到当前镜像
COPY --from=builder /app/my_binary /app/my_binary

# 将配置文件从项目目录复制到当前镜像
COPY config /config

# 设置环境变量 CONFIG_PATH
ENV CONFIG_PATH=/config

# 设置工作目录
WORKDIR /app

# 暴露应用程序使用的端口
EXPOSE 8080

# 运行应用程序
CMD ["/app/my_binary"]