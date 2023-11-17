import request from '@/utils/request';

export async function Login(data: { name: string; password: string }): any {
  return await request.post('/api/user/login', data);
}
