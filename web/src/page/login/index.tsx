import {
  Button,
  Card,
  Form,
  Input,
  Space,
  Typography,
  Message,
} from '@arco-design/web-react';
import { useNavigate } from 'react-router-dom';
import './index.scss';

export default function LoginPage() {
  const navigate = useNavigate();
  
  const handleSubmit = (values: any) => {
    // 模拟登录成功
    Message.success('登录成功！');
    // 跳转到首页
    navigate('/');
  };

  return (
    <div className="login-container">
      <Card className="login-card" hoverable>
        <Space direction="vertical" size="large">
          <Typography.Title heading={4} className="login-title">
            系统登录
          </Typography.Title>
          <Form layout="vertical" onSubmit={handleSubmit}>
            <Form.Item label="用户名" field="username" required>
              <Input placeholder="请输入用户名" allowClear />
            </Form.Item>
            <Form.Item label="密码" field="password" required>
              <Input.Password placeholder="请输入密码" allowClear />
            </Form.Item>
            <Button type="primary" long htmlType="submit">
              登录
            </Button>
          </Form>
        </Space>
      </Card>
    </div>
  );
}
