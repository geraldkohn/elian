import { Button,Form, Input, Card ,message } from "antd";
import React from "react";

import { LoginBox } from "./Login.module";
import { login } from "../../api/user";
import { useNavigate } from "react-router-dom";

export const Login: React.FC = () => {
  const navigate=useNavigate()
  const onFinish = (data: any) => {
    login(data).then(res=>{
      navigate('/')
    }).catch(res=>{
      message.error("后端返回的错误信息")
    })
  };

  const onFinishFailed = (errorInfo: any) => {
    console.log("Failed:", errorInfo);
  };

  return (
    <LoginBox>
      <Card
        title="基于区块链的医疗数据共享平台"
        style={{backgroundColor: "#fcfcfc"}}
        hoverable= {true}
      >
        <br />
        <Form
          name="basic"
          labelCol={{ span: 6 }}
          wrapperCol={{ span: 15 }}
          initialValues={{ remember: true }}
          onFinish={onFinish}
          onFinishFailed={onFinishFailed}
          autoComplete="off"
        >
          <Form.Item
            label="身份证"
            name="id_card_number"
            rules={[{ required: true, message: "请输入您的身份证号码" }]}
          >
            <Input />
          </Form.Item>

          <Form.Item
            label="密码"
            name="password"
            rules={[{ required: true, message: "请输入您的密码" }]}
          >
            <Input.Password />
          </Form.Item>

          <Form.Item wrapperCol={{ offset: 5, span: 16 }}>
            <Button type="primary" htmlType="submit">
              登录
            </Button>
          </Form.Item>
        </Form>
      </Card>
    </LoginBox>
  );
};
