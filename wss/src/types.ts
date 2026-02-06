import WebSocket from "ws"


export type params = {
	user: string
	token: string
	valid: boolean
}

// export interface OatSocket extends WebSocket {
// 	token?: string
// 	user?: string
// }
//

export type Contact = {
	name: string
	active: boolean
	lastSeen: string
}

export type UserDB = {
	name: string
	email: string
	contacts: Contact[]
}


export type Message = {
	dest: string | "server"
	from: string
	time: string
	message: string
	code: number | 200

}

export type Friend = {
	name: string
	lastSeen: string
}

export enum MessageType { PAINT, CHAT }
export type ServerResponse = {
	messageType: MessageType
	body: Message
	paint?: Friend[]
}


