# 告诉makefile 不需要去当前目录找这些同名的文件 只要makefile 文件中 找即可
.PHONY: all build run gotool clean help

# 定义一个变量 项目的名称
BINARY="bbs_server"

# 定义2个目标
all: gotool build

# 我自己的开发机器是mac 所以最终要在linux上部署 需要build出来 linux下的可执行文件
build:
	CGO_ENABled=0 GOOS=linux GOARCH=amd64 go build -o ${BINARY}
# 不加@ 则在make的时候 会打印出来对应的命令
run:
	@go run ./main.go config.yaml

# 格式化代码 并检查
gotool:
	go fmt ./
	go vet ./

# 删除对应的文件
clean:
	@if [ -f ${BINARY} ]; then rm ${BINARY} ; fi

help:
	@echo "make - 格式化 Go代码 并编译成二进制文件"
	@echo "make build - 编译 Go代码 并编译成linux平台下的二进制文件"
	@echo "make run - 直接运行"
	@echo "make clean 移除二进制文件"
	@echo "make gotool 运行go工具 fmt 和vet "


