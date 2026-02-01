import WebSocket, { WebSocketServer } from "ws"

export function handleDisconnection(code: number, reason: Buffer) {
	console.log("client disconnecting:", code, reason.toString())
}

export function handleMessage(ws: WebSocket, data: WebSocket.RawData) {
	console.log("\n\n")
	console.log("[recv] ", data.toString())
	console.log("\n\n")
}

export function handleError(err: Error) {
	console.error("received error from client")
	console.error(err)
}
