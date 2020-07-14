#! /bin/bash

if [ "$1" != "" ]; then
  cd ..
  git add .
  git commit -m "$1"
else
    echo "Please enter a message"
fi