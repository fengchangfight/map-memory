#!/bin/sh

cd /root/go/src/map-memory-backend/
git pull
go install

#export PID=`netstat -tulpn  | grep 8042 | awk '{print $NF}' | sed 's/\/\.\/map-memory-*//'`
export PID=`netstat -tulpn  | grep 8042 | awk '{print $NF}' | tr -d -c 0-9`

echo "Killing process $PID ..."
kill -9 $PID

echo "Switching to go bin directory..."
cd /root/go/bin
echo "Restarting the newly built process..."
nohup ./map-memory-backend prod &
