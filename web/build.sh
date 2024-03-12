#!/bin/bash
###############################################################
# 说明：编译打包 Butterfly 项目
###############################################################
# 1.生成版本信息
git rev-parse HEAD > public/version
# 版本信息
version=$(cat public/version)
# 镜像信息
image="phoenix-web:${version}"

# 2.镜像构建
docker build -t ${image} .

# 3.上传镜像
# docker push ${image}