# tiktok
第五届青训营后端大项目——极简版抖音

## 运行依赖
- Docker
- Docker Compose
## 如何启动
0. 在当前文件夹创建`sqldb`文件夹或修改`/var/lib/mysql`的挂载路径：
```shell
mkdir sqldb
```
1. 启动项目：
```shell
docker-compose up -d
```
2. 暂停项目使用`docker-compose stop`，再次启动使用`docker-compose start`
3. 运行项目（暂时手动运行，端口8086）
