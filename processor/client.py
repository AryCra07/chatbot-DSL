#  Copyright (c) 2023 AryCra07.

from __future__ import print_function

import logging

import grpc
import pb.hello_pb2 as hp2
import pb.hello_pb2_grpc as hp2g


def run():
    # NOTE(gRPC Python Team): .close() is possible on a channel and should be
    # used in circumstances in which the with statement does not fit the needs
    # of the code.
    with grpc.insecure_channel('127.0.0.1:8972') as channel:
        stub = hp2g.GreeterStub(channel)
        resp = stub.SayHello(hp2.HelloRequest(name='AryCra07'))
    print("Greeter client received: " + resp.reply)


if __name__ == '__main__':
    logging.basicConfig()
    run()
