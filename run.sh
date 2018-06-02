#!/bin/sh

# node process kill
killall node

cd assets
yarn run dev &
cd ../

go run main.go