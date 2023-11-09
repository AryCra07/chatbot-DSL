import axios, { AxiosInstance } from 'axios';
import { ElMessage, ElMessageBox } from 'element-plus';

const service: AxiosInstance = axios.create({
  baseURL: 'api/',
  timeout: 5000,
  headers: {
    'Content-Type': 'application/json;charset=utf-8',
  },
});

service.interceptors.request.use(
  (config) => {
    return config;
  },
  (error) => {
    console.log(error);
    return Promise.reject(error);
  },
);

service.interceptors.response.use(
  (response) => {
    const code = response.data.code;
    const message = response.data.message;

    if (code !== 0 && code !== 1) {
      ElMessage({
        message: response.config.url + ': ' + message || 'Error',
        type: 'error',
        duration: 5 * 1000,
      });
      console.log(response);
      return Promise.reject(new Error(message || 'Error'));
    } else {
      return response.data;
    }
  },
  (error) => {
    if (error.response.status === 401) {
      ElMessageBox.confirm(
        'You have been logged out, you can cancel to stay on this page, or log in again',
        'Confirm Logout',
        {
          confirmButtonText: 'Re-Login',
          cancelButtonText: 'Cancel',
          type: 'warning',
          closeOnClickModal: false,
        },
      ).then(() => {
        console.log(1);
      });
    }
  },
);

export default service;
