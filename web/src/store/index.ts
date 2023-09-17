import { defineStore } from 'pinia';
import { useStorage } from '@vueuse/core';

export const settingStore = defineStore({
  id: 'setting',
  state: () => ({
    lang: useStorage('lang', 'zh') as unknown as string,
    token: useStorage('token', '') as unknown as string,
    name: useStorage('name', 'user') as unknown as string,
  }),
  getters: {},
  actions: {},
});
