#!/bin/bash

echo "Hello from builder image!"
cd "$GITHUB_WORKSPACE"
env
pwd
ls -lah
