import React from "react";
import ButterflyLayout from "../components/layout/Layout.jsx";
import RouteLazyLoad from "./RouteLazyLoad.jsx";
import { Navigate } from "react-router";

// 路由数据
export const RouteData = [
  {
    path: "/",
    element: <Navigate to="/dashboard" />,
  },
  {
    path: "/",
    element: <ButterflyLayout />,
    children: [
      {
        path: "dashboard",
        element: RouteLazyLoad(
          React.lazy(() => import("../pages/dashboard/Dashboard.jsx")),
        ),
      },
      {
        path: "node",
        element: RouteLazyLoad(
          React.lazy(() => import("../pages/kubernetes/node/Node.jsx")),
        ),
      },
      {
        path: "users",
        children: [
          {
            path: "list",
            element: RouteLazyLoad(
              React.lazy(() => import("../pages/user/UserList.jsx")),
            ),
          },
        ],
      },
    ],
  },
  // 无需登录的地址
  {
    path: "login",
    element: RouteLazyLoad(
      React.lazy(() => import("../pages/login/LoginBlack.jsx")),
    ),
    notNeedAuth: true,
  },
  {
    path: "403",
    element: RouteLazyLoad(React.lazy(() => import("../pages/error/403.jsx"))),
    notNeedAuth: true,
  },
  {
    path: "404",
    element: RouteLazyLoad(React.lazy(() => import("../pages/error/404.jsx"))),
    notNeedAuth: true,
  },
  {
    path: "500",
    element: RouteLazyLoad(React.lazy(() => import("../pages/error/500.jsx"))),
    notNeedAuth: true,
  },
  {
    path: "*",
    element: <Navigate to="/404" />,
    notNeedAuth: true,
  },
];
