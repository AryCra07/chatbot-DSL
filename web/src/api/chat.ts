import request from '@/utils/request';

export async function fetchHello(): any {
  return await request.post('/api/user/hello');
}

export async function fetchMessage(data: { input: string }): any {
  return await request.post('/api/user/message', data);
}

export async function fetchTimer(data: { now_time: number }): any {
  return await request.post('/api/user/timer', data);
}
