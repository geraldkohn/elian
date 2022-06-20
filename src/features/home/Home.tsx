import { Typography } from "antd";
import { Layout, Menu } from "antd";
import { Outlet, Link } from "react-router-dom";
import React from "react";

const { Header, Content, Footer, Sider } = Layout;
const { Title } = Typography;

const Home: React.FC = () => (
  <Layout>
    {/* Header */}
    <Header className="header" style={{ backgroundColor: "white" }}>
      <Title level={3} style={{ margin: "20px" }}>
        医疗数据共享平台
      </Title>
    </Header>

    {/* 主体内容 */}
    <Content style={{ padding: "0 50px" }}>
      <Layout className="site-layout-background" style={{ padding: "24px 0" }}>
        {/* 侧边栏 */}
        <Sider className="site-layout-background" width={200}>
          <Menu
            mode="inline"
            defaultSelectedKeys={["1"]}
            style={{ height: "100%" }}
          >
            <Menu.Item key={1}>
              <Link to={"/"}>首页</Link>
            </Menu.Item>

            <Menu.SubMenu title="信息管理" key={'sub1'}>
              <Menu.Item key={2}>
                <Link to="/userManager">用户管理</Link>
              </Menu.Item>
              <Menu.Item key={3}>
                <Link to="/notice">通知公告</Link>
              </Menu.Item>
              <Menu.Item key={4}>
                <Link to="/logManager">日志管理</Link>
              </Menu.Item>
              <Menu.Item key={5}>
                <Link to="/dataManager">数据管理</Link>
              </Menu.Item>
              <Menu.Item key={6}>
                <Link to="/currentPerformance">当前性能</Link>
              </Menu.Item>
            </Menu.SubMenu>

            <Menu.SubMenu title="系统管理" key={'sub2'}>
              <Menu.Item key={7}>
                <Link to="/deptManager">部门管理</Link>
              </Menu.Item>
              <Menu.Item key={8}>
                <Link to="/jobManager">岗位管理</Link>
              </Menu.Item>
              <Menu.Item key={9}>
                <Link to="/attrManager">属性管理</Link>
              </Menu.Item>
            </Menu.SubMenu>

            <Menu.SubMenu title="系统监控" key={'sub3'}>
              <Menu.Item key={10}>
                <Link to="/dataMonitor">数据监控</Link>
              </Menu.Item>
              <Menu.Item key={11}>
                <Link to="/onlinePeople">在线用户</Link>
              </Menu.Item>
              <Menu.Item key={12}>
                <Link to="/serviceMonitor">服务监控</Link>
              </Menu.Item>
            </Menu.SubMenu>
          </Menu>
        </Sider>

        {/* 内容 */}
        <Content style={{ padding: "0 24px", minHeight: 280 }}>
          <Outlet />
        </Content>
      </Layout>
    </Content>

    {/* Footer */}
    <Footer style={{ textAlign: "center" }}>
      Smart Health ©2022 Created by TJU Team
    </Footer>
  </Layout>
);

export default Home;
