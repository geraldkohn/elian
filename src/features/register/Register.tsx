import { Button, Form, Input, Card, message } from "antd";
import React from "react";
import { patientLogin } from "../../api/user";
import { useNavigate } from "react-router-dom";
import { RegisterBox } from "./Register.module";

export const Register: React.FC = () => {
  const navigate = useNavigate();
  const onFinish = (data: any) => {
    patientLogin(data)
      .then((res) => {
        navigate("/");
      })
      .catch((res) => {
        message.error("后端返回的错误信息");
      });
  };

  const onFinishFailed = (errorInfo: any) => {
    console.log("Failed:", errorInfo);
  };

  return (
    <RegisterBox>
      <Card
        title="注册"
        style={{ backgroundColor: "#fcfcfc" }}
        hoverable={true}
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

          <Form.Item wrapperCol={{ offset: 4, span: 16 }}>
            <Button type="primary" htmlType="submit">
              注册
            </Button>
          </Form.Item>
        </Form>
      </Card>
    </RegisterBox>
  );
};
