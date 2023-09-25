import request from '@/utils/request';

export async function Sign(data: {
  username: string;
  pwd: string;
}): Promise<{ code: number; msg: number; data: void }> {
  const resp = (await request({
    url: '/api/auth/sign',
    method: 'post',
    data,
  })) as never;
  return resp;
}
