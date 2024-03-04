import { RouteData } from "./RouteData.jsx";
import {useRoutes} from "react-router";

// 生成路由
const RouteRules = () => {
  return useRoutes(RouteData);
};

export default RouteRules;
