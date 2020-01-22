#!/bin/bash -e

docker login -u "$QUAY_USER" -p "$QUAY_PASSWORD" quay.io

MAKE="make"

$MAKE docker-build DOCKER_REGISTRY=quay.io/
$MAKE docker-push DOCKER_REGISTRY=quay.io/

if [ -n "$TRAVIS_TAG" ]; then
    $MAKE docker-push-version DOCKER_REGISTRY=quay.io/ PUSH_DOCKER_TAG="$TRAVIS_TAG"
else
    $MAKE docker-push-version DOCKER_REGISTRY=quay.io/ PUSH_DOCKER_TAG="$TRAVIS_BRANCH"
fi
