import './index.scss';
import { Layout, Menu, Breadcrumb } from '@arco-design/web-react';
import { useState } from 'react';
import { IconHome, IconUser, IconSettings } from '@arco-design/web-react/icon';

const { Header, Sider, Content } = Layout;

export default function HomePage() {
  const [collapsed, setCollapsed] = useState(false);
  
  return (
    <Layout className="layout-wrapper">
      <Header className="header">
        <div className="logo">RBAC管理系统</div>
        <div className="user-info">
          <span>管理员</span>
        </div>
      </Header>
      <Layout>
        <Sider 
          collapsed={collapsed}
          onCollapse={() => setCollapsed(!collapsed)}
          width={220}
          className="sider"
        >
          <Menu
            defaultSelectedKeys={['1']}
            style={{ height: '100%' }}
          >
            <Menu.Item key="1">
              <IconHome />
              <span>首页</span>
            </Menu.Item>
            <Menu.Item key="2">
              <IconUser />
              <span>用户管理</span>
            </Menu.Item>
            <Menu.Item key="3">
              <IconSettings />
              <span>系统设置</span>
            </Menu.Item>
          </Menu>
        </Sider>
        <Content className="content">
          <Breadcrumb style={{ margin: '16px 0' }}>
            <Breadcrumb.Item>首页</Breadcrumb.Item>
          </Breadcrumb>
          <div className="main-content">
            <h1>欢迎来到RBAC系统</h1>
            <p>这是一个基于角色访问控制的系统</p>
          </div>
        </Content>
      </Layout>
    </Layout>
  );
}