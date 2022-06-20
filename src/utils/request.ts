import axios from "axios";

const service = axios.create({
  baseURL: "http://localhost:8080",
  timeout: 10000,
});

service.interceptors.request.use(
  (config: any) => {
    if (localStorage.token) {
      config.headers["Authorization"] = localStorage.token;
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

service.interceptors.response.use(
  function (response) {
    const res = response.data;

    if (res.state !== 200) {
    }
    return response;
  },
  function (error) {
    return Promise.reject(error);
  }
);

export default service;
