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

export type User = {
	name: string
	email: string
	contacts: Contact[]
}



