from concurrent import futures
import datetime
import grpc
import echo_pb2
import echo_pb2_grpc

class Echo(echo_pb2_grpc.EchoServicer):
    def call(self, request, context):
        return echo_pb2.echoResponse(message="yaos, %s, this response from server. Time is %s" % (request.name, datetime.datetime.now()))

class Calculate(echo_pb2_grpc.CalculateServicer):
    def sum(self, request, context):
        return echo_pb2.calculateResponse(result=request.arg1 + request.arg2)

    def min(self, request, context):
        arg1 = request.arg1
        arg2 = request.arg2
        return echo_pb2.calculateResponse(result=arg2 if arg1 > arg2 else arg1)

def serve():
    port = "8717"
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    
    echo_pb2_grpc.add_EchoServicer_to_server(Echo(), server)
    echo_pb2_grpc.add_CalculateServicer_to_server(Calculate(), server)

    server.add_insecure_port("[::]:" + port)
    server.start()
    print("Server started, listening on " + port)
    server.wait_for_termination()

serve()