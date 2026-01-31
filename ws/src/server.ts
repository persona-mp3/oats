import express from "express"
import { createServer } from "node:http"
import WebSocket, { WebSocketServer } from "ws"

import type { IncomingMessage } from "node:http"
import { OatSocket } from "./extensions.js"

import { authClient } from "./middleware/auth.js"
import { handleError, handleDisconnection, handleMessage } from "./handlers.js"

const app = express()
const httpServer = createServer(app)
// const wss = new WebSocketServer({ server: httpServer })
const wss = new WebSocketServer({ noServer: true })

const PORT = 3900

app.get("/", (req, res) => {
	console.log("index route of server got hit")
	res.send("took_a_pill_in_ibixa")
})


function handleNewConnection(ws: WebSocket, req: IncomingMessage) {
	ws.on("close", handleDisconnection)

	ws.on("message", (data) => handleMessage(ws, data))

	ws.on("error", handleError)
}


wss.on("connection", handleNewConnection)

httpServer.on("upgrade", (req, socket, head) => {
	if (!req.url) {
		console.log("ban him chat")
		socket.destroy()
		return
	}

	const url = new URL(req.url)
	const user = url.searchParams.get("user")
	const token = url.searchParams.get("token")

	if (!user || !token) {
		console.log("client did not provide username or token", user, token)
		socket.destroy()
		return
	}

	const isAuth = authClient(user, token)

	if (!isAuth.status) {
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
