package mp3.persona.oats.entities;

public class User {
	private int MIN_CRED_LENGTH = 4;
	// might need to add setters and getters due to the 'OOP' way
	public String user;
	public String password;

	public User(String user, String password) {
		this.user = user;
		this.password = password;
	}

	public boolean isValidCreds() {
		return validateUserName() && validatePassword();
		// return validateUserName() && validatePassword() && validateEmail();
	}

	private boolean validateUserName() {
		return user.length() >= MIN_CRED_LENGTH;
	}

	private boolean validatePassword() {
		return password.length() >= MIN_CRED_LENGTH;
	}

	// private boolean validateEmail() {
	// 	return email.length() > 5 && email.contains("@");
	// }

}
