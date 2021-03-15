docker kill $(docker ps -q)
docker rm ${docker ps -q}
rm -rf ./pgdata/

go fmt ./...

docker-compose build
docker-compose up -d database

echo "Sleeping for 20s..."
sleep 20

docker-compose up backend