# We're using a pre-built image that pulls in our dependencies, instead of 
# taking the build time to pull them in here. See ./pre-build/ for more.

# FROM golang
FROM gcr.io/rg-demo-1/golang-gcpdatastore:latest

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/rgregg/guestbook-go
COPY ./static /static

RUN go install github.com/rgregg/guestbook-go

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/guestbook-go

# Document that the service listens on port 8080.
EXPOSE 8080