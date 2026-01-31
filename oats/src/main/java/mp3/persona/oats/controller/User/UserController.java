package mp3.persona.oats.controller.User;

import mp3.persona.oats.utils.Utils;

import org.springframework.web.bind.annotation.*;
import org.springframework.util.LinkedMultiValueMap;

import org.springframework.util.MultiValueMap;

import org.springframework.http.ResponseEntity;
import org.springframework.http.HttpStatus;
import org.springframework.http.HttpHeaders;

import mp3.persona.oats.repository.Users;
import mp3.persona.oats.entities.User;

@RestController
// might use OAuth for authentication or feeling funky shh
class UserController {
	// will later actual repositories and mongoDB
	private String redirectBaseLink = "http://localhost:3900/";

	@PostMapping("/register")
	public ResponseEntity<?> loginUser(@RequestBody User userReq) {
		if (!userReq.isValidCreds()) {
			return sendBadRequest("credenatials do not meet requirements");
		}

		// ===== simulate connecting to db =====
		Users.onConnect();

		if (!Users.findUser(userReq)) {
			return sendBadRequest("invalid credentials or user doesn't exist");
		}

		String authToken = "dont_be_scandolous";
		MultiValueMap<String, String> mulMap = new LinkedMultiValueMap<>();
		// You can add other query params here
		mulMap.add("authToken", authToken);

		try {
			Utils.parseUrl(redirectBaseLink, mulMap);
		} catch (Exception err) {
			System.out.printf("unexpected erorr occured:", err);
			return ResponseEntity.internalServerError().build();
		}

		// TEST: removing auth_query params for now to test flow
		return sendRedirectToWSServer(authToken = "");
	}

	private ResponseEntity<Void> sendRedirectToWSServer(String authToken) {
		String redirectLink = redirectBaseLink + authToken; // use uri to parse link instead
		return ResponseEntity
				.status(HttpStatus.FOUND).header(HttpHeaders.LOCATION, redirectLink)
				.build();
	}

	private ResponseEntity<String> sendBadRequest(String msg) {
		return ResponseEntity
				.status(HttpStatus.BAD_REQUEST)
				.body(msg);
	}

	@PostMapping("/login")
	public void registerUser() {
	}
}
