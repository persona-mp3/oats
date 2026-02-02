export function handleDisconnection(code, reason) {
    console.log("client disconnecting:", code, reason.toString());
}
export function handleMessage(ws, data) {
    console.log("\n\n");
    console.log("[recv] ", data.toString());
    console.log("\n\n");
}
export function handleError(err) {
    console.error("received error from client");
    console.error(err);
}
