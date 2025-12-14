# docker编译镜像添加proxy

如果你有可用的 HTTP/HTTPS 代理（如 Clash、V2Ray、Surge），可以设置以下环境变量

```shell
# Linux / macOS
export DOCKER_BUILDKIT=1
export HTTPS_PROXY=http://127.0.0.1:7897
export HTTP_PROXY=http://127.0.0.1:7897

docker build -t myapp .
```

