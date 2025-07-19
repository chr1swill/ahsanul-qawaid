#!/bin/bash

#set -xe

root=/home/chr1swill/code/projects/ahsanul-qawaid

[ ! -d "${root}/src" ] && echo "no src folder found" && exit 1

mkdir -p "${root}/bin"

go build -o "${root}/bin/main" "${root}/src/main.go"

exit 0
