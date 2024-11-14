#!/bin/bash
set -eo pipefail

host="127.0.0.1"
db="${1:-dev}"

if mongosh --quiet "$host/$db" --eval 'quit(db.runCommand({ ping: 1 }).ok ? 0 : 2)'; then
	exit 0
fi

exit 1