#!/bin/bash

APP_ID=$(cat APP_ID)
sysOS=$(uname -s)

if [ $sysOS = "Darwin" ]; then
    cpuBrand=$(echo $(sysctl -n machdep.cpu.brand_string) | cut -c -5)
    if [ ""$cpuBrand = "Apple" ]; then
		docker build --platform linux/amd64 --build-arg APP_ID_ARG="${APP_ID}" --build-arg BUILD_ENV=product-local -t "${APP_ID}":"local" .
	else
		docker build --build-arg APP_ID_ARG="${APP_ID}" --build-arg BUILD_ENV=product-local  -t "${APP_ID}":"local" .
	fi
else
	docker build --build-arg APP_ID_ARG="${APP_ID}" --build-arg BUILD_ENV=product-local  -t "${APP_ID}":"local" .
fi