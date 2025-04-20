package main

import (
	"log"
	"os"

	"github.com/fernandomorato/gator/internal/config"
)

type state struct {
	config *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	programState := state{
		config: &cfg,
	}

	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)

	args := os.Args
	if len(args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}
	cmdName := args[1]
	cmdArgs := args[2:]
	err = cmds.run(&programState, command{
		Name: cmdName,
		Args: cmdArgs,
	})
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
