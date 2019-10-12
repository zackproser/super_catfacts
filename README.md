# Super CatFacts Service 
A ridiculously over-engineered CatFacts prank service written in Golang and deployed via Kubernetes

# Pre-requisites 
* A funded Twilio account 
* A valid Twilio phone number 
* A working Kubernetes cluster
* A working Golang installation
* A valid domain name for mapping to the service 

# Getting started 
```
# Clone the repo and build the binary 
git clone github.com/zackproser/super_catfacts 
cd super_catfacts

# Create a config.yml in the working directory, filling in the required env variables (see example below)
vi config.yml 

# Build the container - which will bake in your config file 
./build.sh 

# Tag the resulting image for deployment 
docker tag <image-id> gcr.io/super-catfacts/catfacts:v1

# Push the image 
docker push gcr.io/super-catfacts/catfacts:v1 

# Update k8s/deployment.yaml so that the image field points to the image tag you just pushed
kubectl apply -f k8s/deployment.yaml
```
