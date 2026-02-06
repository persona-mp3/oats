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
				"childish_gambino",
				new User("childish_gambino", "awaken_my_love", "donaldglover@spotify.com"));

		oaties.put(
				"billevans",
				new User("billevans", "password1234", "billevans@gmail.com"));

		oaties.put(
				"west_side_gunn",
				new User("west_side_gunn", "327_tyler", "westsidegunn@spotify.com"));

	}

	public static boolean findUser(User userCreds) {
		if (!oaties.containsKey(userCreds.userName)) {
			return false;
		}

		User userDetails = oaties.get(userCreds.userName);
		if (!userDetails.email.equals(userCreds.email)) {
			System.out.println("failed-auth: email");
			System.out.printf("%s | :%s\n", userDetails.email, userCreds.email);
			return false;
		}

		if (!userDetails.password.equals(userCreds.password)) {
			System.out.printf("failed-auth: passwd");
			System.out.printf("%s | :%s\n", userDetails.password, userCreds.email);
			return false;
		}

		return true;
	}
}
