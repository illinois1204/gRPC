import grpc from "@grpc/grpc-js"
import protoLoader from "@grpc/proto-loader"

const PROTO_PATH = "../../proto/echo.proto"
const protoLoaderOptions = { 
    KeepCase : true, 
    longs : String, 
    enums : String, 
    defaults : true, 
    oneofs : true 
}

const packageDefinition = protoLoader.loadSync(PROTO_PATH, protoLoaderOptions);
const EchoContract = grpc.loadPackageDefinition(packageDefinition).echo;

const STUB = {
    Echo: new EchoContract.Echo("127.0.0.1:8718", grpc.credentials.createInsecure()),
    Calculate: new EchoContract["Calculate"]("127.0.0.1:8717", grpc.credentials.createInsecure())
}

STUB.Echo.call({ name: "illinois" }, (err, res) => {
    if (err) {
        console.log(err);
    }
    else {
        console.log(res);
    }
})

STUB.Calculate.sum({ arg1: 50, arg2: 25 }, (err, res) => {
    if (err) {
        console.log(err);
    }
    else {
        console.log(res);
    }
})

STUB.Calculate.min({ arg1: 35, arg2: 22 }, (err, res) => {
    if (err) {
        console.log(err);
    }
    else {
        console.log(res);
    }
})