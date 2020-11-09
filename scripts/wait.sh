#!/bin/bash

TIMEOUT="${TIMEOUT:-20}"

echo "Waiting ${TIMEOUT} sec for ACP to be ready"
timeout "${TIMEOUT}" bash -c 'while [[ "$(curl -s -k -o /dev/null -w "%{http_code}" https://localhost:8443/alive)" != 200 ]]; do sleep 1; done' || exit 1
echo "ACP is ready"
