# 基于golang镜像，指定1.17版本
FROM golang:1.17

# 指定工作目录
WORKDIR /go/src/gin_project

# 复制文件
COPY . .

# 设置环境变量(解决拉依赖的问题)
ENV GOPROXY="https://goproxy.cn,direct"

# 编译项目
RUN go build

# 暴露端口
EXPOSE 7070

# 执行
ENTRYPOINT ./gin_project
