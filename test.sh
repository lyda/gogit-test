#! /bin/sh

if [ ! -d mirrors ]; then
  mkdir mirrors
  git clone --mirror \
    git://github.com/isagalaev/highlight.js.git \
    mirrors/highlight.js.git
  GIT_DIR=./mirrors/highlight.js.git git gc --aggressive
fi
go run repo.go mirrors
