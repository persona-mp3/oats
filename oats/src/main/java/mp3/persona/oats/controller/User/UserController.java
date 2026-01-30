package mp3.persona.oats.controller.User;

import org.springframework.web.bind.annotation.*;
import mp3.persona.oats.entities.User;

@RestController
// might use OAuth for authentication or feeling funky shh
class UserContoller {
	@PostMapping("/register")
	public String loginUser(@RequestBody User user) {
		String clientUserName = user.userName;
		String clientPassword = user.password;
		String clientEmail = user.email;

		if (!user.isValidCreds()) {
			return "invalid credentails boi";
		}

		System.out.printf("%s,  has a  %s , with a domain of %s\n", clientUserName, clientPassword, clientEmail);
		return "valid credentials boi";
	}

	@PostMapping("/login")
	public void registerUser() {
	}
}
