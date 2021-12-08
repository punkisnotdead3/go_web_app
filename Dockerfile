# 代表这个镜像只参与编译过程
FROM golang:alpine AS builder
# 设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# 移动到工作目录 workdir 和cd 这个命令有点类似
WORKDIR /build

# 复制项目中的 go.mod 和 go.sum 文件并下载依赖信息 之前那个简单的项目是没有第三方依赖的 所以可以忽略这个步骤
COPY go.mod .
COPY go.sum .

RUN go env -w GOPROXY=https://goproxy.cn,direct

RUN go mod download

# 将代码复制到容器中
COPY . .

#将我们的代码编译成二进制可执行文件bbs
RUN go build -o bbs .

# 创建一个小的镜像 linux 最终跑起来的就是这个最小的镜像
FROM scratch

# 拷贝静态文件
COPY config.yaml .
COPY ./html /html


# 从builder镜像中 把build bbs 拷贝到当前目录
COPY --from=builder /build/bbs /

# 需要运行的命令
ENTRYPOINT ["/bbs","config.yaml"]

# 声明服务端口8812
EXPOSE 8812

# 启动容器时运行的命令
CMD ["/dist/app"]



