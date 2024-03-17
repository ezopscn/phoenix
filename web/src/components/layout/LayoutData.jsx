import {
  AppstoreAddOutlined,
  ClusterOutlined,
  DeliveredProcedureOutlined,
  DeploymentUnitOutlined,
  FileProtectOutlined,
  HomeOutlined,
  InsuranceOutlined,
  NodeIndexOutlined,
  PartitionOutlined,
  SettingOutlined,
  SnippetsOutlined,
  TeamOutlined,
  UserOutlined,
} from "@ant-design/icons"; // 菜单数据

// 生成菜单结构
function getMenuItem(label, key, icon, children, type) {
  return {
    key,
    icon,
    children,
    label,
    type,
  };
}

// 菜单数据
export const LayoutMenuData = [
  getItem("工作空间", "/dashboard", <HomeOutlined />),
  getItem("集群面板", "/cluster", <ClusterOutlined />),
  getItem("节点管理", "/node", <NodeIndexOutlined />),
  getItem("名称空间", "/namespace", <AppstoreAddOutlined />),
  getItem("工作负载", "/workload", <DeploymentUnitOutlined />, [
    getItem("工作单元（Pod）", "/workload/pod"),
    getItem("部署副本（Deployment）", "/workload/deployment"),
    getItem("有状态集（StatefulSet）", "/workload/statefulset"),
    getItem("守护进程（DaemonSet）", "/workload/daemonset"),
    getItem("普通任务（Job）", "/workload/job"),
    getItem("定时任务（CronJob）", "/workload/cronjob"),
  ]),
  getItem("服务发现", "/discovery", <PartitionOutlined />, [
    getItem("服务发现（Service）", "/discovery/service"),
    getItem("负载均衡（Ingress）", "/discovery/ingress"),
  ]),
  getItem("存储管理", "/storage", <DeliveredProcedureOutlined />, [
    getItem("存储类别（StorageClass）", "/storage/class"),
    getItem("持久化卷（PV）", "/storage/pv"),
    getItem("持久声明（PVC）", "/storage/pvc"),
  ]),
  getItem("配置密文", "/config", <SnippetsOutlined />, [
    getItem("配置管理（ConfigMap）", "/config/configmap"),
    getItem("密文管理（Secret）", "/config/secret"),
  ]),
  getItem("用户中心", "/users", <TeamOutlined />, [
    getItem("用户列表", "/users/list"),
    getItem("用户分组", "/users/group"),
    getItem("用户角色", "/users/role"),
  ]),
  getItem("系统配置", "/system", <SettingOutlined />, [
    getItem("环境管理", "/system/environment"),
    getItem("菜单管理", "/system/menu"),
    getItem("服务配置", "/system/setting"),
  ]),
  getItem("日志审计", "/log", <InsuranceOutlined />, [
    getItem("登录日志", "/log/login"),
    getItem("改密日志", "/log/password"),
    getItem("操作日志", "/log/operation"),
  ]),
  getItem("个人中心", "/me", <UserOutlined />),
  getItem("获取帮助", "/help", <FileProtectOutlined />),
];
