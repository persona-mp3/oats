import { ServerResponse, MessageType, Message } from ".././types.js"
import { MockMongoDB } from "../db/mongo.js"
class Protocols {
	/**
	 * Makes request to the Redis cache and gets all the friends
	 * this user has. This is only sent upon first connection and never
	 * again
	 * @param user 
	 * @returns 
	 */
	getFirstContentfulPaint(user: string): string {
		const res: ServerResponse = {
			messageType: MessageType.PAINT,
			body: {
				from: "server", dest: user,
				time: Date.now().toLocaleString(),
				message: "paint", code: 200
			}
		}

		const friends = MockMongoDB.findUser(user)
		res.paint = friends
		return JSON.stringify(res)
	}

	/**
	 * A JSON string is returned after wrapped inside a ServerResponse type
	 * Paint messages must not be sent here
	 * @param msg 
	 * @param code 
	 * @returns 
	 */
	customServerMessage(msg: string, code: number): string {
		const resBody: Message = {
			dest: "you",
			from: "SERVER",
			time: Date.now().toString(),
			code: code,
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
			return this.internalServerErrorMessage()
		}

	}

	/**
	 * Sends status code of 500, and "Internal Server Error" message
	 * @returns 
	 */
	internalServerErrorMessage(): string {
		return this.customServerMessage("Internal Server Error", 500)
	}

}

export default Protocols
