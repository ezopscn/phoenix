import React from "react";
import "../../assets/css/login.less";
import { Button, Checkbox, Divider, Form, Input, message, Space } from "antd";
import { LogoWithTitleNoBg } from "../../common/Resource.jsx";
import {
  DingtalkOutlined,
  InsuranceOutlined,
  LockOutlined,
  MailOutlined,
  UserOutlined,
} from "@ant-design/icons";
import { LoginRequest } from "../../utils/RequestAPI.jsx";
import { SetToken } from "../../utils/Token.jsx";
import { useNavigate } from "react-router";

// 用户登录页
const Login = () => {
  // 路由跳转
  const navigate = useNavigate();

  // 登录请求
  const loginHandler = async (data) => {
    try {
      const res1 = await LoginRequest(data);
      if (res1.code === 200) {
        SetToken(res1.data.token, res1.data.expire);
        navigate("/");
      } else {
        message.error(res1.message);
      }
    } catch (e) {
      message.error("服务器异常，请联系管理员");
    }
  };

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
                  onFinish={loginHandler}
                >
                  <Form.Item
                    className="login-form-item"
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
                    className="login-form-item"
                    name="password"
                    rules={[
                      {
                        required: true,
                        message: "请输入您的密码!",
                      },
                    ]}
                  >
                    <Input.Password
                      autoComplete="off"
                      className="login-input"
                      prefix={<LockOutlined className="site-form-item-icon" />}
                      type="password"
                      placeholder="密码"
                    />
                  </Form.Item>

                  {/*手机令牌方式*/}
                  <Form.Item
                    className="login-form-item"
                    name="code"
                    rules={[
                      {
                        required: true,
                        message: "请输入您的验证码!",
                      },
                    ]}
                  >
                    <Input
                      autoComplete="off"
                      className="login-input"
                      prefix={<InsuranceOutlined />}
                      placeholder="手机令牌验证码"
                    />
                  </Form.Item>

                  {/*邮件短信获取验证码方式*/}
                  <Form.Item>
                    <Space direction="horizontal">
                      <Input
                        autoComplete="off"
                        className="login-input"
                        prefix={
                          <MailOutlined className="site-forms-item-icon" />
                        }
                        placeholder="邮件 / 短信验证码"
                        style={{
                          width: "calc(310px - 108px)",
                        }}
                      />
                      <Button type="primary" className="login-code-button">
                        获取验证码
                      </Button>
                    </Space>
                  </Form.Item>

                  <Form.Item className="login-remember-item">
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

                  <Divider className="login-line-change">
                    或者使用钉钉扫码直接登录
                  </Divider>

                  <Button block>
                    <DingtalkOutlined /> 切换到钉钉扫码登录
                  </Button>
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
