const sleep = ms => new Promise(resolve => setTimeout(resolve, ms));

const ITERATOR_COUNT = 5;

async function RunServerStream(call) {
    console.log("Incoming request -> ", call.request);
    for (let i = 0; i < ITERATOR_COUNT; i++) {
        console.log(`Iteration ${i + 1}`);
        console.log("processing...");
        await sleep(1500);
        call.write({ code: Math.random() * 10, status: "CONTINUE" });
    }
    console.log("Finished processing");
    call.write({ code: 0, status: "END" });
    call.end();
}

function RunClientStream(call, callback) {
    const arr = [];
    call.on("data", (value) => {
        console.log("Incoming request -> ", value);
        arr.push({ code: value.code, status: value.msg });
    })

    call.on("end", (_) => {
        callback(null, { info: arr })
    })
}

function RunDuplexStream(call) {
    let iterator = 0;

    call.on("data", async (value) => {
        console.log("Incoming request -> ", value);
        if (value.code > 0) {
            iterator++;
            call.write({ code: 1, status: "CONTINUE" });
        }
        if (iterator == 10) {
            call.write({ code: 2, status: "FINISH" });
            call.end();
        }
    })
    
    call.on("end", (_) => {
        call.write({ code: 0, status: "END" });
        call.end();
    })
}

export const EchoStreamImplementation = { RunServerStream, RunClientStream, RunDuplexStream };