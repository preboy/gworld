package game

func Init() bool {
	if !LoadServerConfig("server_config.json") {
		return false
	}
	return true
}
