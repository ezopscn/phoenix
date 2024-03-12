# 构建环境
FROM node:lts as build-env
# 设置工作目录
WORKDIR /app
# 将项目复制到工作目录，排除某些目录
COPY . ./
# 打包
RUN npm config set registry https://registry.npmmirror.com &&  \
    npm install &&  \
    npm run build

# 基础镜像
FROM nginx:alpine-perl
# 复制构建产物到 Nginx 的默认 Web 根目录
COPY --from=build-env /app/dist /usr/share/nginx/html
# 设置环境变量
ENV RUN_ENV=prod WEB_PATH=/usr/share/nginx/html
# 根据环境变量
RUN cd ${WEB_PATH} && \
    rm -f config.js
# 使用 80 端口运行
EXPOSE 80
# 启动服务
CMD ["cp", "${WEB_PATH}/${RUN_ENV}.js", "${WEB_PATH}/config.js", "&&", "nginx", "-g", "daemon off;"]
