const serverAddr = "ws://localhost:3900"

function manufactureUrl(host_url, params) {
	const url = new URL(serverAddr)
	for (const [key, value] of Object.entries(params)) {
		url.searchParams.set(key, value)
	}
	return url.toString()
}


async function ladiesMan217User() {
	const url = manufactureUrl(serverAddr, { user: "ladiesMan217", token: "ladies_man_217_token" })
	const socket = new WebSocket(url)

	socket.addEventListener("open", (evt) => {
		socket.send(sendJson("answer me, are you laidesman217?", "chief_keef", "ladiesman217"))
	})
	socket.addEventListener("message", (evt) => {
		console.log(" [#] ")
		console.log(evt.data)
	})

	socket.addEventListener("close", (evt) => {
		console.log("ladiesman217 close: code: %d, reason: %s", evt.code, evt.reason)
	})

	socket.addEventListener("error", (err) => {
		console.error("ladiesman217 error:")
		console.error(err)
	})
}

async function chiefKeefUser() {
	const url = manufactureUrl(serverAddr, { user: "chief_keef", token: "your_gonna_be_so_t1redoFsomucHw1nniG" })
	const socket = new WebSocket(url)

	socket.addEventListener("open", (evt) => {
		socket.send(sendJson("uuyi boy don't want war!", "ladiesman217", "chief_keef"))
	})
	socket.addEventListener("message", (evt) => {
		console.log(" [#] %s\n")
		console.log(evt.data)
	})

	socket.addEventListener("close", (evt) => {
		console.log(" sosa close: code: %d, reason: %s", evt.code, evt.reason)
	})

	socket.addEventListener("error", (err) => {
		console.error(" sosa error:")
		console.error(err)
	})
}

function sendJson(text, to, src) {
	const msg = {
		dest: to,
		from: src,
		time: Date.now().toString(),
		message: text
	}

	return JSON.stringify(msg)
}

await ladiesMan217User()
// await chiefKeefUser()
