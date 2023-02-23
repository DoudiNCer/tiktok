FROM alpine:latest
MAINTAINER doudincer <doudi@outlook.lv>

WORKDIR /apps

# 将代码复制到容器中
COPY tiktok ./

EXPOSE 8086

# 启动容器时运行的命令
CMD ["/app/tiktok"]
