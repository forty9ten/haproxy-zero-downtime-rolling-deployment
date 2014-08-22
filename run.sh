#!/bin/sh

./iptables-init.sh
./iptables-add.sh

session=forty9ten
tmux new-session -ds "$session"

tmux send-keys "service haproxy restart" "C-m"
tmux send-keys "sleep 5 && while true; do curl localhost/v1/dowork; sleep 1; echo; done" "C-m"
tmux split-window -h  "go run /app/server.go -port 4910"
tmux split-window -h  "go run /app/server.go -port 4911"
