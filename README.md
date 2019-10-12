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
# Example config 

```
server:
  port: 3000
  admins:
    - "5143277224"
    - "9613262719"
  catfactsuser: furtastic
  catfactspassword: ZywsrdsgSY3241254sSweop
  fqdn: example.com
twilio:
   number: "+13402937949"
   apikey: 812029d43c5957w38e09459c859bf5
   sid: AC4djwiyf5dc871236ffce9001c8addae
   messageIntervalSeconds: 30
``` 

# CatFacts configuration

Here are the configuration variables at a glance 

| Name  | Type  | Description  | Required  |   |
|---|---|---|---|---|
| server.port  | int | The port the CatFacts service will listen on |  :heavy_check_mark |   |
|   |   |   |   |   |
|   |   |   |   |   |

# Controlling your CatFacts deployment 
Management of the service (starting and stopping attacks, checking system status) is done via SMS messages to your configured Twilio phone number. You must be an admin (as configured via the Server.Admins node in config.yml) to control the service. 

# Start an attack 
Text a US phone number to your Twilio CatFacts service number. The service will respond confirming the target is under attack

# Get attack status 
Text "Status" to your Twilio CatFacts service number. The service will respond with a read out of all current attacks, their message counts, and start times. 

# Stop an attack 
Text a US phone number that is already under attack to your Twilio CatFacts service number. The service will respond confirming the target is no longer being attacked 
