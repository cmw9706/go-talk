package main

// StorageType defines available storage types
type StorageType int

const (
	// JSONDataLocation defines where the JSON data files will be saved
	JSONDataLocation = "./data/"

	// JSON will store data in JSON files saved on disk
	JSON StorageType = iota

	// Memory will store data in memory
	Memory
)

//Storage interface: db or in mem storage
type Storage interface {
	GetPeople() ([]Person, error)
	GetPerson(Person) (Person, error)
	SavePerson(Person) error
}

//NewStorage Creates a storage object
func NewStorage(stoType StorageType) (Storage, error) {
	var stg Storage
	var err error

	switch stoType {
	case Memory:
		stg = new(StorageMemory)

		// case JSON:
		// 	// for the moment storage location for JSON files is the current working directory
		// 	stg, err = NewStorageJSON(JSONDataLocation)
	}

	return stg, err
}
