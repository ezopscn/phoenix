import React, { useEffect, useState } from "react";
import { Outlet, useNavigate } from "react-router";
import { Avatar, Cascader, Dropdown, Layout, Menu, message } from "antd";
import { Logo, LogoWithTitle } from "../../common/Resource.jsx";
import { LayoutMenuData } from "./LayoutData.jsx";
import {
  CurrentUserInfoRequest,
  LogoutRequest,
} from "../../utils/RequestAPI.jsx";
import { UserStates } from "../../store/Store.jsx";
import { useSnapshot } from "valtio";
import {
  GetLocalStorageItem,
  SetLocalStorageItem,
} from "../../utils/Storage.jsx";
import { MoreOutlined } from "@ant-design/icons";

const { Header, Sider, Content, Footer } = Layout;

const AdminLayout = () => {
  // 菜单宽度
  const menuWidth = 240;
  const menuCollapsedWidth = 60;

  // 菜单跳转
  const navigate = useNavigate();
  const [collapsed, setCollapsed] = useState(false);

  // 获取用户信息，为了减少获取次数，会将数据保存到 storage 一段时间，默认取缓存中数据
  useEffect(() => {
    (async () => {
      try {
        let cacheKeyName = "user_info";
        let userInfo = GetLocalStorageItem(cacheKeyName);
        if (userInfo === null) {
          const res = await CurrentUserInfoRequest();
          if (res.code === 200) {
            userInfo = res.data.info;
            SetLocalStorageItem(cacheKeyName, userInfo, 60); // 有效期 60 秒
          } else {
            message.error(res.message);
          }
        }
        UserStates.CurrentUserInfo = userInfo;
      } catch (e) {
        console.log(e);
        message.error("服务器异常，请联系管理员");
      }
    })();
  }, []);

  // 级联筛选集群和名称空间
  const clustersAndNamespacesData = [
    {
      value: "test",
      label: "自建机房 | PDC TEST",
      children: [
        {
          value: "kube-system",
          label: "kube-system",
        },
        {
          value: "default",
          label: "default",
        },
      ],
    },
    {
      value: "dev",
      label: "自建机房 | MDC DEV",
      children: [
        {
          value: "kube-system",
          label: "kube-system",
        },
        {
          value: "default",
          label: "default",
        },
      ],
    },
    {
      value: "advance",
      label: "腾讯云广州 | TKE ADVANCE",
      children: [
        {
          value: "kube-system",
          label: "kube-system",
        },
        {
          value: "default",
          label: "default",
        },
      ],
    },
    {
      value: "prod",
      label: "腾讯云上海 | TKE PROD",
      children: [
        {
          value: "kube-system",
          label: "kube-system",
        },
        {
          value: "default",
          label: "default",
        },
      ],
    },
  ];

  const clustersAndNamespacesOnChange = (value, selectedOptions) => {
    console.log(value, selectedOptions);
  };

  // 筛选搜索
  const clustersAndNamespacesFilter = (inputValue, path) =>
    path.some(
      (option) =>
        option.label.toLowerCase().indexOf(inputValue.toLowerCase()) > -1,
    );

  // 用户登出方法
  const logoutHandler = async () => {
    try {
      const res = await LogoutRequest();
      if (res.code === 200) {
        localStorage.clear();
        message.success("用户注销成功");
        navigate("/login");
      } else {
        message.error(res.message);
      }
    } catch (e) {
      message.error("服务器异常，请联系管理员");
    }
  };

  // 用户数据
  const { CurrentUserInfo } = useSnapshot(UserStates);

  // 下拉菜单
  const LayoutDropdownMenuData = [
    {
      key: "1",
      label: (
        <a rel="noopener noreferrer" href="">
          @{CurrentUserInfo?.en_name}（{CurrentUserInfo?.cn_name}）
        </a>
      ),
      disabled: true,
    },
    {
      type: "divider",
    },
    {
      key: "2",
      label: (
        <a rel="noopener noreferrer" href="">
          联系我们
        </a>
      ),
    },
    {
      key: "3",
      label: (
        <a rel="noopener noreferrer" onClick={logoutHandler}>
          注销登录
        </a>
      ),
    },
  ];

  return (
    <Layout>
      <Sider
        className="admin-sider"
        width={menuWidth}
        collapsedWidth={menuCollapsedWidth}
        collapsible
        collapsed={collapsed}
        onCollapse={(value) => setCollapsed(value)}
      >
        <div
          className="admin-layout-logo"
          style={{
            width: collapsed ? menuCollapsedWidth + "px" : menuWidth + "px",
          }}
        >
          <img src={collapsed ? Logo : LogoWithTitle} alt="" />
        </div>
        <Menu
          className="admin-sider-menu"
          theme="dark"
          mode="inline"
          defaultSelectedKeys={["1"]}
          items={LayoutMenuData}
          onClick={({ key }) => {
            console.log(key);
            navigate(key); // 路由跳转
          }}
        />
      </Sider>
      <Layout
        style={{
          marginLeft: collapsed ? menuCollapsedWidth + "px" : menuWidth + "px",
          minHeight: "100vh",
          backgroundColor: "#ffffff",
        }}
      >
        <Header className="admin-header">
          <div className="admin-header-title">
            <Cascader
              style={{ width: "300px" }}
              options={clustersAndNamespacesData}
              onChange={clustersAndNamespacesOnChange}
              placeholder="切换集群和名称空间"
              showSearch={{
                clustersAndNamespacesFilter,
              }}
            />
          </div>
          <div className="admin-header-menu">
            <Dropdown menu={{ items: LayoutDropdownMenuData }}>
              <div className="admin-header-dropdown">
                <Avatar src={CurrentUserInfo?.avatar} size={28} />
                <MoreOutlined />
              </div>
            </Dropdown>
          </div>
        </Header>
        <Content>
          <Outlet />
        </Content>
        <Footer className="admin-footer">
          <b>🧬 PHOENIX </b>© 2024 EZOPS.CN, All Rights Reserved.
        </Footer>
      </Layout>
    </Layout>
  );
};

export default AdminLayout;
