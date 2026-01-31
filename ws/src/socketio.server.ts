import express from "express"
import { createServer } from "node:http"
import { Server } from "socket.io"
import { fileURLToPath } from "node:url"
import { dirname, join } from "node:path"


const WS_PORT = 3900

const app = express()
const server = createServer(app)
const wsServer = new Server(server, { connectionStateRecovery: {} })

const __dirname = dirname(fileURLToPath(import.meta.url))

// main middleware for validating all requests
function authenticateUser(): void {
	console.log("base_ball")
}
authenticateUser()

app.get("/", (req, res) => {
	console.log("someone pingedhere", req.originalUrl)
	res.sendFile(join(__dirname, "index.html"))
})


// WEBSOCKET OPERATIONS
function handleNewConnection(socket: any) {
	console.log("nu_client has connected")

	socket.on("disconnect", handleDisconnection)
	socket.on("message", handleMessageEvent)
}

function handleMessageEvent(msg: any) {
	wsServer.emit("message", msg)
}


function handleDisconnection() {
	console.log("nu_client disconnected")
}

wsServer.on("connection", handleNewConnection)

// wsServer.on("connection", (socket) => {
// 	console.log("new connection made")
// })


server.listen(WS_PORT, () => {
	console.log("wss running on: http://localhost:3900")
})
