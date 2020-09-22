#!/usr/bin/env bash
echo 'Running migrations...'
/go-cloudnative-practice/migrate up >/dev/null 2>&1 &

echo 'Deleting mysql-client...'
apk del mysql-client

echo 'Start application...'
/go-cloudnative-practice/app
