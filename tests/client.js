/**
 * @param {String} host_url
 * @param {Object} params
 * @returns {string}*/
function create_url(host_url, params) {
	const url = new URL(host_url)
	for (const [key, value] of Object.entries(params)) {
		url.searchParams.set(key, value)
	}
	return url.toString()
}

const url = create_url("ws://localhost:3900", {
	"user": "whosisland",
	"token": "mock_jwt_token_by_auth_server"
})

const socket = new WebSocket(url)

socket.addEventListener("open", (evt) => {
	console.log("connecion established")
	socket.send(sendJson())
})

socket.addEventListener("message", (evt) => {
	console.log("[recv]: ", Date.now())
	console.log(evt.data)
	console.log("\n\n")
})


socket.addEventListener("close", (evt) => {
	console.log("_closed", evt.code, evt.reason)
})


socket.addEventListener("error", (err) => {
	console.error("_error: ", err)
})


function sendJson() {
	const msg = {
		dest: "ladiesman217",
		from: "whosisland",
		time: Date.now().toString(),
		message: "yeah came in w the sauceee"
	}

	return JSON.stringify(msg)
}
