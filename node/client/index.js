import grpc from "@grpc/grpc-js"
import protoLoader from "@grpc/proto-loader"
import { echoSample } from "./stub/echo.js";
import { clientStream, duplexStream, serverStream } from "./stub/stream-echo.js";
import { playPingPong } from "./stub/ping-pong.js";

const PORT = 8717;
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
    PingPong: PACKAGE_NAME_SPACE.pingPong
};

export const STUB = {
    Echo: new CONTRACTS.Echo.Echo(`127.0.0.1:${PORT}`, grpc.credentials.createInsecure()),
    Calculate: new CONTRACTS.Echo["Calculate"](`127.0.0.1:${PORT}`, grpc.credentials.createInsecure()),
    StreamEcho: new CONTRACTS.Stream.StreamEcho(`127.0.0.1:${PORT}`, grpc.credentials.createInsecure()),
    PingPong: new CONTRACTS.PingPong.PingPong(`127.0.0.1:${PORT}`, grpc.credentials.createInsecure())
}

// echoSample();
// serverStream();
// clientStream();
// duplexStream();
playPingPong()