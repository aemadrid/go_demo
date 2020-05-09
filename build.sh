#!/usr/bin/env zsh
clear
echo "> building go_demo ..."
docker build -t go_demo:latest .
echo "> running go_demo ..."
docker run -P --name go_demo go_demo:latest
echo "done with go_demo ..."