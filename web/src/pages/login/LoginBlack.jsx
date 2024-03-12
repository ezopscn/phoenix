import React from "react";
import "../../assets/css/login-black.less";
import { Button, Checkbox, Divider, Form, Input } from "antd";
import { LogoWithTitle } from "../../common/Resource.jsx";
import {
  InsuranceOutlined,
  LockOutlined,
  UserOutlined,
} from "@ant-design/icons"; // 用户登录页

// 用户登录页
const LoginBlack = () => {
  const LoginHandler = (values) => {
    console.log("Success:", values);
  };

  return (
    <>
      <div className="login-container">
        <div className="login-header">
          <img src={LogoWithTitle} alt="" draggable="false" />
        </div>
        <div className="login-body">
          <div className="login-box">
            <div className="login-title">登录 / Sign in</div>
            <Divider className="login-line">HI . 欢迎回来</Divider>
            <div className="login-form">
              <Form
                name="login"
                initialValues={{
                  remember: true,
                }}
                onFinish={LoginHandler}
              >
                <Form.Item
                  name="account"
                  rules={[
                    {
                      required: true,
                      message: "请输入您的用户名!",
                    },
                  ]}
                >
                  <Input
                    autoComplete="off"
                    className="login-input"
                    prefix={<UserOutlined className="site-form-item-icon" />}
                    placeholder="工号 / 手机号 / Email"
                  />
                </Form.Item>
                <Form.Item
                  name="password"
                  rules={[
                    {
                      required: true,
                      message: "请输入您的密码!",
                    },
                  ]}
                >
                  <Input.Password
                    className="login-input"
                    prefix={<LockOutlined className="site-form-item-icon" />}
                    type="password"
                    placeholder="用户密码"
                  />
                </Form.Item>

                <Form.Item>
                  <Input
                    className="login-input"
                    prefix={<InsuranceOutlined />}
                    placeholder="VIP ACCESS 验证码"
                  />
                </Form.Item>

                {/*<Form.Item>*/}
                {/*  <Space direction="horizontal">*/}
                {/*    <Input*/}
                {/*      prefix={<MailOutlined className="site-forms-item-icon" />}*/}
                {/*      placeholder="输入验证码"*/}
                {/*      style={{*/}
                {/*        width: 'calc(330px - 108px)',*/}
                {/*      }}*/}
                {/*    />*/}
                {/*    <Button*/}
                {/*      type="primary"*/}
                {/*      style={{ width: 100, letterSpacing: 1 }}*/}
                {/*    >*/}
                {/*      获取验证码*/}
                {/*    </Button>*/}
                {/*  </Space>*/}
                {/*</Form.Item>*/}

                <Form.Item>
                  <Form.Item name="remember" valuePropName="checked" noStyle>
                    <Checkbox style={{ color: "#fff" }}>记住密码</Checkbox>
                  </Form.Item>

                  <a className="login-form-forgot" href="">
                    忘记密码？
                  </a>
                </Form.Item>

                <Form.Item>
                  <Button
                    block
                    type="primary"
                    htmlType="submit"
                    className="login-form-button"
                  >
                    登录
                  </Button>
                </Form.Item>
              </Form>
            </div>
          </div>
        </div>
        <div className="login-footer">
          <b>🧬 PHOENIX </b>© 2024 EZOPS.CN, All Rights Reserved.
        </div>
      </div>
    </>
  );
};

export default LoginBlack;
