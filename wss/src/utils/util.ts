import type { params } from "../types.js"

import dotenv from "dotenv"
dotenv.config()

const WSS_PORT = process.env.WSS_PORT
const SERVER_ADDR = `http://localhost:${WSS_PORT}`

export class ServerUtils {

	static verify_jwt_token(token: string): boolean {
		console.log("token", token)
		return true
	}

	static parse_url(url: string): params {
		const user_params: params = {
			token: "", valid: false
		}
		let cleaned_url = url.replaceAll(" ", "")
		if (!cleaned_url || cleaned_url.length < 5) {
			return user_params
		}

		try {
			const to_url = new URL(SERVER_ADDR + cleaned_url)
			const token = to_url.searchParams.get("token")

			if (!token) {
				return user_params
			}

			user_params.token = token
			user_params.valid = true

			return user_params
		} catch (err) {
			console.log("[error]: ", err)
		}
		return user_params
	}
}



// so the auth_server will generate the 
