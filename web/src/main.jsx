import React from "react";
import ReactDOM from "react-dom/client";
// 由于 antd 组件的默认文案是英文，所以需要修改为中文
import { ConfigProvider } from "antd";
// 中文支持
import zhCN from "antd/es/locale/zh_CN";
// Antd 样式
import "antd/dist/antd.less";
// 字体
import "misans/lib/Normal/MiSans-Regular.min.css";
// 修改默认样式
import "./assets/css/antd-rewrite.less";
import "./assets/css/admin.less";
import App from "./App.jsx";

ReactDOM.createRoot(document.getElementById("root")).render(
  // autoInsertSpaceInButton：解决按钮的文本为两个汉字时中间自动补充空格的问题
  <ConfigProvider locale={zhCN} autoInsertSpaceInButton={false}>
    <App />
  </ConfigProvider>,
);
