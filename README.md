# gin-blog
## Docker 部署

### Dockerfile

```dockerfile
FROM golang:alpine

# 设置Go环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn,direct

# 移动工作目录
WORKDIR /build

# 将代码复制到容器中
COPY . .

# 将我们的代码编译成二进制可执行文件app
RUN go build -o app .

# 移动到用于存放生成的二进制文件的/dist目录
WORKDIR /dist

# 将二进制文件从/build 目录复制到 /dist
RUN cp /build/app . \
    && cp -r /build/conf . \
    && cp -r /build/runtime . 

# 声明服务端口
EXPOSE 8000
CMD ["/dist/app"]
```

### 流程

1. git clone <project url>

2. 修改conf/app.ini中database和redis的Host

3. 构建并运行镜像

```shell
$ docker build  -t gin-blog .
$ docker run -p 8000:8000 --name gin-blog -d gin-blog
```

