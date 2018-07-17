# Rebuild the container image
docker build . -t gcr.io/rg-demo-1/guestbook

# Push it to gcr.io
docker push gcr.io/rg-demo-1/guestbook