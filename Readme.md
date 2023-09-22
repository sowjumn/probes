## Build the Docker Image 
`docker build --tag probes .`

## Spin up a container with the image probe
`docker run -p 8080:8080 probe`