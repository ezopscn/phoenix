import React from 'react';
import ButterflyPageHeader from '../../components/page-header/PageHeader.jsx';
import UserListSearch from './UserListSearch.jsx';
import UserListTable from './UserListTable.jsx';
import UserListBtnGroup from './UserListBtnGroup.jsx';

// 提示信息
const UserNotices = () => {
  return (
    <>
      <ul>
        <li>管理员用户可以对用户进行搜索查看，编辑修改，禁用启用等操作。</li>
        <li>用户状态说明：禁用（可以用于替代用户删除）。</li>
      </ul>
    </>
  );
};

const UserList = () => {
  const title = '用户管理'; // 页面标题

  return (
    <>
      {/*提示信息*/}
      <div className="admin-tips">
        <ButterflyPageHeader title={title} notices={<UserNotices />} />
      </div>

      {/*搜索*/}
      <div className="admin-search">
        <UserListSearch />
      </div>

      {/*按钮组*/}
      <div className="admin-btn-group">
        <UserListBtnGroup />
      </div>

      {/*列表*/}
      <div className="admin-table">
        <UserListTable />
      </div>
    </>
  );
};

export default UserList;
