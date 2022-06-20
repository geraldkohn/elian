import React from "react";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import "./App.css";
import Home from "./features/home/Home";
import { Login } from "./features/login/Login";
import { Register } from "./features/register/Register";

import Main from "./pages/Main";

import CurrentPerformance from "./pages/infoManager/CurrentPerformance";
import DataManager from "./pages/infoManager/DataManager";
import LogManager from "./pages/infoManager/LogManager";
import Notice from "./pages/infoManager/Notice";
import UserManage from "./pages/infoManager/UserManager";

import DeptManager from "./pages/systemManager/DeptManager";
import JobManager from "./pages/systemManager/JobManager";
import AttrManager from "./pages/systemManager/AttrManager";

import DataMonitor from "./pages/systemMonitor/DataMonitor";
import OnlinePeople from "./pages/systemMonitor/OnlinePeople";
import ServiceMonitor from "./pages/systemMonitor/ServiceMonitor";

function App() {
  return (
    <div className="App">
      <BrowserRouter>
        <Routes>
          <Route path="/login" element={<Login />} />
          <Route path="/register" element={<Register />} />
          <Route path="/" element={<Home />}>

            <Route index element={<Main />}></Route>

            <Route path="/currentPerformance" element={<CurrentPerformance />}></Route>
            <Route path="/dataManager" element={<DataManager />}></Route>
            <Route path="/logManager" element={<LogManager />}></Route>
            <Route path="/notice" element={<Notice />}></Route>
            <Route path="/userManager" element={<UserManage />}></Route>

            <Route path="/deptManager" element={<DeptManager />}></Route>
            <Route path="/jobManager" element={<JobManager />}></Route>
            <Route path="/attrManager" element={<AttrManager />}></Route>

            <Route path="/dataMonitor" element={<DataMonitor />}></Route>
            <Route path="/onlinePeople" element={<OnlinePeople />}></Route>
            <Route path="/serviceMonitor" element={<ServiceMonitor />}></Route>
          </Route>
        </Routes>
      </BrowserRouter>
    </div>
  );
}

export default App;
