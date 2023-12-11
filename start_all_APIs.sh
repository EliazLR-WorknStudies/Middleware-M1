#!/bin/bash

OPENDIR=$(pwd)

# Launch API Users
gnome-terminal --tab --title="API Users" --working-directory="$OPENDIR" -- bash -c "cd ./API_Users/ && go run ./cmd/main.go; exec bash"

# Launch API Songs
gnome-terminal --tab --title="API Songs" --working-directory="$OPENDIR" -- bash -c "cd ./API_Songs/tp_middleware_example-main/ && go run ./cmd/main.go; exec bash"

# Launch API Rating
gnome-terminal --tab --title="API Rating" --working-directory="$OPENDIR" -- bash -c "cd ./API_Rating/ && go run ./cmd/main.go; exec bash"

# Launch Flask
gnome-terminal --tab --title="Flask" --working-directory="$OPENDIR" -- bash -c "cd ./Flask/ && PYTHONPATH=$PYTHONPATH:$(pwd) python3 src/app.py; exec bash"
