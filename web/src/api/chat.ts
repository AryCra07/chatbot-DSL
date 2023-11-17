import request from '@/utils/request';

export async function fetchHello(data: { name: string }): any {
  return await request.post('/api/user/hello', data);
}

export async function fetchMessage(data: { name: string; input: string }): any {
  return await request.post('/api/user/message', data);
}
