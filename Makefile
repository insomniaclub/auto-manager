AM = "auto-manager"

all: check build run

check:
	go mod tidy

build:
	go build -o ${AM}

run:
	go run main.go

help:
	@echo "make check - 检查第三方依赖并安装"
	@echo "make build - 编译Go代码，生成可执行文件"
	@echo "make run   - 直接运行main.go"
