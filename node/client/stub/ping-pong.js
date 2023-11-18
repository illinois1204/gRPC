import { STUB } from "../index.js";

const sleep = ms => new Promise(resolve => setTimeout(resolve, ms));

export async function playPingPong() {
    const limit = 20;
    let i = 0;

    const session = STUB.PingPong.Play();
    session.on("status", (_) => console.log("Game over"));
    session.on("data", async (_) => {
        console.log("Pong");
        i++;
        if (i == limit) return session.end();

        await sleep(500);
        process.stdout.write("Ping.......");
        await sleep(1000);
        session.write()
    })
    process.stdout.write("Ping.......");
    session.write()
}