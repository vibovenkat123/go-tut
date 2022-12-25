#!/bin/sh
root="github.com/vibovenkat123/go-tut"
echo "name: "
read filename

go run $root/cmd/$filename
