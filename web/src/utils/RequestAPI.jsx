import { GET, POST } from "./Request.jsx";
import { APIConfig } from "../common/Config.jsx";

// 接口请求
export const LoginRequest = (data) => POST(APIConfig.LoginAPI, data); // 用户登录
export const LogoutRequest = () => GET(APIConfig.LogoutAPI); // 用户登出
export const UserCountRequest = () => GET(APIConfig.UserCountAPI); // 用户总数
export const CurrentUserInfoRequest = () => GET(APIConfig.CurrentUserInfoAPI); // 当前用户信息
export const CurrentUserDepartmentInfoRequest = () =>
  GET(APIConfig.CurrentUserDepartmentInfoAPI); // 当前用户信息
