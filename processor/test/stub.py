#  Copyright (c) 2023 AryCra07.
import unittest
import grpc
from concurrent import futures
import pb.chat_pb2 as chat_pb2
import pb.chat_pb2_grpc as chat_pb2_grpc

class GreetServicer(chat_pb2_grpc.GreetServicer):
    def SayHelloService(self, request, context):
        response = chat_pb2.HelloResponse(words=['Hello'])
        return response

class ChatServicer(chat_pb2_grpc.ChatServicer):

    def AnswerService(self, request, context):
        # 在这里添加你要测试的逻辑
        # 例如，假设你的逻辑是简单地将输入返回作为答案
        response = chat_pb2.ChatResponse(state=request.state, answer=[request.input], wallet={'USD': 100.0, 'EUR': 50.0})
        return response

class TimerServicer(chat_pb2_grpc.TimerServicer):
    def TimerService(self, request, context):
        response = chat_pb2.TimerResponse(is_exit=False, reset=False, state=request.state, answer=['Timer'])
        return response

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    chat_pb2_grpc.add_ChatServicer_to_server(ChatServicer(), server)
    chat_pb2_grpc.add_GreetServicer_to_server(GreetServicer(), server)
    chat_pb2_grpc.add_TimerServicer_to_server(TimerServicer(), server)

    server.add_insecure_port('[::]:50052')
    server.start()
    return server