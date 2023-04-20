env='test'

docker pull icecr/youni:youni-intelligence-server-$env

docker rm -f youni-intelligence-server
docker run --name=youni-intelligence-server -v ~/youni_docker/dev_logs/intelligence:/storage/logs -dp 30000:30000 icecr/youni:youni-intelligence-server-$env --add-host=host.docker.internal:host-gateway
