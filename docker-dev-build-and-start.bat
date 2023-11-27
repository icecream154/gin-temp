@echo off

setlocal enabledelayedexpansion
set timeout=-1
set env=dev

docker build -f ./dockerfiles/Dockerfile.%env% -t my-app:my-app-server-%env% .
docker rm -f my-app-server
docker run --name=my-app-server -dp 30000:30000 my-app:my-app-server-%env% --add-host=host.docker.internal:host-gateway