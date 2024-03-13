import React from "react";
import "../../assets/css/login.less";
import { Button, Checkbox, Divider, Form, Input, Space } from "antd";
import { LogoWithTitleNoBg } from "../../common/Resource.jsx";
import {
  InsuranceOutlined,
  LockOutlined,
  MailOutlined,
  UserOutlined,
} from "@ant-design/icons";
import { APIConfig } from "../../common/Config.jsx";

// 用户登录页
const Login = () => {
  const LoginHandler = (values) => {
    console.log("Success:", values);
  };

  console.log(APIConfig.RunEnv);
  console.log(APIConfig.LoginAPI);

  return (
    <>
      <div className="login">
        <div className="login-container">
          <div className="login-header">
            <img src={LogoWithTitleNoBg} alt="" draggable="false" />
          </div>
          <div className="login-body">
            <div className="login-box">
              <div className="login-title">登录 / Sign in</div>
              <Divider className="login-line">欢迎回来</Divider>
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
                      placeholder="密码"
                    />
                  </Form.Item>

                  <Form.Item
                    name="code"
                    rules={[
                      {
                        required: true,
                        message: "请输入您的验证码!",
                      },
                    ]}
                  >
                    <Input
                      className="login-input"
                      prefix={<InsuranceOutlined />}
                      placeholder="验证码"
                    />
                  </Form.Item>

                  <Form.Item>
                    <Space direction="horizontal">
                      <Input
                        className="login-input"
                        prefix={
                          <MailOutlined className="site-forms-item-icon" />
                        }
                        placeholder="输入验证码"
                        style={{
                          width: "calc(310px - 108px)",
                        }}
                      />
                      <Button type="primary" className="login-code-button">
                        获取验证码
                      </Button>
                    </Space>
                  </Form.Item>

                  <Form.Item>
                    <Form.Item name="remember" valuePropName="checked" noStyle>
                      <Checkbox>记住密码</Checkbox>
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
      </div>
    </>
  );
};

export default Login;
