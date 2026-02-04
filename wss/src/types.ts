import WebSocket from "ws"

export type params = {
	token: string
	valid: boolean
}

export interface OatSocket extends WebSocket {
		token?: string
}


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


