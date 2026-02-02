package mp3.persona.oats.controller;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseBody;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RestController;

import mp3.persona.oats.repository.ClientRepository;
import mp3.persona.oats.entities.Client;

@RestController
@RequestMapping(path = "/client")
public class ClientController {
	@Autowired // getting the bean for clientRepo
	private ClientRepository clientRepo;

	@PostMapping("/add")
	public @ResponseBody String addNewUser(@RequestBody Client client) {
		System.out.println("client[/add route hit]");
		System.out.println("[from]" + client.userName);
		clientRepo.save(client);

		String repoRpr = clientRepo.toString();
		System.out.printf("client saved to repo\n %s\n", repoRpr);
		return client.userName;
	}

}
