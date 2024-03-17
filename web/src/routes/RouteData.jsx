import React from 'react';
import RouteLazyLoad from './RouteLazyLoad.jsx';
import { Navigate } from 'react-router';
import AdminLayout from '../components/layout/Layout.jsx';
import ErrorLayout from '../components/error/ErrorLayout.jsx';

// 路由数据
export const RouteData = [
  {
    path: '/',
    element: <Navigate to="/dashboard" />,
  },
  {
    path: '/',
    element: <AdminLayout />,
    children: [
      {
        path: 'dashboard',
        element: RouteLazyLoad(React.lazy(() => import('../pages/dashboard/Dashboard.jsx'))),
      },
      {
        path: 'cluster',
        element: RouteLazyLoad(React.lazy(() => import('../pages/kubernetes/cluster/ClusterDashboard.jsx'))),
      },
      {
        path: 'node',
        element: RouteLazyLoad(React.lazy(() => import('../pages/kubernetes/node/Node.jsx'))),
      },
      {
        path: 'namespace',
        element: RouteLazyLoad(React.lazy(() => import('../pages/kubernetes/namespace/Namespace.jsx'))),
      },
      {
        path: 'workload',
        children: [
          {
            path: 'pod',
            element: RouteLazyLoad(React.lazy(() => import('../pages/kubernetes/workload/pod/Pod.jsx'))),
          },
          {
            path: 'deployment',
            element: RouteLazyLoad(React.lazy(() => import('../pages/kubernetes/workload/deployment/Deployment.jsx'))),
          },
          {
            path: 'statefulset',
            element: RouteLazyLoad(React.lazy(() => import('../pages/kubernetes/workload/statefulset/StatefulSet.jsx'))),
          },
          {
            path: 'daemonset',
            element: RouteLazyLoad(React.lazy(() => import('../pages/kubernetes/workload/daemonset/DaemonSet.jsx'))),
          },
          {
            path: 'job',
            element: RouteLazyLoad(React.lazy(() => import('../pages/kubernetes/workload/job/Job.jsx'))),
          },
          {
            path: 'cronjob',
            element: RouteLazyLoad(React.lazy(() => import('../pages/kubernetes/workload/cronjob/CronJon.jsx'))),
          },
        ],
      },
      {
        path: 'discovery',
        children: [
          {
            path: 'service',
            element: RouteLazyLoad(React.lazy(() => import('../pages/kubernetes/discovery/service/Service.jsx'))),
          },
          {
            path: 'ingress',
            element: RouteLazyLoad(React.lazy(() => import('../pages/kubernetes/discovery/ingress/Ingress.jsx'))),
          },
        ],
      },
      {
        path: 'storage',
        children: [
          {
            path: 'class',
            element: RouteLazyLoad(React.lazy(() => import('../pages/kubernetes/storage/class/Class.jsx'))),
          },
          {
            path: 'pv',
            element: RouteLazyLoad(React.lazy(() => import('../pages/kubernetes/storage/pv/PV.jsx'))),
          },
          {
            path: 'pvc',
            element: RouteLazyLoad(React.lazy(() => import('../pages/kubernetes/storage/pvc/PVC.jsx'))),
          },
        ],
      },
      {
        path: 'config',
        children: [
          {
            path: 'configmap',
            element: RouteLazyLoad(React.lazy(() => import('../pages/kubernetes/config/configmap/ConfigMap.jsx'))),
          },
          {
            path: 'secret',
            element: RouteLazyLoad(React.lazy(() => import('../pages/kubernetes/config/secret/Secret.jsx'))),
          },
        ],
      },
      {
        path: 'users',
        children: [
          {
            path: 'list',
            element: RouteLazyLoad(React.lazy(() => import('../pages/users/user/UserList.jsx'))),
          },
          {
            path: 'group',
            element: RouteLazyLoad(React.lazy(() => import('../pages/users/group/Group.jsx'))),
          },
          {
            path: 'role',
            element: RouteLazyLoad(React.lazy(() => import('../pages/users/role/Role.jsx'))),
          },
        ],
      },
      {
        path: 'system',
        children: [
          {
            path: 'department',
            element: RouteLazyLoad(React.lazy(() => import('../pages/system/department/Department.jsx'))),
          },
          {
            path: 'menu',
            element: RouteLazyLoad(React.lazy(() => import('../pages/system/menu/Menu.jsx'))),
          },
          {
            path: 'api',
            element: RouteLazyLoad(React.lazy(() => import('../pages/system/api/API.jsx'))),
          },
          {
            path: 'setting',
            element: RouteLazyLoad(React.lazy(() => import('../pages/system/setting/Setting.jsx'))),
          },
        ],
      },
      {
        path: 'log',
        children: [
          {
            path: 'operation',
            element: RouteLazyLoad(React.lazy(() => import('../pages/log/operation/Operation.jsx'))),
          },
          {
            path: 'login',
            element: RouteLazyLoad(React.lazy(() => import('../pages/log/login/Login.jsx'))),
          },
          {
            path: 'password',
            element: RouteLazyLoad(React.lazy(() => import('../pages/log/password/Password.jsx'))),
          },
        ],
      },
      {
        path: 'me',
        element: RouteLazyLoad(React.lazy(() => import('../pages/me/UserCenter.jsx'))),
      },
      {
        path: 'help',
        element: RouteLazyLoad(React.lazy(() => import('../pages/help/Help.jsx'))),
      },
    ],
  },
  // 无需登录的地址
  {
    path: 'login',
    element: RouteLazyLoad(React.lazy(() => import('../pages/login/Login.jsx'))),
    notNeedAuth: true,
  },
  {
    path: 'error',
    element: <ErrorLayout />,
    children: [
      {
        path: '403',
        element: RouteLazyLoad(React.lazy(() => import('../pages/error/403.jsx'))),
        notNeedAuth: true,
      },
      {
        path: '404',
        element: RouteLazyLoad(React.lazy(() => import('../pages/error/404.jsx'))),
        notNeedAuth: true,
      },
      {
        path: '500',
        element: RouteLazyLoad(React.lazy(() => import('../pages/error/500.jsx'))),
        notNeedAuth: true,
      },
    ],
  },
  {
    path: '*',
    element: <Navigate to="/error/404" />,
    notNeedAuth: true,
  },
];
