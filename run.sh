#!/bin/sh

dep ensure

# node process kill
killall node

cd assets
yarn install
yarn run dev &
cd ../

go run main.go