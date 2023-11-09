import grpc
import time
import generator as gn
from concurrent import futures
from pb import chat_pb2
from pb import chat_pb2_grpc

server_name = "py-processor"
m = gn.StateMachine(["./test/parser/case3.txt"])


class ChatServicer(chat_pb2_grpc.ChatServicer):

    def AnswerService(self, request, context):
        user_input = request.input
        state = request.state
        print(
            f"Received message from {server_name}: {user_input}, State: {state}, Money: {request.wallet['debt']}")

        # Process the request as needed

        # Send a response message back to the client

        response = chat_pb2.ChatResponse(answer=[f"Response for '{request.input}'", "Hello"], state=state,
                                         wallet={"balance": 1000})
        return response


class GreetServicer(chat_pb2_grpc.GreetServicer):
    def SayHelloService(self, request, context):
        user_name = request.name
        user_input = request.input
        user_state = request.state

        u = gn.UserInfo(request.state, request.name, request.input, request.wallet)
        print("HelloService -- " + "name=" + request.name)

        r = gn.StateMachine.hello(m, u)

        # 处理客户端的请求，这里简单示范将用户的输入添加到回应中
        response = chat_pb2.HelloResponse(words=r)

        return response


class TimeoutServicer(chat_pb2_grpc.TimeoutServicer):
    def TimeoutService(self, request, context):
        response = chat_pb2.TimeoutResponse(is_exit=request.input=='1', reset=False, state=request.state, answer=[''],
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
