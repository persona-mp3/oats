import { WebSocketServer } from "ws"
import express from "express"
import { createServer } from "node:http"

import { ServerUtils } from "./utils/util.js"
import Protocols from "./protocols/main.js"
import Handlers from "./handlers/main.js"


import dotenv from "dotenv"
dotenv.config()

const WSS_PORT = process.env.WSS_PORT
if (!WSS_PORT) {
	console.error(" env for WSS_PORT has not been set, please set it")
	process.exit()
}

const app = express()
const httpServer = createServer(app)
const wsServer = new WebSocketServer({ noServer: true })

const WELCOME_MESSAGE = " \n\nWelcome To Oats\n\n"

const protocols = new Protocols()
const handlers = new Handlers()

wsServer.on("connection", (socket) => {
	console.log(" new client connected:", socket.user)

	socket.send(protocols.customServerMessage(WELCOME_MESSAGE, 200))
	socket.send(protocols.getFirstContentfulPaint(socket.user))

	socket.on("message", (payload) => handlers.messageEvent(payload, socket, wsServer))

	socket.on("close", (code: number, reason: any) => {
		console.log(" client closed connection")
		console.log(`code: ${code}, reason: ${reason}`)
	})

	socket.on("error", (err) => {
		console.error(" an error occured from ", socket.user)
		console.error(err)
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
			server running on http://localhost:${WSS_PORT}
	`)
})

