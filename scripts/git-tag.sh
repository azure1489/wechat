#!/bin/sh

GIT_TAG=$(git describe --tags $(git rev-list --tags --max-count=1))


echo "test$GIT_TAG"

