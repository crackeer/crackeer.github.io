# go开发如何热加载
go install  github.com/cosmtrek/air@master

```toml
root = "."
tmp_dir = "tmp"

[build]
# Just plain old shell command. You could use `make` as well.
cmd = "go run main.go"
# Binary file yields from `cmd`.
bin = "tmp/main"
# Customize binary.
full_bin = "./tmp/main"
# Watch these filename extensions.
include_ext = ["go"]
# Ignore these filename extensions or directories.
exclude_dir = ["template", "tmp"]
# Watch these directories if you specified.
include_dir = []
# Exclude files.
exclude_file = []
# This log file places in your tmp_dir.
log = "air.log"
# It's not necessary to trigger build each time file changes if it's too frequent.
delay = 1000 # ms
# Stop running old binary when build errors occur.
stop_on_error = true
# Send Interrupt signal before killing process (windows does not support this feature)
send_interrupt = false
# Delay after sending Interrupt signal
kill_delay = 500 # ms

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