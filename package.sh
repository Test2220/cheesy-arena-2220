#!/bin/sh
# set -e
rm -rf db/backups
rm -rf static/logs
mkdir -p static/logs
go clean
go build
zip -r -X cheesy-arena.zip LICENSE README.md access_point_config.tar.gz cheesy-arena cheesy-arena.command db font schedules static switch_config.txt templates
