package mp3.persona.oats;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
public class OatsApplication {

	public static void main(String[] args) {
		System.out.println("server active at: http://localhost:8900");
		SpringApplication.run(OatsApplication.class, args);
	}

}
