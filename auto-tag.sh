#!/bin/sh

#Get the highest tag number
#VERSION=`git describe --abbrev=0 --tags`
VERSION=`git ls-remote --tags --refs --sort="version:refname" git://github.com/vietta-net/agokit.git | awk -F/ 'END{print$NF}'`
VERSION=${VERSION:-'0.1.0'}

#Get number parts
MAJOR="${VERSION%%.*}"; VERSION="${VERSION#*.}"
MINOR="${VERSION%%.*}"; VERSION="${VERSION#*.}"
PATCH="${VERSION%%.*}"; VERSION="${VERSION#*.}"

#Increase version
PATCH=$((PATCH+1))

#Get current hash and see if it already has a tag
GIT_COMMIT=`git rev-parse HEAD`
NEEDS_TAG=`git describe --contains $GIT_COMMIT`

#Create new tag
NEW_TAG="$MAJOR.$MINOR.$PATCH"
echo "Updating to $NEW_TAG"

#Only tag if no tag already (would be better if the git describe command above could have a silent option)
if [ -z "$NEEDS_TAG" ]; then
  echo "Tagged with $NEW_TAG (Ignoring fatal:cannot describe - this means commit is untagged) "
    git tag "$NEW_TAG" -m "Release $NEW_TAG"
    git push origin "$NEW_TAG"
    git push
else
    echo "Already a tag on this commit"
fi