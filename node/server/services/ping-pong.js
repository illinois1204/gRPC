function Play(call) {
    call.on("data", async (_) => {
        console.log("accept ping");
        call.write();
    })
    
    call.on("end", (_) => call.end())
}

export const PingPongImplementation = { Play };