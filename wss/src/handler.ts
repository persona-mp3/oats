import WebSocket, { WebSocketServer } from "ws"
import { Message, ServerResponse, MessageType } from "./types.js"

const utf_8_encoding = "utf-8"


export function handleMessageEventNew(payload: Message, conn: WebSocket, wss: WebSocketServer) { }


// so we want to receive a json based-message
// and then send the message to the target
export function handleMessageEvent(payload: WebSocket.RawData, conn: WebSocket, wss: WebSocketServer) {
	console.log(" handling_message")
	console.log("%s", payload.toString())
	try {
		const msg: Message = JSON.parse(payload.toString(utf_8_encoding))

		if (msg.dest.replaceAll(" ", "") == "") {
			console.log(" no valid dest provided, echo back to client or prompt to provide destpayload")
			conn.send(createEmptyDestStatus(conn.user))
		}

		const allClients = wss.clients
		let isSent = false
		allClients.forEach((client) => {
			if (client !== conn && client.user === msg.dest && client.readyState === WebSocket.OPEN) {
				console.log(" found client", client.token, client.user)
				isSent = true
				client.send(payload)
			}
		})


		const res = createServerMessage(" a ninja bike would be nice")
		conn.send(res)

		// we'd want to cache this or put inside kafka or rabbitMQ, a message broker
		if (!isSent) {
			console.log(" could not find dest user: %s for %s:", msg.dest, conn.user)
			return
		}
		// if (!isSent) {
		// 	conn.send(createSendLaterStatus(conn.user))
		// } else {
		// 	return
		// }
		//
		return


	} catch (err) {
		console.log(" could not parse msg")
		console.error(err)
	}

	return
}

const ERR_INTERNAL_SERVER_ERR_MSG = "Internal Server Error"

function internalServerErrorMessage(): string {
	const msg: Message = {
		dest: "you",
		from: "SERVER",
		time: Date.now().toString(),
		code: 500,
		message: ERR_INTERNAL_SERVER_ERR_MSG,
	}

	return JSON.stringify(msg)
}
export function createServerMessage(msg: string, code?: number): string {
	const resBody: Message = {
		dest: "you",
		from: "SERVER",
		time: Date.now().toString(),
		code: 200,
		message: msg
	}

	const response: ServerResponse = {
		messageType: MessageType.CHAT,
		body: resBody,
		paint: []
	}

	try {
		return JSON.stringify(response)
	} catch (err) {
		console.log(" parsing error:", err)
		return internalServerErrorMessage()
	}
}









function createEmptyDestStatus(user: string): string {
	const msg: Message = {
		dest: user,
		from: "SERVER",
		time: Date.now().toString(),
		code: 400,
		message: "Provide destination to send, Message rejected"
	}

	const response: ServerResponse = {
		messageType: MessageType.CHAT,
		body: msg,
		paint: []
	}

	return JSON.stringify(msg)
}

function createSendLaterStatus(user: string): string {
	const msg: Message = {
		dest: user,
		from: "SERVER",
		time: Date.now().toString(),
		code: 200,
		message: `Send later to ${user}`
	}
	const response: ServerResponse = {
		messageType: MessageType.CHAT,
		body: msg,
		paint: []
	}

	return JSON.stringify(msg)
}

function createSuccessfulSend(user: string): string {
	const msg: Message = {
		dest: user,
		from: "SERVER",
		time: Date.now().toString(),
		code: 200,
		message: `Successful`
	}

	const response: ServerResponse = {
		messageType: MessageType.CHAT,
		body: msg,
		paint: []
	}
	return JSON.stringify(msg)
}
