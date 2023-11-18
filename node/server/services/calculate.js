function sum(call, callback) {
    const { arg1, arg2 } = call.request;
    const res = { result: arg1 + arg2 };
    callback(null, res);
}

function min(call, callback) {
    const { arg1, arg2 } = call.request;
    const result = arg1 > arg2 ? arg2 : arg1;
    callback(null, { result });
}

export const CalculateImplementation = { sum, min }