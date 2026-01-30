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
				"bill_evans_onkey",
				new User("bill_evans_onkey", "peacepeice3000", "billevans@spotify.com"));

		oaties.put(
				"west_side_gunn",
				new User("west_side_gunn", "327_tyler", "westsidegunn@spotify.com"));

	}

	public static boolean findUser(User userCreds) {
		System.out.println("good problems?");
		if (!oaties.containsKey(userCreds.userName)) {
			return false;
		}

		User userDetails = oaties.get(userCreds.userName);
		if (!userDetails.email.equals(userCreds.email)) {
			System.out.println("different emails, for userName, emailAuth");
			System.out.printf("%s vs :%s\n", userDetails.email, userCreds.email);
			return false;
		}

		if (!userDetails.password.equals(userCreds.password)) {
			System.out.printf("Invalid password");
			System.out.printf("%s vs :%s\n", userDetails.password, userCreds.email);
			return false;
		}

		System.out.println("good problems?");

		return true;
	}
}
