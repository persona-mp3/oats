import WebSocket, { WebSocketServer } from "ws"
import { Message } from "../types.js"
import Protocols from "../protocols/main.js"

const protocol = new Protocols()

class Handlers {
	messageEvent(payload: WebSocket.RawData, conn: WebSocket, server: WebSocketServer): void {
		try {
			const request: Message = JSON.parse(payload.toString())
			console.log(`*new_msg from:${request.from},to: ${request.dest}`)

			if (request.dest.replaceAll(" ", "").length === 0) {
				conn.send(protocol.customServerMessage("Invalid destination", 400))
				return
			}

			// check if the message type is actually a chat message
			let msgSent = false
			server.clients.forEach((client) => {
				if (client !== conn
					&& client.user === request.dest
					&& client.readyState === WebSocket.OPEN) {
					console.log(" found destination", client.user)
					msgSent = true

					client.send(JSON.stringify(request))
				}
			})

			if (!msgSent) {
				// send to a message broker
				console.log("couldn't find dest", request.dest)
			}
		} catch (err) {
			console.error(" an error occured")
			console.error(err)
			conn.send(protocol.internalServerErrorMessage())
		}
	}


}

export default Handlers

