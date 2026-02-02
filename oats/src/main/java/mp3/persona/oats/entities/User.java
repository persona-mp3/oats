package mp3.persona.oats.entities;


public class User {
	private int MIN_CRED_LENGTH = 8;
	// might need to add setters and getters due to the 'OOP' way
	public String userName;
	public String email;
	public String password;

	public User(String userName, String password, String email) {
		this.userName = userName;
		this.password = password;
		this.email = email;
	}

	public boolean isValidCreds() {
		return validateUserName() && validatePassword() && validateEmail();
	}

	private boolean validateUserName() {
		return userName.length() >= MIN_CRED_LENGTH;
	}

	private boolean validatePassword() {
		return password.length() >= MIN_CRED_LENGTH;
	}

	private boolean validateEmail() {
		return email.length() > 5 && email.contains("@");
	}

}
