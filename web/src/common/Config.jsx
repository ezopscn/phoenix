const backendUrl = window.CONFIG.backendUrl; // 后端接口前缀

// 接口信息
const APIConfig = {
  RunEnv: window.CONFIG.env, // 运行环境
  BaseURL: backendUrl, // 基础后端连接
  LoginAPI: backendUrl + "/login", // 登录接口
  LogoutAPI: backendUrl + "/logout", // 登出接口
  CurrentUserInfoAPI: backendUrl + "/user/info", // 当前用户信息接口
};

export { APIConfig };
