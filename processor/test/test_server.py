import unittest
import grpc
from concurrent import futures
from stub import serve
import pb.chat_pb2 as chat_pb2
import pb.chat_pb2_grpc as chat_pb2_grpc
class TestServer(unittest.TestCase):

    @classmethod
    def setUpClass(cls):
        # 在一个单独的线程中启动gRPC服务器
        cls.server_thread = futures.ThreadPoolExecutor(max_workers=1)
        cls.server_future = cls.server_thread.submit(serve)

    @classmethod
    def tearDownClass(cls):
        # 在测试后停止gRPC服务器
        cls.server_future.result().stop(0)

    def test_hello(self):
        with grpc.insecure_channel('localhost:50052') as channel:
            stub = chat_pb2_grpc.GreetStub(channel)
            request = chat_pb2.HelloRequest(state=1, name="John")
            response = stub.SayHelloService(request)
            self.assertEqual(response.words, ['Hello'])

    def test_chat(self):
        with grpc.insecure_channel('localhost:50052') as channel:
            stub = chat_pb2_grpc.ChatStub(channel)
            request = chat_pb2.ChatRequest(state=1, name="Alice", input="Hello", wallet={"USD": 100.0, "EUR": 50.0})
            response = stub.AnswerService(request)
            self.assertEqual(response.state, 1)
            self.assertEqual(response.answer, ['Hello'])
            self.assertEqual(response.wallet["USD"], 100.0)
            self.assertEqual(response.wallet["EUR"], 50.0)

    def test_timer(self):
        with grpc.insecure_channel('localhost:50052') as channel:
            stub = chat_pb2_grpc.TimerStub(channel)
            request = chat_pb2.TimerRequest(state=1, last_time=0, now_time=5)
            response = stub.TimerService(request)
            self.assertEqual(response.state, 1)
            self.assertEqual(response.answer, ['Timer'])
            self.assertFalse(response.is_exit)
            self.assertFalse(response.reset)


if __name__ == '__main__':
    unittest.main()
