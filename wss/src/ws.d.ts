import "ws"

declare module "ws" {
	interface WebSocket {
		token: string
		user: string
		valid: boolean
	}
}
