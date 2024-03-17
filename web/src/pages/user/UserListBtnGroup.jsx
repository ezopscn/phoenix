import React from 'react';
import { Button, Col, Dropdown, Row, Space } from 'antd';
import { DownloadOutlined, DownOutlined, UploadOutlined, UserAddOutlined } from '@ant-design/icons';

// 用户批量操作下拉菜单
const userMultiHandleItems = [
  {
    label: '锁定选中用户',
    key: '1',
  },
  {
    label: '解锁选中用户',
    key: '2',
  },
  {
    label: '激活选中用户',
    key: '3',
  },
  {
    label: '删除选中用户',
    key: '4',
  },
];

// 执行方法
const userMultiHandle = (e) => {
  console.log('数据：', e);
};

// 对象
const userMultiHandleProps = {
  items: userMultiHandleItems,
  onClick: userMultiHandle,
};

// 用户按钮组
const UserListBtnGroup = () => {
  return (
    <>
      {/* 按钮组定义 */}
      <Row>
        <Col span={12} className="admin-btn-group-left">
          <Button type="primary" icon={<UserAddOutlined />}>
            新建用户
          </Button>
          <Dropdown menu={userMultiHandleProps}>
            <Button danger>
              <Space>
                <DownOutlined />
                批量操作
              </Space>
            </Button>
          </Dropdown>
        </Col>
        <Col span={12} className="admin-align-right admin-btn-group-right">
          <Button icon={<DownloadOutlined />}>下载模板</Button>
          <Button icon={<UploadOutlined />}>导入用户</Button>
        </Col>
      </Row>
    </>
  );
};

export default UserListBtnGroup;
