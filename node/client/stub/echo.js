import { STUB } from "../index.js";

export function echoSample() {
    STUB.Echo.call({ name: "illinois" }, (err, res) => {
        if (err) console.log(err);
        else console.log(res);
    })

    STUB.Calculate.sum({ arg1: 50, arg2: 25 }, (err, res) => {
        if (err) console.log(err);
        else console.log(res);
    })

    STUB.Calculate.min({ arg1: 35, arg2: 22 }, (err, res) => {
        if (err) console.log(err);
        else console.log(res);
    })
}

