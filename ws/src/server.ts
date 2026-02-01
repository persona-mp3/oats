import express from "express"
import { createServer } from "node:http"
import WebSocket, { WebSocketServer } from "ws"
import type { IncomingMessage } from "node:http"

import { OatSocket } from "./extensions.js"
import { authClient } from "./middleware/auth.js"
import { handleError, handleDisconnection, handleMessage } from "./handlers.js"

const app = express()
const httpServer = createServer(app)
const wss = new WebSocketServer({ noServer: true })

const PORT = 3900

app.get("/", (_req, res) => {
	console.log("index route of server got hit")
	res.send("took_a_pill_in_ibixa")
})

function handleNewConnection(ws: WebSocket, req: IncomingMessage) {
	console.log("[nu_conn] new connection")

	ws.send("[host]: welcome to OATS")

	ws.on("close", handleDisconnection)

	ws.on("message", (data) => handleMessage(ws, data))

	ws.on("error", handleError)
}

wss.on("connection", handleNewConnection)

httpServer.on("upgrade", (req, socket, head) => {
	console.log("listening for upgrade")
	if (!req.url) {
		console.log("ban him chat")
		socket.destroy()
		return
	}

	const url = new URL("http://localhost:3900" + req.url)
	const user = url.searchParams.get("userName")
	const token = url.searchParams.get("token")

	if (!user || !token) {
		console.log("client did not provide username or token", user, token)
		socket.destroy()
		return
	}

	const isAuth = authClient(user, token)
	console.log("clients auth status:", isAuth)

	if (!isAuth.status) {
		console.log("clients socket will be destroyed")
		socket.destroy()
		return
	}

	wss.handleUpgrade(req, socket, head, (ws: OatSocket) => {
		ws.user = user // userName provided
		ws.token = token // auth token parsed from url_params
		ws.info = isAuth.info // includes friends, groups, no exposed credentials
		wss.emit("connection", ws, req)
	})
})

httpServer.listen(PORT, () => {
	console.log("web-socket-server: http://localhost:3900")
})
