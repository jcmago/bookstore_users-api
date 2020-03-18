# bookstore_users_api
USERS API

# GO - Udemy course - How-to-design-and-build-rest-microservices-in-go

## Basic commands

```bash
export GPPATH=/Users/juancarlos.gomez/go
go build main.go
go run main.go
go fmt main.go
go install main.go
go get main.go
go test main.go
```

## Libraries

- GIN - https://gin-gonic.com - HTTP Engine

## Executable or reusable

```bash
go get -u github.com/gin-gonic/gin
go get -u github.com/jcmago/bookstore_users-api
go get -u github.com/jcmago/bookstore_items-api
go get -u github.com/jcmago/bookstore_oauth-api
```

## Lessons

- https://www.udemy.com/course/golang-how-to-design-and-build-rest-microservices-in-go/learn/lecture/16451510#overview


## Benchmarking

- https://opensource.com/article/18/6/copying-files-go


## Debugging

Hot reload

https://github.com/cosmtrek/air

```bash
go get -u github.com/cosmtrek/air
touch .air.conf
air -c .air.conf
```

```.air.conf
# Config file for [Air](https://github.com/cosmtrek/air) in TOML format

# Working directory
# . or absolute path, please note that the directories following must be under root.
root = "."
tmp_dir = "tmp"

[build]
# Just plain old shell command. You could use `make` as well.
cmd = "go build -o ./tmp/main ."
# Binary file yields from `cmd`.
bin = "tmp/main"
# Customize binary.
full_bin = "APP_ENV=dev APP_USER=air ./tmp/main"
# Watch these filename extensions.
include_ext = ["go", "tpl", "tmpl", "html"]
# Ignore these filename extensions or directories.
exclude_dir = ["assets", "tmp", "vendor", "frontend/node_modules"]
# Watch these directories if you specified.
include_dir = []
# Exclude files.
exclude_file = []
# It's not necessary to trigger build each time file changes if it's too frequent.
delay = 1000 # ms
# Stop to run old binary when build errors occur.
stop_on_error = true
# This log file places in your tmp_dir.
log = "air_errors.log"

[log]
# Show log time
time = false

[color]
# Customize each part's color. If no color found, use the raw app log.
main = "magenta"
watcher = "cyan"
build = "yellow"
runner = "green"

[misc]
# Delete tmp directory on exit
clean_on_exit = true
```