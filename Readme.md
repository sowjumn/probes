## Build the Docker Image 
`docker build --tag probes .`

## Spin up a container with the image `probe`
`docker run -p 8080:8080 probe`

## Tag the image so its suits Docker hub format
`docker tag probes mountainpinelake/probe-hello`