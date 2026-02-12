package mp3.persona.oats.repository;

import mp3.persona.oats.entities.User;

import java.util.*;

public class Users {
	private static Map<String, User> oaties = new HashMap<>();

	public static void onConnect() {
		loadUsers();
	}

	private static void loadUsers() {
		oaties.put(
				"master_user",
				new User("master_user", "m@ster_password"));

		oaties.put(
				"billy",
				new User("billy_password", "peacepeice3000"));


	}

	public static boolean findUser(User userCreds) {
		System.out.println("good problems?");
		if (!oaties.containsKey(userCreds.user)) {
			return false;
		}

		User userDetails = oaties.get(userCreds.user);
		// if (!userDetails.email.equals(userCreds.email)) {
		// 	System.out.println("different emails, for userName, emailAuth");
		// 	System.out.printf("%s vs :%s\n", userDetails.email, userCreds.email);
		// 	return false;
		// }

		if (!userDetails.password.equals(userCreds.password)) {
			System.out.printf("Invalid password");
			System.out.printf("%s vs :%s\n", userDetails.password, userCreds.password);
			return false;
		}

		System.out.println("good problems?");

		return true;
	}
}
