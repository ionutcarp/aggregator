package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/ionutcarp/aggregator/internal/config"
	"github.com/ionutcarp/aggregator/internal/database"
	"time"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %v <name>", cmd.Name)
	}

	loginUser, err := s.db.GetUser(context.Background(), cmd.Args[0])
	if err != nil {
		return fmt.Errorf("user does not exist: %w", err)
	}

	err = s.config.SetUser(loginUser.Name)
	if err != nil {
		return fmt.Errorf("could not set user: %w", err)
	}

	fmt.Printf("User has been set to: %s\n", s.config.CurrentUserName)
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      cmd.Args[0],
	})
	if err != nil {
		return fmt.Errorf("could not create user: %w", err)
	}

	err = s.config.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("could not set user: %w", err)
	}

	fmt.Println("User created successfully:")
	printUser(user)
	return nil
}

func handlerListUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("could not get users: %w", err)
	}
	cfg, err := config.Read()
	if err != nil {
		return fmt.Errorf("could not read config: %w", err)
	}

	for _, user := range users {
		if user == cfg.CurrentUserName {
			fmt.Printf(" * %v (current)\n", user)
			continue
		}
		fmt.Printf(" * %v\n", user)
	}
	return nil
}

func printUser(user database.User) {
	fmt.Printf(" * ID:      %v\n", user.ID)
	fmt.Printf(" * Name:    %v\n", user.Name)
}
