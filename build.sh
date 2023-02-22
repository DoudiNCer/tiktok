#!/usr/bin/env sh
# 静态编译以便在 alpine 上运行
go build -ldflags '-s -w -L /usr/lib -linkmode "external" -extldflags "-static"'

# 构建 Docker 镜像
docker build -t doudincer/tiktok:1 .