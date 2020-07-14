#! /bin/bash

if [ "$1" != "" ]; then
  cd ..
  git add .
  git commit -m "$1"
  git push
else
    echo "Please enter a message"
fi