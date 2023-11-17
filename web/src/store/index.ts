import { defineStore } from 'pinia';
import { useStorage } from '@vueuse/core';

export const settingStore = defineStore('user-setting', {
  state: () => {
    return {
      lang: useStorage('lang', 'zh-CN') as unknown as string,
      token: useStorage('token', '') as unknown as string,
      name: useStorage('name', 'agent') as unknown as string,
    };
  },
  actions: {
    setUserName(name: string) {
      this.name = name;
    },
    setToken(token: string) {
      this.token = token;
    },
    clearUserData() {
      this.userName = '';
      this.token = '';
    },
  },
  getters: {
    getUserName(): string {
      return this.name;
    },
    getToken(): string {
      return this.token;
    },
  },
});
