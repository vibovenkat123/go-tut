#!/bin/sh
root="github.com/vibovenkat123/go-tut"
echo "name: "
read filename

go build -o bin/$filename $root/cmd/$filename  
