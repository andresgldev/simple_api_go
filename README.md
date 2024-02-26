# simple_api_go


cd cmd/api/
docker build -t "myapp:latest" .
docker run --publish 8080:8080 myapp:latest