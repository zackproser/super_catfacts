# Super CatFacts Service
A ridiculously over-engineered CatFacts prank service written in Golang and deployed via Kubernetes and Google Cloud

# Prank at a glance

As a service administrator, you text the phone number of your prank target to the service. The service then launches an attack on your target - and allows you to monitor and stop it at will.

Your target is then entered into CatFacts hell. They'll receive text messages at the interval you've configured (defaults to 30 seconds) and if they text the number back they'll get one of several infuriating messages about their command not being recognized or receiving additional CatFacts for free.

![Commanding your catfacts service](/static/img/initiate.png)

# Pre-requisites
* A funded Twilio account
* A valid Twilio phone number
* A working Google Cloud Kubernetes cluster
* A working Golang installation
* A valid domain name for mapping to the service

# Getting started
```
# Clone the repo and build the binary
git clone github.com/zackproser/super_catfacts

# Run tests
go test -v ./...

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
   apikey: e8e6fedwyc73tdjsjsdhfoiefua03w37dh
   sid: AC4djwiyf5dc871236ffce9001c8addae
   messageIntervalSeconds: 30
```

# CatFacts configuration

Here are the configuration variables at a glance

| Name  | Type  | Description  | Required  |
|---|---|---|---|
| server.port  | int | The port the CatFacts service will listen on |  :heavy_check_mark: |
|  server.admins | array of strings  |  These are the phone numbers of your service's administrators. They will have total control over the service | :heavy_check_mark:  |
|  server.catfactsuser | string  | Basic auth username, used by Twilio to access your service | :heavy_check_mark:  |
| server.catfactspassword | string | Basic auth password, used by Twilio to access your service | :heavy_check_mark: |
| twilio.number | string | The Twilio number you are mapping to your CatFacts service | :heavy_check_mark: |
| twilio.apikey | string | Your API Key (also called an auth token) from your Twilio dashboard | :heavy_check_mark: |
| twilio.sid | string  | Your SID from your Twilio dashboard | :heavy_check_mark: |
| twilio.messageIntervalSeconds | int | The number of seconds to pause between sending text messages when attacking a target | :x: |

# Devops
Super CatFacts is a web service designed to be deployed via Kubernetes and to integrate with Twilio via user-supplied Twilio account credentials. From the perspective of Kubernetes, Super Catfacts is a deployment of a Catfacts Docker image and a k8s service that exposes it such that Twilio can interact with it. 

You will also need a domain name to map to your Kubernetes service. Once you have your CatFacts k8s service running successfully, point a DNS A record at the ipv4 address of your loadBalancer and then update config.yml such that FQDN is set to your domain name. Rebuild the image via ```build.sh```, tag and deploy it and then configure your Twilio webhooks to point to your domain name (including basic auth credentials). 

**E.G:** If your *domain* is catfacts.com, and your *catfactsusername* is furry and your *catfactspassword* is furB4l1, then the URL you'd enter in your Twilio dashboard as your webhooks would like something like: 

```https://furry:furB4l1@


 your config.yml's ```server.fqdn``` field should be set to ```catfacts.com```, and the full URL including 

# Security and authentication

To prevent abuse of your service, Super CatFacts requires HTTP Basic Auth by default. In your config, you define a Basic Auth username and password, and Super CatFacts automatically requires these credentials on all API endpoints. 

When configuring your URLs in the Twilio dashboard, you must supply all webhook URLs correctly formatted with the HTTP Basic Auth ```username:password```syntax, e.g: ```https://twiliousername:su3e97r7ehdgdh@mycatfactsservice.com```

# Controlling your CatFacts deployment
Management of the service (starting and stopping attacks, checking system status) is done via SMS messages to your configured Twilio phone number. You must be an admin (as configured via the Server.Admins node in config.yml) to control the service.

# Start an attack
Text a US phone number to your Twilio CatFacts service number. The service will respond confirming the target is under attack

# Get attack status
Text "Status" to your Twilio CatFacts service number. The service will respond with a read out of all current attacks, their message counts, and start times.

# Stop an attack
Text a US phone number that is already under attack to your Twilio CatFacts service number. The service will respond confirming the target is no longer being attacked
