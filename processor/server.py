import grpc
import time
import argparse
from dsl import dsl_engine as gn
from concurrent import futures
from pb import chat_pb2
from pb import chat_pb2_grpc

parser = argparse.ArgumentParser(description='parse files path')
parser.add_argument('--fp', type=str, default='./script/script3.txt', help='Path to the files')
args = parser.parse_args()

server_name = "py-processor"
m = gn.StateMachine([args.fp])


class GreetServicer(chat_pb2_grpc.GreetServicer):
    def SayHelloService(self, request, context):
        u = gn.UserInfo(request.state, request.name, "", {})
        print('HelloService -- ' + 'name=' + request.name + ' state=' + m.states[u.state])
        r = gn.StateMachine.hello(m, u)
        response = chat_pb2.HelloResponse(words=r)

        return response


class ChatServicer(chat_pb2_grpc.ChatServicer):

    def AnswerService(self, request, context):
        u = gn.UserInfo(request.state, request.name, request.input, request.wallet)
        print('AnswerService -- ' + 'name=' + request.name + ' state=' + str(m.states[u.state]) + ' input=' + u.input)

        # Process the request as needed
        m.condition_transform(u)

        # Send a response message back to the client
        response = chat_pb2.ChatResponse(state=u.state, answer=u.answer, wallet=u.wallet)
        return response


class TimerServicer(chat_pb2_grpc.TimerServicer):
    def TimerService(self, request, context):
        u = gn.UserInfo(request.state, "", "")
        last_time = request.last_time
        now_time = request.now_time
        print('TimerService -- ' + ' state=' + str(m.states[u.state]) + ' last_time=' + str(last_time) + ' now_time=' + str(now_time))
        r, is_exit, reset = gn.StateMachine.timeout_transform(m, u, last_time, now_time)

        response = chat_pb2.TimerResponse(is_exit=is_exit, reset=reset, state=u.state, answer=r.answer)
        return response


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    chat_pb2_grpc.add_ChatServicer_to_server(ChatServicer(), server)
    chat_pb2_grpc.add_GreetServicer_to_server(GreetServicer(), server)
    chat_pb2_grpc.add_TimerServicer_to_server(TimerServicer(), server)

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
