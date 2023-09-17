export interface ChatResponseData {
  code: number;
  msg: number;
  data: Array<{ content: string; time: number }>;
}
