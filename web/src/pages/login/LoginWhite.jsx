import React from "react";
import {
  DingtalkOutlined,
  InsuranceOutlined,
  KeyOutlined,
  UserOutlined,
} from "@ant-design/icons";
import { Button, Checkbox, Divider, Form, Input } from "antd";

import "../../assets/css/login-white.less";

import { LogoWithTitleNoBg } from "../../common/Resource.jsx";
import QRCode from "../../assets/image/background/qr-code.png";
import TeamBgDots from "../../assets/image/background/team-bg-dots.png";
import TeamBgTriangle from "../../assets/image/background/team-bg-triangle.png";
import YellowTriangle from "../../assets/image/background/yellow-triangle.png";
import ServiceHalfCycle from "../../assets/image/background/service-half-cycle.png";
import SeoHalfCycle from "../../assets/image/background/seo-half-cycle.png";
import SeoBall from "../../assets/image/background/seo-ball.png";
import BlueHalfCycle from "../../assets/image/background/blue-half-cycle.png";
import FeatureBg from "../../assets/image/background/feature-bg.png";

// 注释
const LoginWhite = () => {
  const onFinish = (values) => {
    console.log("Received values of form: ", values);
  };
  return (
    <>
      <div className="hero">
        <img
          className="login-img-fluid login-change-bg-1 up-down-animation"
          src={FeatureBg}
          alt=""
        />
        <img
          className="login-img-fluid login-change-bg-5 up-down-animation"
          src={BlueHalfCycle}
          alt=""
        />
        <img
          className="login-img-fluid login-change-bg-4 up-down-animation"
          src={SeoBall}
          alt=""
        />
        <img
          className="login-img-fluid login-change-bg-3 left-right-animation"
          src={SeoHalfCycle}
          alt=""
        />
        <img
          className="login-img-fluid login-change-bg-8 left-right-animation"
          src={ServiceHalfCycle}
          alt=""
        />
        <img
          className="login-img-fluid login-change-bg-6 up-down-animation"
          src={YellowTriangle}
          alt=""
        />
        <img
          className="login-img-fluid login-change-bg-7 left-right-animation"
          src={TeamBgTriangle}
          alt=""
        />
        <img
          className="login-img-fluid login-change-bg-2 up-down-animation"
          src={TeamBgDots}
          alt=""
        />
      </div>
      <div className="login">
        <div className="login-header">
          <img src={LogoWithTitleNoBg} alt="" />
        </div>
        <div className="login-main">
          <div className="login-box">
            <div className="login-box-left">
              <img src={LogoWithTitleNoBg} alt="" />
              <div className="login-box-left-bg"></div>
            </div>
            <div className="login-box-right">
              <div className="login-qr-code">
                <img src={QRCode} alt="" />
              </div>
              <div className="login-title">Sign in</div>
              <Divider className="login-line">欢迎回来</Divider>
              <div>
                <Form
                  name="login"
                  className="login-form"
                  initialValues={{
                    remember: true,
                  }}
                  onFinish={onFinish}
                >
                  <Form.Item
                    name="account"
                    rules={[
                      {
                        required: true,
                        message: "请使用工号 / 手机号 / 邮箱登录!",
                      },
                    ]}
                  >
                    <Input
                      className="login-input"
                      prefix={<UserOutlined />}
                      placeholder="工号 / 手机号 / 邮箱"
                    />
                  </Form.Item>
                  <Form.Item
                    name="password"
                    rules={[
                      {
                        required: true,
                        message: "请输入你的密码!",
                      },
                    ]}
                  >
                    <Input.Password
                      className="login-input"
                      prefix={<KeyOutlined />}
                      placeholder="密码"
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
                  {/*      prefix={<MessageOutlined />}*/}
                  {/*      placeholder="短信验证码"*/}
                  {/*      style={{*/}
                  {/*        width: '170px',*/}
                  {/*      }}*/}
                  {/*    />*/}
                  {/*    <Button className="login-btn login-code-btn">获取验证码</Button>*/}
                  {/*  </Space>*/}
                  {/*</Form.Item>*/}
                  {/*<Form.Item>*/}
                  {/*  <Space direction="horizontal">*/}
                  {/*    <Input*/}
                  {/*        prefix={<MailOutlined />}*/}
                  {/*        placeholder="邮件验证码"*/}
                  {/*        style={{*/}
                  {/*          width: '170px',*/}
                  {/*        }}*/}
                  {/*    />*/}
                  {/*    <Button className="login-btn login-code-btn">获取验证码</Button>*/}
                  {/*  </Space>*/}
                  {/*</Form.Item>*/}
                  <Form.Item>
                    <Form.Item name="remember" valuePropName="checked" noStyle>
                      <Checkbox>记住密码 |</Checkbox>
                    </Form.Item>
                    <a href="">忘记密码？</a>
                  </Form.Item>

                  <Form.Item>
                    <Button
                      htmlType="submit"
                      className="login-btn login-form-button"
                    >
                      登录
                    </Button>
                    <div>
                      没有账号？
                      <a href="">
                        直接钉钉扫码自动创建 <DingtalkOutlined />
                      </a>
                    </div>
                  </Form.Item>
                </Form>
              </div>
            </div>
          </div>
        </div>
      </div>
    </>
  );
};

export default LoginWhite;
