import React, { useEffect, useState } from 'react';
import { Avatar, Badge, Descriptions, message, Space, Table, Tag } from 'antd';
import { UserStates } from '../../../store/Store.jsx';
import { useSnapshot } from 'valtio';
import { UserListRequest } from '../../../utils/RequestAPI.jsx';
import moment from 'moment';

// 用户表格定义
const userListColumns = [
  Table.EXPAND_COLUMN, // 展开
  Table.SELECTION_COLUMN, // 选择
  {
    title: '用户',
    dataIndex: 'avatar',
    width: '60px',
    align: 'center',
    render: (text) => <Avatar src={text} size={18} />,
  },
  {
    title: '英文名',
    dataIndex: 'en_name',
  },
  {
    title: '中文名',
    dataIndex: 'cn_name',
  },
  {
    title: '性别',
    dataIndex: 'gender',
    width: '60px',
    render: (gender) => (gender === 1 ? <Tag color="blue">男</Tag> : gender === 2 ? <Tag color="magenta">女</Tag> : <Tag color="yellow">未知</Tag>),
  },
  {
    title: '手机号',
    dataIndex: 'phone',
  },
  {
    title: '邮箱',
    dataIndex: 'email',
  },
  {
    title: '工号',
    dataIndex: 'job_id',
  },
  {
    title: '部门',
    dataIndex: ['department', 'name'],
  },
  {
    title: '职位',
    dataIndex: 'job_name',
  },
  {
    title: '角色',
    render: (record) =>
      record.role?.keyword === 'administrator' ? (
        <Tag color="red">
          {record.role?.name} / {record.role?.keyword}
        </Tag>
      ) : (
        <Tag className="admin-gray-tag">
          {record.role?.name} / {record.role?.keyword}
        </Tag>
      ),
  },
  {
    title: '状态',
    dataIndex: 'status',
    align: 'center',
    render: (status) => (status === 1 ? <Badge status="success" /> : <Badge status="error" />),
  },
  {
    title: '操作',
    align: 'center',
    render: (_, record) => (
      <Space size="middle">
        <a
          onClick={() => {
            console.log(record.job_id);
            // UserStates.UserEditModelOpen = true;
            // UserStates.EditUserInfo = record;
          }}>
          修改
        </a>
        <a>禁用</a>
        <a>锁定</a>
        <a>删除</a>
      </Space>
    ),
  },
];

// 选中用户
const userListRowSelection = {
  onChange: (selectedRowKeys, selectedRows) => {
    console.log(`selectedRowKeys: ${selectedRowKeys}`, 'selectedRows: ', selectedRows);
  },
  // getUserListCheckboxProps: (record) => ({
  //   disabled: record.username === 'admin',
  //   username: record.username,
  // }),
};

// 用户列表
const UserListTable = () => {
  // 局部状态
  const [userList, setUserList] = useState([]);
  const [isLoading, setIsLoading] = useState(false);

  // 用户搜索参数
  const { UserListSearchParams } = useSnapshot(UserStates);

  // 数据总量
  const [pageTotal, setPageTotal] = useState(0);

  // 请求数据
  const [requestPageParams, setRequestPageParams] = useState({
    // 分页数据
    no_pagination: false, // 是否不分页
    page_number: 1, // 默认页码
    page_size: 1, // 每页显示数量
  });

  // 组合对象
  const userListParams = { ...requestPageParams, ...UserListSearchParams };

  // 请求用户列表
  useEffect(() => {
    // 获取用户列表，使用异步会导致短暂的 Warning 提示
    (async () => {
      try {
        const res = await UserListRequest({ params: userListParams });
        if (res.code === 200) {
          setUserList(res.data.list);
          setIsLoading(false);
          // 设置分页信息
          setPageTotal(res.data.page.total);
        } else {
          message.error(res.message);
        }
      } catch (e) {
        console.log(e);
        message.error('服务器异常，请联系管理员');
      }
    })();
  }, [requestPageParams]); // 次数不能跟踪 userListParams，会死循环

  return (
    <>
      {/*数据*/}
      {isLoading ? (
        <div>正在加载中...</div>
      ) : (
        <Table
          rowSelection={{
            type: 'checkbox',
            ...userListRowSelection,
          }}
          expandable={{
            expandedRowRender: (record) => (
              <div className="admin-list-expand-content">
                <Descriptions column={1} size="small">
                  <Descriptions.Item label="用户账户">{record.job_id}</Descriptions.Item>
                  <Descriptions.Item label="用户籍贯">
                    {record.native_province.name} - {record.native_city.name}
                  </Descriptions.Item>
                  <Descriptions.Item label="办公地点">
                    {record.office_province.name} - {record.office_city.name} - {record.office_area.name} - {record.office_street.name} - {record.office_address} - {record.office_station}
                  </Descriptions.Item>
                  <Descriptions.Item label="管理人员">{record.leader === 1 ? <span style={{ color: '#cf1322' }}>是</span> : '否'}</Descriptions.Item>
                  <Descriptions.Item label="入职时间">{moment(record.join_time).format('YYYY-MM-DD')}</Descriptions.Item>
                  <Descriptions.Item label="用户生日">{moment(record.birthday).format('YYYY-MM-DD')}</Descriptions.Item>
                  <Descriptions.Item label="创建时间">{record.created_at}</Descriptions.Item>
                  <Descriptions.Item label="最后登录">{record.last_login_time}</Descriptions.Item>
                </Descriptions>
              </div>
            ),
          }}
          columns={userListColumns} // 列
          dataSource={userList} // 用户数据
          bordered
          rowKey="id"
          size="small"
          pagination={{
            total: pageTotal,
            showTotal: () => '总共 ' + pageTotal + ' 条数据',
            defaultCurrent: requestPageParams.page_number,
            defaultPageSize: requestPageParams.page_size,
            showSizeChanger: true,
            hideOnSinglePage: true,
            onChange: (page, pageSize) => {
              setRequestPageParams({
                page_number: page,
                page_size: pageSize,
              });

              // 重置为空，用于解决切换页面显示数量由大变小偶尔出现报错的问题。
              // 原因在于：由于异步请求还未完成，导致列表溢出的问题，虽然该问题会在请求完成后自己解决，但是控制台会提示：
              // Warning: [antd: Table] `dataSource` length is less than `pagination.total` but large than `pagination.pageSize`.
              // Please make sure your config correct data with async mode.
              if (userList.length > pageSize) {
                setUserList([]);
              }
            },
          }}
        />
      )}
    </>
  );
};

export default UserListTable;
