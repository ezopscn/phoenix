import { proxy } from "valtio"; // 地区数据

// 地区
export const RegionStates = proxy({
  Provinces: [], //省份数据
});

// 部门
export const DepartmentStates = proxy({
  Departments: [], //部门数据
});

// 角色
export const RoleStates = proxy({
  Roles: [], //角色数据
});

// 用户
export const UserStates = proxy({
  UserSearchFieldExpand: false, //是否展开所有搜索字段
  UserListSearchParams: {}, // 用户搜索参数
});
