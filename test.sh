#! /bin/sh

if [ ! -d mirrors ]; then
  mkdir mirrors
  #git clone --mirror \
  #  git://github.com/isagalaev/highlight.js.git \
  #  mirrors/highlight.js.git
  #GIT_DIR=./mirrors/highlight.js.git git gc --aggressive
  curl -L \
    'https://www.dropbox.com/s/p7smueklddt8bws/highlight.js.git.tar?dl=0' \
    | tar xf - -C mirrors
fi
go run repo.go mirrors
