import grpc
import time
from pb import chat_pb2
from pb import chat_pb2_grpc


class ChatClient:
    def __init__(self, name):
        self.name = name
        self.channel = grpc.insecure_channel('localhost:50051')
        self.stub1 = chat_pb2_grpc.ChatStub(self.channel)
        self.stub2 = chat_pb2_grpc.GreetStub(self.channel)
        self.stub3 = chat_pb2_grpc.TimeoutStub(self.channel)

    def chat(self):
        user_state = 0
        while True:
            user_input = input("Enter your message: ")
            user_state += 1
            request = chat_pb2.UserRequest(name=self.name, input=user_input, state=user_state,
                                           wallet={"balance": 1000, "debt": 0})
            response = self.stub1.AnswerService(request)
            print(
                f"Received message from server: {response.answer}, State: {response.state}, Money: {response.wallet['balance']}")


    def hello(self):
        while True:
            user_input = input("Enter your message: ")
            user_request = chat_pb2.UserRequest(
                state=0,
                name="YourName",
                input=user_input,
                wallet={"balance": 1000, "debt": 0}
            )

            response = self.stub2.SayHelloService(user_request)

            print(f"Received message from server: {response.words}")

    def timeout(self):
        while True:
            user_input = input("Enter your message: ")
            user_request = chat_pb2.UserRequest(
                state=0,
                name="YourName",
                input=user_input,
                wallet={"balance": 1000, "debt": 0}
            )

            response = self.stub3.TimeoutService(user_request)

            print(f"Received message from server: {response.is_exit}")


if __name__ == '__main__':
    name = "client"  # Set the client's name
    client = ChatClient(name)
    client.timeout()
