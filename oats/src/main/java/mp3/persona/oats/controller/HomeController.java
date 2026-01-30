package mp3.persona.oats.controller;

import org.springframework.web.bind.annotation.*;
import org.springframework.stereotype.Controller;

@RestController
class HomeController {
	private String welcomeMessage = """
			\n
				Welcome to Oats. This is cli application
				that is a simulation of chat-rooms through the
				terminal.
				1. To learn more visit /docs.
				2. To Login visit /login
				3. To register visit /register
				4. To get guidance or contact customer service visit /help-me
				\n
				""";

	private String helpDoc = """
			\n
				This is the help-me forum
				I've got so many theories and suspicion
				Today is the day I follow my intuition
				\n
					""";

	@GetMapping("/welcome")
	public String welcomePage() {
		return welcomeMessage;
	}

	@GetMapping("/docs")
	public String documentation() {
		return "documentation.md";
	}

	@GetMapping("/help-me")
	public String helpMeHandler() {
		return helpDoc;
	}
}
