#  Copyright (c) 2023 AryCra07.

from __future__ import print_function

import logging

import grpc
from pb import chat_pb2, chat_pb2_grpc

def run():
    # NOTE(gRPC Python Team): .close() is possible on a channel and should be
    # used in circumstances in which the with statement does not fit the needs
    # of the code.
    with grpc.insecure_channel('127.0.0.1:8972') as channel:
        stub = chat_pb2_grpc.AnswerGeneratorStub(channel)
        resp = stub.Answer(chat_pb2.ChatRequest(userName='AryCra07', userInput='Heyy', userState = 0))
    print("Chat client received: " + resp.answer)


if __name__ == '__main__':
    logging.basicConfig()
    run()
