## Docker Build

```sh
#!/usr/bin/env bash

# 使用 DOCKER_NAME=product-local DOCKER_TAG=0.0.1 sh build_docker.sh

# 使用 DOCKER_NAME=shepherd_svc_main DOCKER_TAG=local APP_ID=shepherd_svc_main sh build_docker.sh
# 使用 DOCKER_NAME=shepherd_svc_alg DOCKER_TAG=local APP_ID=shepherd_svc_alg BUILD_ENV=alg-product-local sh build_docker.sh
# 使用 DOCKER_NAME=shepherd_svc_daemon DOCKER_TAG=local APP_ID=shepherd_svc_daemon MAIN=cmd/shepherd/main.go sh build_docker.sh

# 拉取 gopkg
# git clone git@git.lianjia.com:vrlab_server/gopkg.git

# 构建docker镜像
APP_ID=$(cat APP_ID)
sysOS=$(uname -s)


if [ $sysOS = "Darwin" ]; then
    cpuBrand=$(echo $(sysctl -n machdep.cpu.brand_string) | cut -c -5)
    if [ ""$cpuBrand = "Apple" ]; then
    echo '1'
		#docker build --platform linux/amd64 --build-arg APP_ID_ARG="${APP_ID}" --build-arg BUILD_ENV="${BUILD_ENV}" -t "${DOCKER_NAME}":"${DOCKER_TAG}" .
		docker build --platform linux/amd64 --build-arg APP_ID_ARG="${APP_ID}" --build-arg BUILD_ENV=product-local -t "${APP_ID}":"local" .
	else
	  echo '2'
		#docker build --build-arg APP_ID_ARG="${APP_ID}" --build-arg BUILD_ENV="${BUILD_ENV}" -t "${DOCKER_NAME}":"${DOCKER_TAG}" .
		docker build --build-arg APP_ID_ARG="${APP_ID}" --build-arg BUILD_ENV=product-local  -t "${APP_ID}":"local" .
	fi
else
  echo '2'
	#docker build --build-arg APP_ID_ARG="${APP_ID}" --build-arg BUILD_ENV="${BUILD_ENV}" -t "${DOCKER_NAME}":"${DOCKER_TAG}" .
	docker build --build-arg APP_ID_ARG="${APP_ID}" --build-arg BUILD_ENV=product-local  -t "${APP_ID}":"local" .
fi

```

## 一份Dockerfile

```docker
# 使用go
# FROM business_base:local

# ------------------------- go mod 基础库 ---------------------------------
FROM node:16.17-alpine

MAINTAINER realsee.com

WORKDIR /data0/www/htdocs

RUN apk update  && apk add --no-cache git tar

# Internal Gitlab
# ------------------------- 编译文件 ---------------------------------

ARG APP_ID_ARG
ARG BUILD_ENV
ENV APP_ID=${APP_ID_ARG}

ADD ./ /usr/local/${APP_ID}

# 指定工作路径
WORKDIR /usr/local/${APP_ID}
RUN APP_ID=${APP_ID} BUILD_ENV=${BUILD_ENV} sh -e build_local.sh

# ------------------------- 输出最终镜像 ---------------------------------

FROM node:16-alpine

ARG APP_ID_ARG
ENV APP_ID=${APP_ID_ARG}

RUN apk update && apk add zip tzdata curl

WORKDIR /data0/www/htdocs/${APP_ID}/bin

COPY --from=0 /usr/local/${APP_ID}/releases/${APP_ID}/bin /data0/www/htdocs/${APP_ID}/bin

# 容器对外爆露的端口
EXPOSE 80

CMD sh /data0/www/htdocs/${APP_ID}/bin/run.sh run

```