# Key server
```
_              ____  ______     __
| | _____ _   _/ ___||  _ \ \   / /
| |/ / _ \ | | \___ \| |_) \ \ / /
|   <  __/ |_| |___) |  _ < \ V /
|_|\_\___|\__, |____/|_| \_\ \_/
         |___/

A lightweight and secure key server
```
## Description
Simple server which generates key with given length in binary format

## Quick start
> go run cmd/main.go --max-size 999 --srv-port 8000
### Build
> docker build --platform="linux/amd64" . -t besedi/key-server
### Local docker run
> docker run -p 1123:1123 besedi/key-server
### Push image to Dockerhub
> docker push besedi/key-server

## Helm
helm upgrade -i -f helm/values.yaml key-server helm/ -n YOUR_NAMESPACE

## Monitoring
- Alert on high rate of 4XX/5XX codes
- Alert when the 70th, 90th, or 99th percentile of key length distribution exceeds defined thresholds
