#!/bin/zsh
limactl start
sleep 1
./docker.sh -u
sleep 5
lima nerdctl ps
