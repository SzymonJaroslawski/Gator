package main

import "errors"

type Commands struct {
	Cmds map[string]func(*State, Command) error
}

type Command struct {
	Name string
	Args []string
}

func (c *Commands) register(name string, f func(*State, Command) error) {
	c.Cmds[name] = f
}

func (c *Commands) run(s *State, cmd Command) error {
	f, ok := c.Cmds[cmd.Name]
	if !ok {
		return errors.New("command not found")
	}
	return f(s, cmd)
}
