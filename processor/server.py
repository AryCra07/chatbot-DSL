import grpc
import time
import argparse
import dsl_engine as gn
from concurrent import futures
from pb import chat_pb2
from pb import chat_pb2_grpc

parser = argparse.ArgumentParser(description='Your description here')
parser.add_argument('--file_path', type=str, default='./test/parser/case4.txt', help='Path to the file')
args = parser.parse_args()

server_name = "py-processor"
m = gn.StateMachine([args.file_path])


class ChatServicer(chat_pb2_grpc.ChatServicer):

    def AnswerService(self, request, context):
        u = gn.UserInfo(request.state, request.name, request.input, request.wallet)
        print('AnswerService -- ' + 'name=' + request.name + ' state=' + str(m.states[u.state]) + ' input=' + u.input)
        r = gn.StateMachine.condition_transform(m, u)

        # Process the request as needed

        # Send a response message back to the client

        response = chat_pb2.ChatResponse(answer=u.answer, state=u.state,
                                         wallet={'balance': u.wallet['balance']})
        print('AnswerService -- ' + 'name=' + request.name + ' state=' + str(u.state) + ' input=' + u.input)
        return response


class GreetServicer(chat_pb2_grpc.GreetServicer):
    def SayHelloService(self, request, context):
        u = gn.UserInfo(request.state, request.name, request.input, request.wallet)
        print('HelloService -- ' + 'name=' + request.name + ' state=' + m.states[u.state])
        r = gn.StateMachine.hello(m, u)
        response = chat_pb2.HelloResponse(words=r)

        return response


class TimeoutServicer(chat_pb2_grpc.TimeoutServicer):
    def TimeoutService(self, request, context):
        u = gn.UserInfo(request.state, request.name, request.input, request.wallet)
        print('HelloService -- ' + 'name=' + request.name + ' state=' + m.states[u.state])

        response = chat_pb2.TimeoutResponse(is_exit=request.input == '1', reset=False, state=request.state, answer=[''],
                                            wallet={'balance': 1000})
        return response


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    chat_pb2_grpc.add_ChatServicer_to_server(ChatServicer(), server)
    chat_pb2_grpc.add_GreetServicer_to_server(GreetServicer(), server)
    chat_pb2_grpc.add_TimeoutServicer_to_server(TimeoutServicer(), server)

    server.add_insecure_port('[::]:50051')
    server.start()
    print("Server started on port 50051")
    try:
        while True:
            time.sleep(3600)
    except KeyboardInterrupt:
        server.stop(0)


if __name__ == '__main__':
    serve()
