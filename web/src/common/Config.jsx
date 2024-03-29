const backendUrl = window.CONFIG.backendUrl; // 后端接口前缀

// 接口信息
const APIConfig = {
  RunEnv: window.CONFIG.env, // 运行环境
  BaseURL: backendUrl, // 基础后端连接
  LoginAPI: backendUrl + '/login', // 登录接口
  LogoutAPI: backendUrl + '/logout', // 登出接口
  UserCountAPI: backendUrl + '/user/count', // 用户总数接口
  CurrentUserInfoAPI: backendUrl + '/user/info', // 当前用户信息接口
  CurrentUserDepartmentInfoAPI: backendUrl + '/department/info', // 当前用户部门信息接口
  CurrentUserMenuTreeAPI: backendUrl + '/menu/tree', // 当前用户的菜单树
  CurrentUserMenuListAPI: backendUrl + '/user/menu/list', // 当前用户的菜单列表
  UserListAPI: backendUrl + '/user/list', // 用户列表
};

export { APIConfig };
