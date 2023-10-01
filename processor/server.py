#  Copyright (c) 2023 AryCra07.

import grpc
from concurrent import futures
from pb import chat_pb2, chat_pb2_grpc
import logging

class AnswerGenerator(chat_pb2_grpc.AnswerGeneratorServicer):
    def Answer(self, request, context):
        logging.info(request)
        return chat_pb2.ChatResponse(answer="Hello " + request.userName + '!' + str(request.userState), state=-1)

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    logging.info("Server Start")
    chat_pb2_grpc.add_AnswerGeneratorServicer_to_server(AnswerGenerator(), server)
    server.add_insecure_port('[::]:8972')
    server.start()
    server.wait_for_termination()

if __name__ == '__main__':
    logging.basicConfig(level=logging.INFO, format='%(asctime)s - %(name)s - %(levelname)s - %(message)s')
    serve()