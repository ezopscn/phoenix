package data

import (
	"errors"
	"gorm.io/gorm"
	"phoenix/common"
	"phoenix/model"
)

// 菜单数据
var menus = []model.Menu{
	{
		BaseModel: model.BaseModel{Id: 1000},
		Name:      "工作空间",
		Icon:      "HomeOutlined",
		Path:      "/dashboard",
		Sort:      0,
		ParentId:  0,
		Roles: []model.Role{
			roles[1],
		},
	},
	{
		BaseModel: model.BaseModel{Id: 1010},
		Name:      "集群面板",
		Icon:      "ClusterOutlined",
		Path:      "/cluster",
		Sort:      0,
		ParentId:  0,
	},
	{
		BaseModel: model.BaseModel{Id: 1020},
		Name:      "节点管理",
		Icon:      "NodeIndexOutlined",
		Path:      "/node",
		Sort:      0,
		ParentId:  0,
	},
	{
		BaseModel: model.BaseModel{Id: 1030},
		Name:      "名称空间",
		Icon:      "AppstoreAddOutlined",
		Path:      "/namespace",
		Sort:      0,
		ParentId:  0,
	},
	{
		BaseModel: model.BaseModel{Id: 1040},
		Name:      "工作负载",
		Icon:      "DeploymentUnitOutlined",
		Path:      "/workload",
		Sort:      0,
		ParentId:  0,
		Children: []model.Menu{
			{
				BaseModel: model.BaseModel{Id: 1041},
				Name:      "工作单元（Pod）",
				Icon:      "",
				Path:      "/workload/pod",
				Sort:      0,
				ParentId:  1040,
			},
			{
				BaseModel: model.BaseModel{Id: 1042},
				Name:      "部署副本（Deployment）",
				Icon:      "",
				Path:      "/workload/deployment",
				Sort:      0,
				ParentId:  1040,
			},
			{
				BaseModel: model.BaseModel{Id: 1043},
				Name:      "有状态集（StatefulSet）",
				Icon:      "",
				Path:      "/workload/statefulset",
				Sort:      0,
				ParentId:  1040,
			},
			{
				BaseModel: model.BaseModel{Id: 1044},
				Name:      "守护进程（DaemonSet）",
				Icon:      "",
				Path:      "/workload/daemonset",
				Sort:      0,
				ParentId:  1040,
			},
			{
				BaseModel: model.BaseModel{Id: 1045},
				Name:      "普通任务（Job）",
				Icon:      "",
				Path:      "/workload/job",
				Sort:      0,
				ParentId:  1040,
			},
			{
				BaseModel: model.BaseModel{Id: 1046},
				Name:      "定时任务（CronJob）",
				Icon:      "",
				Path:      "/workload/cronjob",
				Sort:      0,
				ParentId:  1040,
			},
		},
	},
	{
		BaseModel: model.BaseModel{Id: 1050},
		Name:      "服务发现",
		Icon:      "PartitionOutlined",
		Path:      "/discovery",
		Sort:      0,
		ParentId:  0,
		Children: []model.Menu{
			{
				BaseModel: model.BaseModel{Id: 1051},
				Name:      "服务发现（Service）",
				Icon:      "",
				Path:      "/discovery/service",
				Sort:      0,
				ParentId:  1050,
			},
			{
				BaseModel: model.BaseModel{Id: 1052},
				Name:      "负载均衡（Ingress）",
				Icon:      "",
				Path:      "/discovery/ingress",
				Sort:      0,
				ParentId:  1050,
			},
		},
	},
	{
		BaseModel: model.BaseModel{Id: 1060},
		Name:      "存储管理",
		Icon:      "DeliveredProcedureOutlined",
		Path:      "/storage",
		Sort:      0,
		ParentId:  0,
		Children: []model.Menu{
			{
				BaseModel: model.BaseModel{Id: 1061},
				Name:      "存储类别（StorageClass）",
				Icon:      "",
				Path:      "/storage/class",
				Sort:      0,
				ParentId:  1060,
			},
			{
				BaseModel: model.BaseModel{Id: 1062},
				Name:      "持久化卷（PV）",
				Icon:      "",
				Path:      "/storage/pv",
				Sort:      0,
				ParentId:  1060,
			},
			{
				BaseModel: model.BaseModel{Id: 1063},
				Name:      "持久声明（PVC）",
				Icon:      "",
				Path:      "/storage/pvc",
				Sort:      0,
				ParentId:  1060,
			},
		},
	},
	{
		BaseModel: model.BaseModel{Id: 1070},
		Name:      "配置密文",
		Icon:      "SnippetsOutlined",
		Path:      "/config",
		Sort:      0,
		ParentId:  0,
		Children: []model.Menu{
			{
				BaseModel: model.BaseModel{Id: 1071},
				Name:      "配置管理（ConfigMap）",
				Icon:      "",
				Path:      "/config/configmap",
				Sort:      0,
				ParentId:  1070,
			},
			{
				BaseModel: model.BaseModel{Id: 1072},
				Name:      "密文管理（Secret）",
				Icon:      "",
				Path:      "/config/secret",
				Sort:      0,
				ParentId:  1070,
			},
		},
	},
	{
		BaseModel: model.BaseModel{Id: 1090},
		Name:      "用户中心",
		Icon:      "TeamOutlined",
		Path:      "/users",
		Sort:      0,
		ParentId:  0,
		Children: []model.Menu{
			{
				BaseModel: model.BaseModel{Id: 1091},
				Name:      "用户管理",
				Icon:      "",
				Path:      "/users/list",
				Sort:      0,
				ParentId:  1090,
			},
			{
				BaseModel: model.BaseModel{Id: 1092},
				Name:      "分组管理",
				Icon:      "",
				Path:      "/users/group",
				Sort:      0,
				ParentId:  1090,
			},
			{
				BaseModel: model.BaseModel{Id: 1093},
				Name:      "角色管理",
				Icon:      "",
				Path:      "/users/role",
				Sort:      0,
				ParentId:  1090,
			},
		},
	},
	{
		BaseModel: model.BaseModel{Id: 1100},
		Name:      "系统配置",
		Icon:      "SettingOutlined",
		Path:      "/system",
		Sort:      0,
		ParentId:  0,
		Children: []model.Menu{
			{
				BaseModel: model.BaseModel{Id: 1101},
				Name:      "部门管理",
				Icon:      "",
				Path:      "/system/department",
				Sort:      0,
				ParentId:  1100,
			},
			{
				BaseModel: model.BaseModel{Id: 1102},
				Name:      "菜单管理",
				Icon:      "",
				Path:      "/system/menu",
				Sort:      0,
				ParentId:  1100,
			},
			{
				BaseModel: model.BaseModel{Id: 1103},
				Name:      "接口管理",
				Icon:      "",
				Path:      "/system/api",
				Sort:      0,
				ParentId:  1100,
			},
			{
				BaseModel: model.BaseModel{Id: 1104},
				Name:      "服务配置",
				Icon:      "",
				Path:      "/system/setting",
				Sort:      0,
				ParentId:  1100,
			},
		},
	},
	{
		BaseModel: model.BaseModel{Id: 1110},
		Name:      "日志审计",
		Icon:      "InsuranceOutlined",
		Path:      "/log",
		Sort:      0,
		ParentId:  0,
		Roles: []model.Role{
			roles[1],
		},
		Children: []model.Menu{
			{
				BaseModel: model.BaseModel{Id: 1111},
				Name:      "操作日志",
				Icon:      "",
				Path:      "/log/operation",
				Sort:      0,
				ParentId:  1110,
				Roles: []model.Role{
					roles[1],
				},
			},
			{
				BaseModel: model.BaseModel{Id: 1112},
				Name:      "登录日志",
				Icon:      "",
				Path:      "/log/login",
				Sort:      0,
				ParentId:  1110,
				Roles: []model.Role{
					roles[1],
				},
			},
			{
				BaseModel: model.BaseModel{Id: 1113},
				Name:      "改密日志",
				Icon:      "",
				Path:      "/log/password",
				Sort:      0,
				ParentId:  1110,
				Roles: []model.Role{
					roles[1],
				},
			},
		},
	},
	{
		BaseModel: model.BaseModel{Id: 1120},
		Name:      "个人中心",
		Icon:      "UserOutlined",
		Path:      "/me",
		Sort:      0,
		ParentId:  0,
		Roles: []model.Role{
			roles[1],
		},
	},
	{
		BaseModel: model.BaseModel{Id: 1130},
		Name:      "获取帮助",
		Icon:      "FileProtectOutlined",
		Path:      "/help",
		Sort:      0,
		ParentId:  0,
		Roles: []model.Role{
			roles[1],
		},
	},
}

// 递归插入数据方法
func insertMenusData(menus []model.Menu) {
	var menu model.Menu
	for _, item := range menus {
		// 查看数据是否存在，如果不存在才执行创建
		err := common.DB.Where("id = ? or name = ? or path = ?", item.Id, item.Name, item.Path).First(&menu).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			common.DB.Create(&item)
		}
		// 递归插入子菜单
		if len(item.Children) > 0 {
			insertMenusData(item.Children)
		}
	}
}

// 菜单数据初始化
func MenuInit() {
	insertMenusData(menus)
}
