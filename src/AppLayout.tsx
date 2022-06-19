import {Typography} from 'antd';
import { Layout, Menu } from 'antd';
import { Outlet,Link } from 'react-router-dom';
import React from 'react';

const { Header, Content, Footer, Sider } = Layout;
const { Title } = Typography;


const AppLayout: React.FC = () => (
  <Layout>
    {/* Header */}
    <Header className="header" style={{backgroundColor: 'white'}}>
      <Title level={3} style={{margin: "20px"}}>
        医疗数据共享平台
      </Title>
    </Header>

    {/* 主体内容 */}
    <Content style={{ padding: '0 50px' }}>

      <Layout className="site-layout-background" style={{ padding: '24px 0' }}>
        {/* 侧边栏 */}
        <Sider className="site-layout-background" width={200}>
          <Menu
            mode="inline"
            defaultSelectedKeys={['1']}
            defaultOpenKeys={['sub1']}
            style={{ height: '100%' }}>
                <Menu.Item key={1}> 
                    <Link to={'/'}>首页</Link>
                </Menu.Item>

                <Menu.SubMenu title="信息管理" key={1}>
                <Menu.Item key={2}>
                        <Link to='/home'>用户管理</Link>
                    </Menu.Item>
                    <Menu.Item key={2}>
                        <Link to='/home'>通知公告</Link>
                    </Menu.Item>
                    <Menu.Item key={2}>
                        <Link to='/home'>日志管理</Link>
                    </Menu.Item>
                    <Menu.Item key={2}>
                        <Link to='/home'>数据管理</Link>
                    </Menu.Item>
                    <Menu.Item key={2}>
                        <Link to='/home'>当前性能</Link>
                    </Menu.Item>
                </Menu.SubMenu>

                <Menu.SubMenu title="系统管理">
                    <Menu.Item key={2}>
                        <Link to='/home'>部门管理</Link>
                    </Menu.Item>
                    <Menu.Item key={2}>
                        <Link to='/home'>岗位管理</Link>
                    </Menu.Item>
                    <Menu.Item key={2}>
                        <Link to='/home'>属性管理</Link>
                    </Menu.Item>
                </Menu.SubMenu>

                <Menu.SubMenu title="系统监控">
                    <Menu.Item key={2}>
                        <Link to='/home'>数据监控</Link>
                    </Menu.Item>
                    <Menu.Item key={2}>
                        <Link to='/home'>在线用户</Link>
                    </Menu.Item>
                    <Menu.Item key={2}>
                        <Link to='/home'>服务监控</Link>
                    </Menu.Item>
                </Menu.SubMenu>
          </Menu>
        </Sider>

        {/* 内容 */}
        <Content style={{ padding: '0 24px', minHeight: 280 }}> 
            <Outlet/>
        </Content>
      </Layout>
    </Content>

    {/* Footer */}
    <Footer style={{ textAlign: 'center' }}>Smart Health ©2022 Created by TJU Team</Footer>
  </Layout>
);

export default AppLayout;