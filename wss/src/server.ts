import express from "express"
import { createServer } from "node:http"
import WebSocket, { WebSocketServer } from "ws"
import dotenv from "dotenv"
import { ServerUtils as s_utils } from "./utils/util.js"

import type { OatSocket } from "./types.js"

dotenv.config()
const WSS_PORT = process.env.WSS_PORT
// const SERVER_ADDR = `http://localhost:${WSS_PORT}`
if (!WSS_PORT) {
	console.error("port for wsserver hasn't been set")
	process.exit()
}

const app = express()
const http_server = createServer(app)
const ws_server = new WebSocketServer({ noServer: true })

app.get("/welcome", (_req, res) => {
	console.log("[/welcome]: visited")
	res.send("hello there")
})


// ==== handlers =====
function handle_message(oat_socket: OatSocket, msg: any, broadcast?: Function) { }
// ====§====§====§====

ws_server.on("connection", (oat_socket) => {

	oat_socket.on("message", (msg: WebSocket.RawData) => {

		ws_server.clients.forEach((client) => {
			if (client !== oat_socket || client.readyState === WebSocket.OPEN) {
				client.send(msg)
			}
		})

	})


	oat_socket.send(" welcome to the plate of oats... ")

	// oat_socket.on("close", (evt) => {
	// })
	//
	// oat_socket.on("error", (err) => {
	// })
})


http_server.on("upgrade", (req, socket, head) => {
	console.log("[event]: handling upgrade request")

	const redirect_url = req.url
	if (!redirect_url) {
		socket.write("400: provide redirect-url")
		socket.destroy()
		return
	}

	let non_shared_buffer = head.toString()
	console.log("what was passed in as head", non_shared_buffer)

	const params = s_utils.parse_url(redirect_url)
	if (!params.valid) {
		socket.write("403: invalid credentials")
		socket.destroy()
		return
	}

	const valid_token = s_utils.verify_jwt_token(params.token)
	if (!valid_token) {
		socket.write("403: invalid token")
		socket.destroy()
		return
	}

	ws_server.handleUpgrade(req, socket, head, (oat_socket: OatSocket) => {
		// check cache and also attach user_info here
		oat_socket.token = params.token
		ws_server.emit("connection", oat_socket)
	})

})


http_server.listen(WSS_PORT, () => {
	console.log(`
		http-ws_server running on http://localhost:${WSS_PORT}
		ws_protocol on ws://localhost:${WSS_PORT}
	`)
})
