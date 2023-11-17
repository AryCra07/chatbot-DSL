import axios, { AxiosInstance } from 'axios';
import { ElMessage, ElMessageBox } from 'element-plus';
import i18n from '@/lang';
import { settingStore } from '@/store';

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
    console.log('token: ' + token);

    if (token.length > 0) {
      // let each request carry token
      // ['X-Token'] is a custom headers key
      // please modify it according to the actual situation
      // axios.defaults.headers.common['token'] = getToken()

      // console.log(axios.defaults.headers.common['token'])
      const header = config.headers as any;
      header.Authorization = token;
    }
    return config;
  },
  (error) => {
    // do something with request error
    console.log(error); // for debug
    return Promise.reject(error);
  },
);

// response interceptor
request.interceptors.response.use(
  /**
   * If you want to get http information such as headers or status
   * Please return  response => response
   */

  /**
   * Determine the request status by custom code
   * Here is just an example
   * You can also judge the status by HTTP Status Code
   */
  (response) => {
    const store = settingStore();
    const res = response.data;
    console.log(res);

    // if the custom code is not 20000, it is judged as an error.
    if (res.code !== 0) {
      ElMessage({
        message: res.msg || 'Error',
        type: 'error',
        duration: 5 * 1000,
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
          location.reload();
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
    console.log('err' + error); // for debug
    ElMessage({
      message: error.message,
      type: 'error',
      duration: 5 * 1000,
    });
    return Promise.reject(error);
  },
);

export default request;
