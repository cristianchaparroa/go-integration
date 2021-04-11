#!/bin/sh
# wait-for-sftp

set -e

host="$1"
shift
cmd="$@"

until sftp -p 22 127.0.0.11 22; do
  >&2 echo "SFTP server is unavailable - sleeping"
  sleep 1
done

>&2 echo "SFTP is up - executing command"
exec $cmd
