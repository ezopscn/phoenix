import React, { useState } from 'react';
import { useSnapshot } from 'valtio';
import { DepartmentStates, RegionStates, RoleStates, UserStates } from '../../../store/Store.jsx';
import { Button, Col, Form, Input, Row, Select } from 'antd';
import { DownOutlined, UpOutlined } from '@ant-design/icons';

// 可供搜索的字段
const userListSearchFields = [
  { name: 'username', label: '用户名', type: 'text' },
  { name: 'name', label: '姓名', type: 'text' },
  { name: 'mobile', label: '手机号', type: 'text' },
  { name: 'email', label: '邮箱', type: 'text' },
  { name: 'job_number', label: '工号', type: 'text' },
  { name: 'system_role_id', label: '角色', type: 'select' },
  { name: 'job_name', label: '岗位名称', type: 'text' },
  { name: 'system_department_id', label: '部门名称', type: 'select' },
  { name: 'native_province_id', label: '籍贯省份', type: 'select' },
  { name: 'native_city_id', label: '籍贯城市', type: 'select' },
  { name: 'active', label: '激活状态', type: 'select' },
  { name: 'unlocked', label: '锁定状态', type: 'select' },
  { name: 'gender', label: '性别', type: 'select' },
  { name: 'creator', label: '创建人', type: 'text' },
  { name: 'office_province_id', label: '办公省份', type: 'select' },
  { name: 'office_city_id', label: '办公城市', type: 'select' },
  { name: 'office_address', label: '办公地点', type: 'text' },
];

// 生成搜索表单
const generateUserListSearchForm = () => {
  // 全局状态信息
  const { Provinces } = useSnapshot(RegionStates); // 地区信息
  const { Departments } = useSnapshot(DepartmentStates); // 部门信息
  const { Roles } = useSnapshot(RoleStates); // 角色信息
  const { UserSearchFieldExpand } = useSnapshot(UserStates); // 用户信息

  // 联动城市数据
  const [searchOfficeCities, setSearchOfficeCities] = useState([]);
  const [searchNativeCities, setSearchNativeCities] = useState([]);

  // 默认显示的搜索框数量
  let count = 7;

  // 如果用户展开了更多搜索，则为定义的搜索列表长度
  if (UserSearchFieldExpand) {
    count = userListSearchFields.length;
  }

  // 用于存储传递给表单的数据
  const children = [];

  // 遍历生成搜索框数据
  for (let i = 0; i < count; i++) {
    // 每项数据和唯一的 Key
    let item = userListSearchFields[i];

    // 输入提示信息，Placeholder
    let tp = '通过输入' + item.label + '进行搜索';
    let sp = '通过选择' + item.label + '进行搜索';

    // 生成搜索框，并保存起来
    children.push(
      <Col span={6} key={item.name}>
        <Form.Item name={item.name} label={item.label} colon={false} labelCol={{ span: 6 }} wrapperCol={{ span: 18 }}>
          {/* Select 数据需要根据实际选择 */}
          {item.name === 'gender' ? (
            // 性别
            <Select initialvalues="1" placeholder={sp}>
              <Option value="1">男</Option>
              <Option value="2">女</Option>
              <Option value="3">未知</Option>
            </Select>
          ) : item.name === 'active' ? (
            // 激活
            <Select initialvalues="1" placeholder={sp}>
              <Option value="0">未激活</Option>
              <Option value="1">已激活</Option>
            </Select>
          ) : item.name === 'unlocked' ? (
            // 锁定
            <Select initialvalues="1" placeholder={sp}>
              <Option value="0">已锁定</Option>
              <Option value="1">未锁定</Option>
            </Select>
          ) : item.name === 'system_role_id' ? (
            // 角色
            <Select placeholder={sp} showSearch={true} filterOption={(input, option) => option.label.toLowerCase().includes(input.toLowerCase())}>
              {Roles.map((item) => (
                <Select.Option key={item.id} label={item.name} value={item.id}>
                  {item.name}
                </Select.Option>
              ))}
            </Select>
          ) : item.name === 'system_department_id' ? (
            // 部门
            <Select placeholder={sp} showSearch={true} filterOption={(input, option) => option.label.toLowerCase().includes(input.toLowerCase())}>
              {Departments.map((item) => (
                <Select.Option key={item.id} label={item.name} value={item.id}>
                  {item.name}
                </Select.Option>
              ))}
            </Select>
          ) : item.name === 'office_province_id' ? (
            // 办公省份
            <Select
              placeholder={sp}
              showSearch={true}
              filterOption={(input, option) => option.label.toLowerCase().includes(input.toLowerCase())}
              onChange={(id) => {
                // GetCityListByProvinceIdHandle(id).then((val) =>
                //   setSearchOfficeCities(val),
                // );
              }}>
              {Provinces.map((item) => (
                <Select.Option key={item.id} label={item.name} value={item.id}>
                  {item.name}
                </Select.Option>
              ))}
            </Select>
          ) : item.name === 'office_city_id' ? (
            // 办公城市
            <Select placeholder={sp} showSearch={true} filterOption={(input, option) => option.label.toLowerCase().includes(input.toLowerCase())}>
              {searchOfficeCities.map((item) => (
                <Select.Option key={item.id} label={item.name} value={item.id}>
                  {item.name}
                </Select.Option>
              ))}
            </Select>
          ) : item.name === 'native_province_id' ? (
            // 籍贯省份
            <Select
              placeholder={sp}
              showSearch={true}
              filterOption={(input, option) => option.label.toLowerCase().includes(input.toLowerCase())}
              onChange={(id) => {
                // GetCityListByProvinceIdHandle(id).then((val) =>
                //   setSearchNativeCities(val),
                // );
              }}>
              {Provinces.map((item) => (
                <Select.Option key={item.id} label={item.name} value={item.id}>
                  {item.name}
                </Select.Option>
              ))}
            </Select>
          ) : item.name === 'native_city_id' ? (
            // 籍贯城市
            <Select placeholder={sp} showSearch={true} filterOption={(input, option) => option.label.toLowerCase().includes(input.toLowerCase())}>
              {searchNativeCities.map((item) => (
                <Select.Option key={item.id} label={item.name} value={item.id}>
                  {item.name}
                </Select.Option>
              ))}
            </Select>
          ) : (
            // 默认
            <Input placeholder={tp} />
          )}
        </Form.Item>
      </Col>,
    );
  }

  return children;
};

// 用户搜索
const UserListSearch = () => {
  const [form] = Form.useForm();

  // 全局状态
  const { UserSearchFieldExpand } = useSnapshot(UserStates); // 用户信息

  return (
    <>
      <Form form={form} name="user-search">
        <Row gutter={24}>
          {/*搜索框*/}
          {generateUserListSearchForm()}
          {/*搜索按钮*/}
          <Col span={6} key="search-submit" className="admin-submit-btn-group admin-btn-group-left">
            <Button type="primary" htmlType="submit">
              搜索用户
            </Button>
            <Button
              onClick={() => {
                form.resetFields();
              }}>
              清空条件
            </Button>
            <a
              onClick={() => {
                UserStates.UserSearchFieldExpand = !UserSearchFieldExpand;
              }}>
              {UserSearchFieldExpand ? <UpOutlined /> : <DownOutlined />}
              {UserSearchFieldExpand ? ' 收起选项' : ' 更多选项'}
            </a>
          </Col>
        </Row>
      </Form>
    </>
  );
};

export default UserListSearch;
