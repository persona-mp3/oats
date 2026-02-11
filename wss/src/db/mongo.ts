
import type { Friend } from "../types.js"
export class MockMongoDB {
	constructor() {
	}

	public static findUser(user: string): Friend[] {
		const mockDb = new Map()
		const ladiesMan217User: Friend = {
			name: "ladiesMan217User",
			lastSeen: "3 months ago"
		}

		const sosaUser: Friend = {
			name: "chiefSosa",
			lastSeen: "inactive"
		}

		const mcdonaldsHairline: Friend = {
			name: "mcdonaldsHairline",
			lastSeen: "banned",
		}

		const undefinedIsUndefined: Friend = {
			name: "undefinedIsUndefined",
			lastSeen: Date.now().toLocaleString()
		}

		const goClient: Friend = {
			name: "go_client",
			lastSeen: Date.now().toLocaleString()
		}

		const nodeClient: Friend = {
			name: "node_client",
			lastSeen: Date.now().toLocaleString()
		}

		const goClientContacts = [ladiesMan217User, sosaUser, mcdonaldsHairline, nodeClient]
		const nodeClientContacts = [mcdonaldsHairline, undefinedIsUndefined, goClient]

		mockDb.set("master_user", goClientContacts)
		mockDb.set("node_client", nodeClientContacts)

		console.log(" searching for:", user)

		const paint: Friend[] = mockDb.get(user)

		return paint
	}
}
