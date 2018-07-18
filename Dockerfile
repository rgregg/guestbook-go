# We're using a pre-built image that pulls in our dependencies, instead of 
# taking the build time to pull them in here. See ./pre-build/ for more.

# FROM golang
FROM gcr.io/rg-demo-1/guestbook:source

# Document that the service listens on port 8080.
EXPOSE 8080