package mp3.persona.oats.repository;

import org.springframework.data.repository.CrudRepository;
import mp3.persona.oats.entities.Client;

public interface ClientRepository extends CrudRepository<Client, Integer> {
}
