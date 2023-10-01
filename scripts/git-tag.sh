#!/bin/sh

#GIT_TAG=$(git describe --tags $(git rev-list --tags --max-count=1))

tagV=v2.2.1

echo "------------$tagV------------------------------"

git tag -a $tagV -m "$tagV"

git push origin $tagV