#!/bin/bash

# This sample script will generate a few GUIDs and
# append them to the keys file in the /var/lib/myapp/

DATA_DIR=/var/lib/myapp

echo "Adding 5 keys to keys.txt"
for i in {1..5}; do
    uuidgen >>$DATA_DIR/keys.txt
done

echo "Done. Exiting."
