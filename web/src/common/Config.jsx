const backendUrl = window.CONFIG.backendUrl; // 后端接口前缀

// 接口信息
const APIConfig = {
  RunEnv: window.CONFIG.env,
  LoginAPI: backendUrl + "/login",
};

export { APIConfig };
