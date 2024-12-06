package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/SzymonJaroslawski/Gator/internal/config"
	"github.com/SzymonJaroslawski/Gator/internal/database"
	"github.com/golang/glog"
	_ "github.com/lib/pq"
)

type State struct {
	Config *config.Config
	Db     *database.Queries
}

func main() {
	config, err := config.Read()
	if err != nil {
		glog.Fatalf("Error: %s", err)
	}

	state := State{
		Config: config,
	}

	db, err := sql.Open("postgres", state.Config.DbURL)
	if err != nil {
		log.Fatal(err)
	}

	state.Db = database.New(db)

	cmds := Commands{
		Cmds: make(map[string]func(*State, Command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", middlewareLogin(handleAddFeed))
	cmds.register("feeds", handleFeeds)
	cmds.register("follow", middlewareLogin(handleFollow))
	cmds.register("following", middlewareLogin(handleFollowing))
	cmds.register("unfollow", middlewareLogin(handleUnfollow))
	cmds.register("browse", middlewareLogin(handleBrowse))

	if len(os.Args) < 2 {
		fmt.Println("Usage: cli <command> [args...]")
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = cmds.run(&state, Command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}
}
