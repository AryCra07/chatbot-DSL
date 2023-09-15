import { settingStore } from '@/store';

export function getToken(): string {
    const store = settingStore();
    const token = store.$state.token;
    if (token === null) {
        return '';
    }
    return token;
}

export function setToken(token: string): void {
    const store = settingStore();
    store.$state.token = token;
}

export function removeToken(): void {
    const store = settingStore();
    store.$state.token = '';
}
