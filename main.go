package main

func main() {
	startRepl(cfg)
}

var cfg = &config{
	commands: getCommands(),
}
