import grpc
import time
import argparse
import dsl_engine as gn
from concurrent import futures
from pb import chat_pb2
from pb import chat_pb2_grpc

parser = argparse.ArgumentParser(description='parse files path')
parser.add_argument('--file_path', type=str, default='./script/script1.txt', help='Path to the file')
args = parser.parse_args()

server_name = "py-processor"
m = gn.StateMachine([args.file_path])


class GreetServicer(chat_pb2_grpc.GreetServicer):
    def SayHelloService(self, request, context):
        u = gn.UserInfo(request.state, request.name, "", {'balance': request.balance, 'bill': request.bill})
        print('HelloService -- ' + 'name=' + request.name + ' state=' + m.states[u.state])
        r = gn.StateMachine.hello(m, u)
        response = chat_pb2.HelloResponse(words=r)

        return response


class ChatServicer(chat_pb2_grpc.ChatServicer):

    def AnswerService(self, request, context):
        u = gn.UserInfo(request.state, request.name, request.input, {'balance': request.balance, 'bill': request.bill})
        print('AnswerService -- ' + 'name=' + request.name + ' state=' + str(m.states[u.state]) + ' input=' + u.input)
        r = gn.StateMachine.condition_transform(m, u)

        # Process the request as needed

        # Send a response message back to the client

        response = chat_pb2.ChatResponse(answer=u.answer, state=u.state,
                                         wallet={'balance': u.wallet['balance']})
        print('AnswerService -- ' + 'name=' + request.name + ' state=' + str(u.state) + ' input=' + u.input)
        return response


class TimeoutServicer(chat_pb2_grpc.TimerServicer):
    def TimerService(self, request, context):
        u = gn.UserInfo(request.state, request.name, "", {})
        now_time = request.now_time
        print('TimeoutService -- ' + 'name=' + request.name + ' state=' + m.states[u.state])
        answer, is_exit, reset  = gn.StateMachine.timeout_transform(m, u, now_time)

        response = chat_pb2.TimerResponse(is_exit=is_exit, reset=reset, state=u.state, answer=answer)
        return response


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    chat_pb2_grpc.add_ChatServicer_to_server(ChatServicer(), server)
    chat_pb2_grpc.add_GreetServicer_to_server(GreetServicer(), server)
    chat_pb2_grpc.add_TimerServicer_to_server(TimeoutServicer(), server)

    server.add_insecure_port('[::]:50051')
    server.start()
    print('Server ' + server_name + ' started on port 50051')
    try:
        while True:
            time.sleep(3600)
    except KeyboardInterrupt:
        server.stop(0)


if __name__ == '__main__':
    serve()
