#!/bin/sh

# if container starts with root user,
# we fix data volume permissions then run CMD as appuser
if [ "$(id -u)" = "0" ]; then
    echo "Fixing data volume permissions"
    chown -R appuser /data

    exec gosu appuser "$@"
fi

# else we just run CMD if container starts as appuser user
exec "$@"
