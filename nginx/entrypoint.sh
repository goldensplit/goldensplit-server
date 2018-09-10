#!/bin/sh

set +e
nginx -t
while [[ $? -ne 0 ]]; do
    sleep 2
    echo "Failed to parse configuration files... restarting"
    nginx -t
done
nginx -g 'daemon off;'
