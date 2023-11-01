import grpc

import echo_pb2
import echo_pb2_grpc

STUB = {}

def run():
    with grpc.insecure_channel("127.0.0.1:8718") as channel:
        STUB["echo"] = echo_pb2_grpc.EchoStub(channel)
        STUB["calculate"] = echo_pb2_grpc.CalculateStub(channel)

        response = STUB["echo"].call(echo_pb2.echoRequest(name="SUI"))
        print(response.message)

        response = STUB["calculate"].sum(echo_pb2.calculateRequest(arg1=5, arg2=15))
        print()
        print(response)

        response = STUB["calculate"].min(echo_pb2.calculateRequest(arg1=7546, arg2=246186))
        print()
        print(response)

run()