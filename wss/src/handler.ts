import WebSocket, { WebSocketServer } from "ws"

const utf_8_encoding = "utf-8"
type MessageJson = {
	dest: string | "server"
	from: string
	time: string
	message: string
	code: number | 200
}

// so we want to receive a json based-message
// and then send the message to the target
export function handleMessageEvent(payload: WebSocket.RawData, conn: WebSocket, wss: WebSocketServer) {
	console.log(" handling_message")
	console.log("%s", payload.toString())
	try {
		const msg: MessageJson = JSON.parse(payload.toString(utf_8_encoding))

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

		if (!isSent) {
			conn.send(createSendLaterStatus(conn.user))
		}

		conn.send(createSuccessfulSend(conn.user))
		return


	} catch (err) {
		console.log(" could not parse msg")
		console.error(err)
	}

	return
}


function createEmptyDestStatus(user: string): string {
	const msg: MessageJson = {
		dest: user,
		from: "SERVER",
		time: Date.now().toString(),
		code: 400,
		message: "Provide destination to send, Message rejected"
	}

	return JSON.stringify(msg)
}

function createSendLaterStatus(user: string): string {
	const msg: MessageJson = {
		dest: user,
		from: "SERVER",
		time: Date.now().toString(),
		code: 200,
		message: `Send later to ${user}`
	}

	return JSON.stringify(msg)
}

function createSuccessfulSend(user: string): string {
	const msg: MessageJson = {
		dest: user,
		from: "SERVER",
		time: Date.now().toString(),
		code: 200,
		message: `Successful`
	}

	return JSON.stringify(msg)
}
