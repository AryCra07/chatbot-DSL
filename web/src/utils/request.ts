import axios, { AxiosInstance } from 'axios';
import { ElMessage, ElMessageBox } from 'element-plus';
import i18n from '@/lang';
import { settingStore } from '@/store';
import router from '@/router';

const request: AxiosInstance = axios.create({
  baseURL: process.env.VITE_APP_API,
  timeout: 5000,
  headers: {
    'Content-Type': 'application/json;charset=utf-8',
  },
});

request.interceptors.request.use(
  (config) => {
    const store = settingStore();
    // do something before request is sent
    const token = store.getToken;

    if (token.length > 0) {
      const header = config.headers as any;
      header.Authorization = token;
    }
    return config;
  },
  (error) => {
    // do something with request error
    return Promise.reject(error);
  },
);

// response interceptor
request.interceptors.response.use(
  (response) => {
    const store = settingStore();
    const res = response.data;

    if (res.code !== 0) {
      ElMessage({
        message: res.msg || 'Error',
        type: 'error',
        duration: 1000,
      });

      // -2: wrong token
      if (res.code === -2) {
        // to re-login
        const t = i18n.global.t;
        ElMessageBox.confirm(t('re_login_text'), t('logout'), {
          confirmButtonText: t('re_login'),
          cancelButtonText: t('cancel'),
          type: 'warning',
        }).then(() => {
          store.setToken('');
          //push to login page
          router.push({path: '/login'}).then();
          // location.;
        });
      }
      return Promise.reject(new Error(res.msg || 'Error'));
    } else if (res.code === 0) {
      // store.setToken(res.data.token);
      return res;
    } else {
      return res;
    }
  },
  (error) => {
    return Promise.reject(error);
  },
);

export default request;
