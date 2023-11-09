import request from '@/utils/request';
import {ChatResponseData} from '@/types/APITypes';

export async function getHello(data: {
  mode: string;
}): Promise<ChatResponseData> {
  return (await request({
    url: '/api/hello',
    method: 'get',
    data,
  })) as never;
}

export async function getMessage(data: {
  mode: string;
}): Promise<ChatResponseData> {
  return (await request({
    url: '/api/message',
    method: 'get',
    data,
  })) as never;
}
