import React, { useEffect, useState } from 'react';
import { Outlet, useLocation, useNavigate } from 'react-router';
import { Avatar, Cascader, Dropdown, Layout, Menu, message } from 'antd';
import { FooterText, Logo, LogoWithTitle } from '../../common/Resource.jsx';
import { CurrentUserInfoRequest, CurrentUserMenuListRequest, LogoutRequest } from '../../utils/RequestAPI.jsx';
import { LayoutStates, UserStates } from '../../store/Store.jsx';
import { useSnapshot } from 'valtio';
import { MoreOutlined } from '@ant-design/icons';
import { GenerateMenuTree } from '../../utils/Menu.jsx';

const { Header, Sider, Content, Footer } = Layout;

const AdminLayout = () => {
  // 菜单宽度
  const menuWidth = 240;
  const menuCollapsedWidth = 60;

  // 用于获取请求连接
  const { pathname } = useLocation();

  // 菜单跳转
  const navigate = useNavigate();
  const [collapsed, setCollapsed] = useState(false);
  const [menuList, setMenuList] = useState([]);
  const [menuTree, setMenuTree] = useState([]);
  const { MenuSiderCollapsed, MenuOpenKeys, MenuSelectKeys } = useSnapshot(LayoutStates);

  // 用户数据
  const { CurrentUserInfo } = useSnapshot(UserStates);

  // 获取用户信息
  useEffect(() => {
    (async () => {
      try {
        const res = await CurrentUserInfoRequest();
        if (res.code === 200) {
          UserStates.CurrentUserInfo = res.data.info;
        } else if (res.code === 1000) {
          message.error('用户认证失效，请重新登录');
          localStorage.clear();
          navigate('/login');
        } else {
          message.error(res.message);
        }
      } catch (e) {
        console.log(e);
        message.error('服务器异常，请联系管理员');
      }
    })();
  }, []);

  useEffect(() => {
    (async () => {
      try {
        // 查询菜单
        const res = await CurrentUserMenuListRequest();
        if (res.code === 200) {
          // 处理菜单树
          setMenuList(res.data.list);
          const tree = GenerateMenuTree(0, res.data.list);
          setMenuTree(tree);
        } else {
          message.error(res.message);
        }
      } catch (e) {
        console.log(e);
        message.error('服务器异常，请联系管理员');
      }
    })();
  }, []);

  const findKeyList = (path, menus) => {
    const result = [];
    let fmenu = {};

    // 先找到对应的菜单
    menus.forEach((menu) => {
      if (menu.path === path) {
        fmenu = menu;
        result.push(path);
      }
    });

    // 通过找到的菜单，查询它的所有上级菜单
    const findMenu = (menu) => {
      if (menu.parent_id !== 0) {
        menus.forEach((item) => {
          if (item.id === menu.parent_id) {
            result.push(item.path);
            findMenu(item);
          }
        });
      }
    };

    findMenu(fmenu);
    return result;
  };

  useEffect(() => {
    if (menuList.length > 0) {
      // 修改默认打开和选中菜单
      let keys = findKeyList(pathname, menuList);
      LayoutStates.MenuSelectKeys = keys;

      // 解决收起菜单会弹出子菜单的问题
      if (MenuSiderCollapsed) {
        LayoutStates.MenuOpenKeys = [];
      } else {
        LayoutStates.MenuOpenKeys = keys;
      }
    }
  }, [pathname, menuList]);

  // 级联筛选集群和名称空间
  const clustersAndNamespacesData = [
    {
      value: 'test',
      label: '自建机房 | PDC TEST',
      children: [
        {
          value: 'kube-system',
          label: 'kube-system',
        },
        {
          value: 'default',
          label: 'default',
        },
      ],
    },
    {
      value: 'dev',
      label: '自建机房 | MDC DEV',
      children: [
        {
          value: 'kube-system',
          label: 'kube-system',
        },
        {
          value: 'default',
          label: 'default',
        },
      ],
    },
    {
      value: 'advance',
      label: '腾讯云广州 | TKE ADVANCE',
      children: [
        {
          value: 'kube-system',
          label: 'kube-system',
        },
        {
          value: 'default',
          label: 'default',
        },
      ],
    },
    {
      value: 'prod',
      label: '腾讯云上海 | TKE PROD',
      children: [
        {
          value: 'kube-system',
          label: 'kube-system',
        },
        {
          value: 'default',
          label: 'default',
        },
      ],
    },
  ];

  const clustersAndNamespacesOnChange = (value, selectedOptions) => {
    console.log(value, selectedOptions);
  };

  // 筛选搜索
  const clustersAndNamespacesFilter = (inputValue, path) => path.some((option) => option.label.toLowerCase().indexOf(inputValue.toLowerCase()) > -1);

  // 用户登出方法
  const logoutHandler = async () => {
    try {
      const res = await LogoutRequest();
      if (res.code === 200) {
        localStorage.clear();
        message.success('用户注销成功');
        navigate('/login');
      } else {
        message.error(res.message);
      }
    } catch (e) {
      message.error('服务器异常，请联系管理员');
    }
  };

  // 下拉菜单
  const LayoutDropdownMenuData = [
    {
      key: '1',
      label: (
        <a rel="noopener noreferrer" href="">
          @{CurrentUserInfo?.en_name}（{CurrentUserInfo?.cn_name}）
        </a>
      ),
      disabled: true,
    },
    {
      type: 'divider',
    },
    {
      key: '2',
      label: (
        <a
          rel="noopener noreferrer"
          onClick={() => {
            navigate('/me');
          }}>
          个人中心
        </a>
      ),
    },
    {
      key: '3',
      label: (
        <a rel="noopener noreferrer" onClick={logoutHandler}>
          注销登录
        </a>
      ),
    },
  ];

  return (
    <Layout>
      <Sider className="admin-sider" width={menuWidth} collapsedWidth={menuCollapsedWidth} collapsible collapsed={collapsed} onCollapse={(value) => setCollapsed(value)}>
        <div
          className="admin-layout-logo"
          style={{
            width: collapsed ? menuCollapsedWidth + 'px' : menuWidth + 'px',
          }}>
          <img src={collapsed ? Logo : LogoWithTitle} alt="" />
        </div>
        <Menu
          className="admin-sider-menu"
          theme="dark"
          defaultSelectedKeys="['/dashboard']"
          openKeys={MenuOpenKeys}
          selectedKeys={MenuSelectKeys}
          mode="inline"
          items={menuTree}
          onOpenChange={(key) => {
            // 解决 404 等页码第一次点击折叠菜单不展开和收起菜单栏不选中问题
            // LayoutStates.MenuOpenKeys = [key[key.length - 1]];
            LayoutStates.MenuOpenKeys = key;
          }}
          // 菜单点击事件，能够返回对应的 Key
          // 文档中提示可获取到 item, key, keyPath, domEvent
          onClick={({ key }) => {
            navigate(key);
          }}
        />
      </Sider>
      <Layout
        style={{
          marginLeft: collapsed ? menuCollapsedWidth + 'px' : menuWidth + 'px',
          minHeight: '100vh',
          backgroundColor: '#ffffff',
        }}>
        <Header className="admin-header">
          <div className="admin-header-title">
            <Cascader
              style={{ width: '300px' }}
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
                <MoreOutlined
                  style={{
                    marginLeft: '5px',
                  }}
                />
              </div>
            </Dropdown>
          </div>
        </Header>
        <Content>
          <Outlet />
        </Content>
        <Footer className="admin-footer">
          <FooterText />
        </Footer>
      </Layout>
    </Layout>
  );
};

export default AdminLayout;
