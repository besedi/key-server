# Key server
_              ____  ______     __
| | _____ _   _/ ___||  _ \ \   / /
| |/ / _ \ | | \___ \| |_) \ \ / /
|   <  __/ |_| |___) |  _ < \ V /
|_|\_\___|\__, |____/|_| \_\ \_/
         |___/
## Quick start
> go run cmd/main.go --max-size 999 --srv-port 8000
### Build
> docker build --platform="linux/amd64" . -t besedi/key-server
### Local docker run
> docker run -p 1123:1123 besedi/key-server
### Push image to Dockerhub
> docker push besedi/key-server
