<!--suppress HtmlDeprecatedAttribute -->
<h1 align="center">Hi 🥳, 前端开发文档</h1>

<p align="center">
  <a>
    <img src="https://img.shields.io/badge/-Ant Design-blue?style=flat-square&logo=antdesign&logoColor=white" alt="">
  </a>
</p>

<hr>

### 🤔 技术栈

开发依赖安装包：

```bash
# 自动格式化代码
npm i --save-dev --save-exact prettier
npm i --save-dev @types/node 

# 字体
npm i misans

# UI 和图标
npm i antd@4.24.14
npm i @ant-design/icons
npm i moment

# 路由
npm i react-router
npm i react-router-dom

# 样式
npm i less less-loader

# 网络请求
npm i axios

# 数据
npm i valtio
```

<br>

### ⚡ 特别说明

关于镜像的设计理念：

> 为了适配 docker 一次构建到处运行的理念。针对前端项目，我们选择将配置文件单独成一个 js 文件并放到 public 目录中。该目录有个特点，下面的文件文件在 build 之后都会原封不动的被拷贝到项目的根目录下。我们再从 index.js 中引入特定的 js 文件。

如何做到一个镜像运行不同的环境？

> 在 build 的时候，不同环境的配置都被打包到了项目的根目录，但实际我们只是引入了 config.js 文件。此时就意味着我们需要 docker 启动命令中，根据传入的 env 环境变量，将默认的 config.js 替换成真实环境的 js 文件即可。这样做的好处在于，相较于传统的暴力替换，这样的替换更稳定靠谱。