import { useRoutes } from "react-router";
import { RouteData } from "./RouteData.jsx";

// 生成路由
const RouteRules = () => {
  return useRoutes(RouteData);
};

export default RouteRules;
