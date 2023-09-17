import request from '@/utils/request';
import { ChatResponseData } from '@/types/APITypes';

export async function getStart(data: {
  mode: string;
}): Promise<ChatResponseData> {
  try {
    const response = (await request({
      url: '/api/bot/chat/init',
      method: 'get',
      data,
    })) as never;
    return response;
  } catch (e) {
    console.error(e);
  }
}

export async function getMessage(data: {
  mode: string;
}): Promise<ChatResponseData> {
  const response = (await request({
    url: '/api/bot/chat/message',
    method: 'get',
    data,
  })) as never;
  return response;
}
