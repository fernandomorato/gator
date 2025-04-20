package main

import "fmt"

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	username := cmd.Args[0]

	err := s.config.SetUser(username)
	if err != nil {
		return fmt.Errorf("couldn't set user %s: %v", username, err)
	}

	fmt.Printf("user set to: %s\n", username)
	return nil
}
