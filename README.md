# **go-gin-scaffold**
用于go开发restful api和网页的脚手架  
使用gin提供http服务  
使用beego验证器  
使用gorose的orm  
使用logrus作为日志管理

# **Docker镜像**
1. 同时修改config.ini和Dockerfile里的端口号为自己喜欢的数字,默认8889
2. 执行 `docker build -t go:api .` 构建镜像
3. 执行 `docker run -itd -p 8889:8889 --name goapi go:api /bin/sh`
4. 如果Mysql/Redis用的镜像,执行 `docker run -itd -p 服务器端口:镜像端口(修改Dockerfile里的EXPOSE) --link=mysql镜像名:别名 --link=redis镜像名:别名 --name goapi go:api /bin/sh` ,并将config.ini中mysql/redis的ip地址改成对应的别名
5. http://127.0.0.1:8889 访问
