import WebSocket, { WebSocketServer } from "ws"

export function handleDisconnection(code: number, reason: Buffer) {
	console.log("disconnection handler")
}

export function handleMessage(ws: WebSocket, data: WebSocket.RawData) {
	console.log("message handler")
}

export function handleError(err: Error) {
	console.error("received error from client")
	console.error(err)
}
