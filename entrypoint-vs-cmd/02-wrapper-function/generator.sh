#!/bin/bash

# This sample script will generate a few GUIDs and
# append them to the keys file in the /var/lib/myapp/

DATA_DIR=/data

echo "Adding 5 keys to keys.txt ..."
for i in {1..5}; do
    uuidgen >>$DATA_DIR/keys.txt
done

echo "The following keys have been generated and added to $DATA_DIR/keys.txt:"
echo
tail -n 5 ${DATA_DIR}/keys.txt

echo
echo "Done. Exiting."
