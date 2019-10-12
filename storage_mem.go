package main

// StorageMemory data storage layered save only in memory
type StorageMemory struct {
	People []Person
}

//GetPeople gets people from memory
func (s *StorageMemory) GetPeople() ([]Person, error) {
	return s.People, nil
}

//GetPerson gets person from memory
func (s *StorageMemory) GetPerson(person Person) (Person, error) {
	var foundPerson Person

	for _, p := range s.People {
		if p.ID == person.ID {
			foundPerson = p
		}
	}

	return foundPerson, nil
}

//SavePerson saves a new person
func (s *StorageMemory) SavePerson(p Person) error {
	s.People = append(s.People, p)

	return nil
}
