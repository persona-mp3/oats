import WebSocket, { WebSocketServer } from "ws"
import express from "express"
import { createServer } from "node:http"

import { handleMessageEvent, createServerMessage } from "./handler.js"
import { MockMongoDB } from "./db/mongo.js"
import { ServerResponse, MessageType } from "./types.js"

import dotenv from "dotenv"

dotenv.config()
import { ServerUtils } from "./utils/util.js"

const WSS_PORT = process.env.WSS_PORT

if (!WSS_PORT) {
	console.error(" env for WSS_PORT has not been set, please set it")
	process.exit()
}

const app = express()
const httpServer = createServer(app)
const wsServer = new WebSocketServer({ noServer: true })


const WELCOME_MESSAGE = " Welcome To Oats"

// enum MessageType { paint }
// type ResponseMessage = {
// 	type: MessageType
// 	body?: any
// }


function firstContentfulPaint(user: string): string {
	const res: ServerResponse = {
		messageType: MessageType.PAINT,
		body: {
			from: "server", dest: user,
			time: Date.now().toLocaleString(),
			message: "paint", code: 200
		}
	}

	const friends = MockMongoDB.findUser(user)
	res.paint = friends
	return JSON.stringify(res)
}

wsServer.on("connection", (socket: WebSocket) => {
	console.log("new client has connected: ", socket.user)

	socket.send(createServerMessage(WELCOME_MESSAGE, 200))
	socket.send(firstContentfulPaint(socket.user))


	socket.on("message", (msg) => {
		handleMessageEvent(msg, socket, wsServer)
	})

	socket.on("close", (code: number, reason: any) => {
		console.log("client closing connection")
		console.log("code: %d, reason: %s", code, reason)
	})

	socket.on("error", (err) => {
		console.log(" error occured with a connected client: ", socket.user)
		console.log(err)
	})
})




httpServer.on("upgrade", (req, socket, head) => {
	const url = req.url
	if (!url) {
		socket.write("400: Forbidden")
		socket.destroy()
		return
	}

	const params = ServerUtils.parseUrl(url)
	if (!params.valid) {
		socket.write("400: Forbidden")
		socket.destroy()
		return
	}

	wsServer.handleUpgrade(req, socket, head, (socket) => {
		socket.token = params.token
		socket.user = params.user
		wsServer.emit("connection", socket)
	})

})

httpServer.listen(WSS_PORT, () => {
	console.log(`
		serve active on @ http://localhost:${WSS_PORT}
	`)
})

