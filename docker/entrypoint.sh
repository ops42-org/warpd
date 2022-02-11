#!/bin/bash

# echo "Hello from Warpd image!"
cd "$GITHUB_WORKSPACE"
# env
# pwd
# ls -lah

exec warpd $@
