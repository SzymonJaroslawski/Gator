# Gator

CLI RSS post aggregator

## Requirements

You will need to install Postgres, Go and Goose on your system

1. Install Postgres, Go and Goose
2. Start Postgres server (Setup it as per your distribution docs or Postgres docs)
3. Create new database
4. Clone this repo
5. Run these
```
  go install github.com/pressly/goose/v3/cmd/goose@latest
  // From root in the folder where you cloned the repo
  cd sql/schema
  goose <link_to_the_database> up
  cd ../..
  go install .
```
6. Run "Gator" to generate the config file
7. Add your database connection link to the config (XDG_CONFIG_HOME) example:
```
  nvim ~/.config/gator/.gatorconfig.json

  {
    "db_url":"postgres://postgres:postgres@localhost:5432/gator?sslmode=disable",
    "current_user":"user1"
  }
```

You will probably want to remove the cloned repo after installing the program.

## Usage

First you need to register new user. Command register always logins you into the new user.

```
  Gator register user1
```

You can switch users with "Gator login <user_name>"

```
  Gator login user2
```

List register users with "Gator users"

```
  Gator users
  /*
  Users: 2
  * user1
  * user2 (current)
  */
```

You can reset the whole database with "Gator reset"

```
  Gator reset
```

Add feeds with "Gator addfeed <name> <url>". The user you are login will also start following it.

```
  Gator add "Boot.dev" "https://blog.boot.dev/index.xml"
```

List all feeds with "Gator feeds"

```
  Gator feeds
```

Follow feeds added by other users with "Gator follow <url>"

```
  Gator follow "https://blog.boot.dev/index.xml"
```

List follows with "Gator following"

```
  Gator following
```

Unfollow feeds with "Gator unfollow <url>"

```</url>
  Gator unfollow "https://blog.boot.dev/inde.xml"
```

To aggregate posts from feeds use "Gator agg <interval>"

```
  // aggregate every minute
  Gator agg 1m

  // aggregate every hour
  Gator agg 1h
```

Use "Gator browse <limit>" to browse post from followed feeds by the user 

```
  Gator browse 10
```

## Build

Versions:
```
  - postgres (PostgreSQL) 16.3
  - go version go1.23.4 linux/amd64
  - goose version: v3.23.0
  - sqlc v1.27.0
```
