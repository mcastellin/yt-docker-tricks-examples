#!/bin/sh

# IF container starts with root user,
# we fix data volume permissions then run CMD as appuser
if [ "$(id -u)" = "0" ]; then
    echo "Fixing data volume permissions"; echo
    chown -R appuser /data

    exec gosu appuser "$@"
fi

# ELSE we just run CMD if container starts as appuser
exec "$@"
