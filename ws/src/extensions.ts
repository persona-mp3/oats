import WebSocket from "ws"

export interface OatSocket extends WebSocket {
	user?: string
	token?: string
	info?: any
}
