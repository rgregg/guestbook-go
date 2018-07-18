#!/bin/bash
export PROJECT_ID=rg-demo-1
export GOOGLE_APPLICATION_CREDENTIALS=/secret/private-key.json
export IMAGE_NAME=$1

if ["$IMAGE_NAME" == ""]; then
    export IMAGE_NAME='gcr.io/rg-demo-1/guestbook'
fi

echo "Running image ${IMAGE_NAME}"


docker run -v /Users/ryangregg/private-key.json:/secret/private-key.json -p 8080:8080 -e PROJECT_ID -e GOOGLE_APPLICATION_CREDENTIALS ${IMAGE_NAME}