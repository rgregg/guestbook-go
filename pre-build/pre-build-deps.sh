#!/bin/bash

export PROJ_ID=rg-demo-1

# Prebuild our GO dependencies, so we don't have to wait for them
docker build . -t gcr.io/${PROJ_ID}/golang-gcpdatastore:latest

# Push to Google Container Registry
docker push gcr.io/${PROJ_ID}/golang-gcpdatastore:latest

echo "Pushed new container image to gcr.io/${PROJ_ID}/golang-gcpdatastore:latest"