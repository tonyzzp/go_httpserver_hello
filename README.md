# go_httpserver_hello

一个空的http服务

https://hub.docker.com/r/tonyzzp/go_httpserver_hello


# 使用

## 直接使用命令

```bash
docker run -itd \
  -p 80:80 \
  tonyzzp/go_httpserver_hello
```

#### 使用https
运行时增加环境变量

- HTTP_SSL=true
- HTTP_HOST=myhost.com,www.myhost.com
- HTTP_EMAIL=mymail@mail.com

如:
```
docker run -itd \
    -p 443:443 \
    -p 80:80 \
    -e HTTP_SSL=true \
    -e HTTP_HOST=myhost.com,www.myhost.com \
    -e HTTP_EMAIL=mymail@mail.com \
    tonyzzp/go_httpserver_hello
```

## 使用docker-compose
```
docker-compose -f docker-compose-sample.yml up -d
```

