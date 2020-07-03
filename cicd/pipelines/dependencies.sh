#!/usr/bin/env sh

set -eu

# Add python pip and bash
#apk add --no-cache py-pip bash
apk add --no-cache python3 py3-pip3 bash

# Install docker-compose via pip
pip install --no-cache-dir docker-compose
docker-compose -v
