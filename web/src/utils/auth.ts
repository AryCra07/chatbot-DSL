import { settingStore } from '@/store';

export function setToken(token: string): void {
  const store = settingStore();
  console.log(store);
  store.setToken(token);
}

export function getToken(): string {
  const store = settingStore();
  const token = store.getToken;
  if (token === null || token === undefined) {
    return '';
  }
  return token;
}
