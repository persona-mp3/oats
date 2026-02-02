package mp3.persona.oats.entities;

import jakarta.persistence.Entity;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.GenerationType;
import jakarta.persistence.Id;

@Entity // tells hibernate? to make a table out of the class
public class Client {
	@Id
	@GeneratedValue(strategy=GenerationType.AUTO)
	private Integer id;
	public String userName;
	public String password;

	// how do beans work in springboot/java
}
