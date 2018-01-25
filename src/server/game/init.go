package game

func Init() bool {
	if !LoadServerConfig("config.json") {
		return false
	}

	LoadServerData()

	return true
}
