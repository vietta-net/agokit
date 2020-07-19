#!/bin/sh

cd ..

#Get the highest tag currency
#VERSION=`git describe --abbrev=0 --tags`
VERSION=`git ls-remote --tags --refs --sort="version:refname" git://github.com/vietta-net/agokit.git | awk -F/ 'END{print$NF}'`
VERSION=${VERSION:-'0.1.0'}

#Get currency parts
MAJOR="${VERSION%%.*}"; VERSION="${VERSION#*.}"
MINOR="${VERSION%%.*}"; VERSION="${VERSION#*.}"
PATCH="${VERSION%%.*}"; VERSION="${VERSION#*.}"

#Get current hash and see if it already has a tag
GIT_COMMIT=`git rev-parse HEAD`
NEEDS_TAG=`git describe --contains $GIT_COMMIT`

#Only tag if no tag already (would be better if the git describe command above could have a silent option)
if [ -z "$NEEDS_TAG" ]; then
    #Increase version
    PATCH=$((PATCH+1))
    #Create new tag
    NEW_TAG="$MAJOR.$MINOR.$PATCH"
    git tag "$NEW_TAG" -m "Release $NEW_TAG"
    echo "Updating to $NEW_TAG"
    git push origin "$NEW_TAG"
    echo "Push"
    git push
else
    VERSION="$MAJOR.$MINOR.$PATCH"
    echo "Already a tag (version $VERSION) on this commit"
fi