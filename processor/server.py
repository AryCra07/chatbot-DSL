#  Copyright (c) 2023 AryCra07.

import grpc
from concurrent import futures
import pb.hello_pb2 as hpb2
import pb.hello_pb2_grpc as hpb2g
import logging

class Greeter(hpb2g.GreeterServicer):
    def SayHello(self, request, context):
        logging.info("dqwdq")
        logging.info(request)
        return hpb2.HelloResponse(reply="Hello " + request.name)

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    logging.info("Server Start")
    hpb2g.add_GreeterServicer_to_server(Greeter(), server)
    server.add_insecure_port('[::]:8972')
    server.start()
    server.wait_for_termination()

if __name__ == '__main__':
    logging.basicConfig(level=logging.INFO, format='%(asctime)s - %(name)s - %(levelname)s - %(message)s')
    serve()