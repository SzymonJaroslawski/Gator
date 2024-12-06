package main

import (
	"context"

	"github.com/SzymonJaroslawski/Gator/internal/database"
)

func middlewareLogin(handler func(s *State, cmd Command, user database.User) error) func(*State, Command) error {
	return func(s *State, c Command) error {
		user, err := s.Db.GetUser(context.Background(), s.Config.CurrentUserName)
		if err != nil {
			return err
		}

		return handler(s, c, user)
	}
}
