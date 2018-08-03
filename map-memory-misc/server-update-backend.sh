#!/bin/sh

cd /root/go/src/map-memory-backend/
git pull
go install

export PID=`netstat -tulpn  | grep 8042 | awk '{print $NF}' | sed 's/\/\.\/map-memory-*//'`
kill -9 $PID

cd /root/go/bin
nohup ./map-memory-backend prod &
