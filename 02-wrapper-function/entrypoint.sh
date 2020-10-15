#!/bin/sh

# if container starts with root user,
# we fix data volume permissions then run cmd as appuser
if [ "$(id -u)" = "0" ]; then
    echo "Fixing data volume permissions"
    chown -R appuser /data

    exec gosu appuser "$@"
fi

# else CMD is started as current user
exec "$@"
