# 设置基础镜像
FROM golang:1.17-alpine

# 设置工作目录
WORKDIR /app

# 复制应用程序源代码到容器中
COPY main.go .

# 构建应用程序
RUN go build -o app main.go

# 暴露应用程序端口
EXPOSE 8080

# 运行应用程序
CMD ["./app"]


# 在上面的 Dockerfile 中，我们使用了官方的 Go 语言 Docker 镜像 golang:1.17-alpine 作为基础镜像。
# 然后，我们将当前目录下的 main.go 文件复制到容器的 /app 目录中，并运行 go build 命令来构建应用程序。
# 最后，我们使用 EXPOSE 命令暴露应用程序的端口，并使用 CMD 命令来运行应用程序。

# 使用该 Dockerfile 构建镜像后，可以使用 docker run 命令来启动容器并运行 Go 语言程序。
# 例如，以下命令将构建名为 myapp 的 Docker 镜像，并启动一个容器来运行该应用程序：
# docker build -t myapp .
# docker run -p 8080:8080 myapp
# 其中 -p 8080:8080 参数用于将容器的 8080 端口映射到主机的 8080 端口，以便可以通过浏览器访问该应用程序。