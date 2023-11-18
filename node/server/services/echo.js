function call(call, callback) {
    const req = call.request;
    const res = { message: `yaos, ${req.name}, this response from server. Time is ${Date.now()}` };
    callback(null, res);
}

export const EchoImplementation = { call }