# backend-app/Dockerfile

FROM golang:1.21.3

WORKDIR /app

# 将go mod和go sum文件复制并下载所有依赖项
COPY go.mod go.sum ./
RUN GOPROXY=https://goproxy.cn go mod download

# 将Go代码复制到容器中
COPY . .

# 编译Go应用
RUN go build -o main .

# 暴露端口8000
EXPOSE 8000

CMD ["./main"]

