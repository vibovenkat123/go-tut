#!/bin/sh
root="github.com/vibovenkat123/go-tut"
filename=$1
go build -o bin/$filename $root/cmd/$filename  
