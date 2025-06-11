# Key server
## Quick start
> go run cmd/main.go
### Build
> docker build --platform="linux/amd64" . -t besedi/key-server
### Local docker run
> docker run -p 1123:1123 besedi/key-server
### Push image to Dockerhub
> docker push besedi/key-server
