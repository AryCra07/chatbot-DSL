import request from '@/utils/request';

export async function Login(data: {
  username: string;
  pwd: string;
}): Promise<{ code: number; msg: number; data: void }> {
  return (await request({
    url: '/api/login',
    method: 'post',
    data,
  })) as never;
}
