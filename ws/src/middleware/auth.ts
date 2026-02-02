type User = {
	authToken?: string
	userName: string,
	password: string,
	email: string
}


function mockSharedDB(): Map<string, User> {
	const _masterUserName = "childish_gambino"
	const db = new Map<string, User>()
	const masterUser: User = {
		authToken: "don't_be_scandalous",
		userName: "childish_gambino",
		password: "awaken_my_love",
		email: "donaldglover@spotify.com"
	}

	db.set(_masterUserName, masterUser)

	return db
}

type Friend = {
	name: string,
	status: boolean,
	lastSeen: string
}

type AuthStatus = {
	status: boolean,
	info?: Map<string, Friend>
}

function masterCompadres(): Map<string, Friend> {
	const friends: Friend = { name: "jia_tian", status: false, lastSeen: "" }
	return new Map<string, Friend>([["jia_tian", friends]])

}

class Database {
	#users: Map<string, User>
	constructor() {
		this.#users = mockSharedDB()
	}

	simulateConnection(): boolean {

		let counter = 0;
		let interval = setInterval(function () {
			counter += 1
			if (counter === 2) {
				console.log("connected to the database")
				clearInterval(interval)
			}
			console.log("connecting to database...")

		}, 1300)

		console.log("connected to databse")
		return true;
	}

	findUser(userName: string, bearerToken: string): AuthStatus {
		let authStats: AuthStatus = {
			status: false,
		}; 

		const userExists = this.#users.has(userName)
		if (!userExists) {
			console.error("user doesn't exist", userName, userExists)
			return authStats
		}

		console.log("user exists, extracting data")
		const user: User | undefined = this.#users.get(userName)
		if (!user) {
			console.error("little patience for ts, but incase something like this happens...")
			throw new Error("little patience for ts, but incase something like this happens...")
		}


		console.log("extracted-token", bearerToken)
		if (user.authToken === bearerToken) {
			console.log("tokens match")
			authStats.info = masterCompadres()
			authStats.status = true
			return authStats
		}

		return authStats
	}
}

const DB = new Database()

export function authClient(userName: string, bearerToken: string): AuthStatus {

	const isConnected = DB.simulateConnection()
	if (!isConnected) throw new Error("Yoo twin....Database could not be connected yo, what shall ye be done????😭✌️")

	return DB.findUser(userName, bearerToken)
}
