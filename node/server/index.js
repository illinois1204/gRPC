import grpc from "@grpc/grpc-js"
import protoLoader from "@grpc/proto-loader"
import { EchoMethods } from "./services/echo.js"
import { CalculateMethods } from "./services/calculate.js"

const PORT = 8717
const PROTO_PATH = "../../proto/echo.proto"
const protoLoaderOptions = { 
    KeepCase : true, 
    longs : String, 
    enums : String, 
    defaults : true, 
    oneofs : true 
}

const packageDefinition = protoLoader.loadSync([PROTO_PATH], protoLoaderOptions);
const PACKAGE_NAME_SPACE = grpc.loadPackageDefinition(packageDefinition);
const CONTRACTS = {
    echo: PACKAGE_NAME_SPACE.echo
};

const server = new grpc.Server();
server.addService(CONTRACTS.echo.Echo.service, EchoMethods);
server.addService(CONTRACTS.echo.Calculate.service, CalculateMethods);
server.bindAsync(`0.0.0.0:${PORT}`, grpc.ServerCredentials.createInsecure(), () => {
    server.start();
    console.log(`Server is running on port: ${PORT}`);
});
