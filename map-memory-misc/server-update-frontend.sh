#!/bin/sh
export PID=`netstat -tulpn  | grep 8043 | awk '{print $NF}' | sed 's/\/node//'`
kill -9 $PID

cd /opt/map-memory
git pull --rebase
rm -rf /opt/workspace/map-memory
cp -rf /opt/map-memory /opt/workspace/
cp /opt/workspace/http-commons.js /opt/workspace/map-memory/common/
cd /opt/workspace/map-memory
npm install
nohup npm run build &
