package gamelogger


type GamesLogger interface {
	CreateNewCollection(string) (error)
	ArchiveCollection(string) (error)
	CollectionLength(string) (int, error)
	AddGameToCollection(string, string) error
}

