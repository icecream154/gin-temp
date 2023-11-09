set timeout -1

env='dev'

docker build --tag my-app:my-app-server-$env -f ./dockerfiles/Dockerfile.$env .
docker rm -f my-app-server
docker run --name=my-app-server -dp 30000:30000 my-app:my-app-server-$env --add-host=host.docker.internal:host-gateway