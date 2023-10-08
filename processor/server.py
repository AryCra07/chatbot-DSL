import grpc
import time
from concurrent import futures
from pb import chat_pb2
from pb import chat_pb2_grpc


class AnswerGeneratorServicer(chat_pb2_grpc.AnswerGeneratorServicer):
    def __init__(self):
        self.clients = set()

    def Answer(self, request_iterator, context):
        user_name = "server"  # Set the server's name
        for request in request_iterator:
            user_input = request.userInput
            state = request.userState
            print(f"Received message from {user_name}: {user_input}, State: {state}, Money: {request.userWallet['debt']}")

            # Process the request as needed

            # Send a response message back to the client
            response = chat_pb2.ChatResponse(answer=f"Response for '{request.userInput}'", state=state, userWallet={"balance": 1000})
            yield response

    def send_message_to_clients(self, message):
        for client in self.clients:
            client.on_server_message(message)


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    chat_pb2_grpc.add_AnswerGeneratorServicer_to_server(AnswerGeneratorServicer(), server)
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
