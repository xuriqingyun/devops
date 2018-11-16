#!/bin/sh

cd /apps/gopath/src/devops
git pull
cd /apps/gopath/src/devops/newWeb
go build
ps -ef |grep newWeb |awk '{print $2}'|xargs kill -9
./newWeb &
echo "deploy progress is complete!"
