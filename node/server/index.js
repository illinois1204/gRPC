import grpc from "@grpc/grpc-js"
import protoLoader from "@grpc/proto-loader"
import { EchoImplementation } from "./services/echo.js"
import { CalculateImplementation } from "./services/calculate.js"
import { EchoStreamImplementation } from "./services/echo-stream.js"
import { PingPongImplementation } from "./services/ping-pong.js"

const PORT = 8717
const PROTO_PATH = [
    "../../proto/echo.proto",
    "../../proto/stream.proto",
    "../../proto/ping-pong.proto"
]
const protoLoaderOptions = { 
    KeepCase : true, 
    longs : String, 
    enums : String, 
    defaults : true, 
    oneofs : true 
}

const packageDefinition = protoLoader.loadSync(PROTO_PATH, protoLoaderOptions);
const PACKAGE_NAME_SPACE = grpc.loadPackageDefinition(packageDefinition);
const CONTRACTS = {
    Echo: PACKAGE_NAME_SPACE.echo,
    Stream: PACKAGE_NAME_SPACE.stream,
    PingPong: PACKAGE_NAME_SPACE.pingPong,
};

const server = new grpc.Server();
server.addService(CONTRACTS.Echo.Echo.service, EchoImplementation);
server.addService(CONTRACTS.Echo.Calculate.service, CalculateImplementation);
server.addService(CONTRACTS.Stream.StreamEcho.service, EchoStreamImplementation);
server.addService(CONTRACTS.PingPong.PingPong.service, PingPongImplementation);

server.bindAsync(`0.0.0.0:${PORT}`, grpc.ServerCredentials.createInsecure(), () => {
    server.start();
    console.log(`Server is running on port: ${PORT}`);
});
