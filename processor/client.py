import grpc
import time
from pb import chat_pb2
from pb import chat_pb2_grpc

class ChatClient:
    def __init__(self, name):
        self.name = name
        self.channel = grpc.insecure_channel('localhost:50051')
        self.stub = chat_pb2_grpc.AnswerGeneratorStub(self.channel)

    def chat(self):
        user_state = 0
        while True:
            user_input = input("Enter your message: ")
            user_state += 1
            request = chat_pb2.ChatRequest(userName=self.name, userInput=user_input, userState=user_state, userWallet={"balance": 1000, "debt": 0})
            response_iterator = self.stub.Answer(iter([request]))
            for response in response_iterator:
                print(f"Received message from server: {response.answer}, State: {response.state}, Money: {response.userWallet['balance']}")

    def on_server_message(self, message):
        print(f"Received message from server: {message}")

if __name__ == '__main__':
    name = "client"  # Set the client's name
    client = ChatClient(name)
    client.chat()
