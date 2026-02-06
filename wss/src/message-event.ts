import type { Message } from "./types.js"
import WebSocket, { WebSocketServer } from "ws"
import { handleMessageEventNew } from "./handler.js"


// now if we wanted the user the user 
// to search for friends how would it be done?
// first we can have a user pointing to a collection of their friends in the db
// and then when we recv a messageType, persay, "find-friends", we just show 
// them the list of friends they already have
// master -> <ghostFriend1, ghostFriend2, ghostFriend3>
// type ?: boolean
// Due to how this project is far more than I predicted 😭✌️ 
// I'll have to use protobufs because of these types
// Because the protocol is slowly building

// hacking it now, for a find event, we can get the clients to send
//
// f-userName to find <userName>
// or f-all to get  alist of all users 
// but this should also be done on initial connection instead
// so all they have to do is to find a particalar user
// obviously we won't send all 1000 of them but lets say first 5/10 for now

type ReqMessage = {
	body: Message
	type: string
}

const EVENT_FIND = "f"
const EVENT_CHAT = "i"

function handleFindEvent() { }

export function eventManager(payload: WebSocket.RawData, conn: WebSocket, wsServer: WebSocketServer) {
	try {
		const request: ReqMessage = JSON.parse(payload.toString())


		switch (request.type) {
			case EVENT_CHAT:
				console.log(" normal chat event")
				handleMessageEventNew(request.body, conn, wsServer)
				break

			case EVENT_FIND:
				console.log(" get users friends and send back to them")
				handleFindEvent()
				break
			default:
				console.log("unidentified event")

		}


	} catch (err: Error | any) {
		if (err.name == "SyntaxError") {
			console.log(" couldn't parse request send bad request to websocket")
			return
		}

		console.log(" unexpected error:")
		console.log(err)
	}
}
