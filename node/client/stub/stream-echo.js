import { STUB } from "../index.js";

const sleep = ms => new Promise(resolve => setTimeout(resolve, ms));

export async function serverStream() {
    STUB.StreamEcho.RunServerStream({ code: 101, msg: "hello grpc" })
    .on("data", (value) => {
        console.log("Receive response from server");
        console.log(value);
    })
    .on("status", (status) => {
        console.log("Stream is canceled");
        console.log(status);
    })
}

export async function clientStream() {
    const clientStream = STUB.StreamEcho.RunClientStream((err, res) => {
        if (err) console.log(err);
        else console.log(res);
    });
    for (let i = 0; i < 5; i++) {
        clientStream.write({ code: Math.random() * 10, msg: `Random text ${i}` });
        await sleep(1500);
    }
    clientStream.end();
}

export async function duplexStream() {
    let stop = false;
    const duplexStream = STUB.StreamEcho.RunDuplexStream();
    duplexStream.on("status", (status) => {
        stop = true;
        console.log("Stream is canceled");
        console.log(status);
    })
    duplexStream.on("data", (value) => {
        console.log("Receive response from server");
        console.log(value);
    })

    let counter = 0;
    while (true) {
        if (stop) break;
        if (counter == 15) {
            duplexStream.end();
            break;
        }

        counter++;
        duplexStream.write({ code: Math.random() * 50 - 25, msg: "lorem ipsum..." })
        await sleep(500);
    }
}
