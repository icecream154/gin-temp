env='test'

docker buildx build --platform linux/amd64 --tag youni:youni-intelligence-server-$env -f ./dockerfiles/Dockerfile.$env .

docker tag youni:youni-intelligence-server-$env icecr/youni:youni-intelligence-server-$env
docker push icecr/youni:youni-intelligence-server-$env
