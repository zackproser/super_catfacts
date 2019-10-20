# Super CatFacts Service
A ridiculously over-engineered CatFacts prank service written in Golang and deployed via Kubernetes and Google Cloud.

[![Build Status](https://travis-ci.com/zackproser/super_catfacts.svg?branch=master "Travis CI Status")](https://travis-ci.com/zackproser/super_catfacts)
[![Go Report Card](https://goreportcard.com/badge/github.com/zackproser/super_catfacts)](https://goreportcard.com/report/github.com/zackproser/super_catfacts) [![GoDoc](https://godoc.org/github.com/zackproser/super_catfacts?status.svg)](https://godoc.org/github.com/zackproser/super_catfacts/cmd)

# Prank at a glance

As a service administrator, you text the phone number of your prank target to the service. The service then launches an attack on your target - and allows you to monitor and stop it at will.

![Commanding your catfacts service](/static/img/initiate.png)

This service prioritizes expediency and precision of command, because successful offensives are founded upon the core principles of **speed**, **violence**, and **momentum**.

Your target is thusly entered into CatFacts hell. They'll receive text messages at the interval you've configured (defaults to 30 seconds) and if they text the number back they'll get one of several infuriating messages about their command not being recognized or receiving additional CatFacts for free.

![Receiving CatFacts](/static/img/2nd.png)

![Texting back](/static/img/3rd.png)

If they call the number, they get an additionally hellish phone tree experience, complete with piercing angry meowing sounds. [Watch a recording of what the phone tree is like here.](https://www.youtube.com/watch?v=_Tx5LtcOIgg)

# Pre-requisites
* A funded Twilio account
* A valid Twilio phone number
* A working Google Cloud Kubernetes cluster
* A working Golang installation
* A valid domain name for mapping to the service

# Quickstart
```
# Clone the repo and build the binary
git clone github.com/zackproser/super_catfacts

cd super_catfacts

# Run tests
go test -v ./...

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

# Deployment
Out of the box, this service runs smoothly via Kubernetes on Google Cloud, though it could be easily modified to run in any other Kubernetes cluster. That said, here are the high level steps to deploying this via Google Cloud, which are also spelled out in the quickstart above:

1. [Complete the steps to get a working Google Cloud Kubernetes cluster and authenticate with it via ```gcloud``` and ```kubectl```](https://cloud.google.com/kubernetes-engine/docs/quickstart)
2. Check out the repository
3. Ensure tests pass and that go build returns without error
4. Create a config.yml file in the root of the working directory, and configure it as explained above
5. Run the build.sh script to get a Docker image with your specific config file baked into the container
6. [Tag the resulting image for Google Container registry and push it](https://cloud.google.com/container-registry/docs/pushing-and-pulling)
7. Update your deployment.yaml so that the ```spec.containers.image``` node points at your pushed image
8. ```kubectl apply -f k8s/deployment.yaml```
9. Ensure the container came up and is healthy without restarts ```kubectl get po,svc```
10. Monitor logs with ```kubectl logs -f <your-deployment-id-from-previous-step>```

# API
The service also exposes an API for scripting against or in case you need more flexibility in integrating it with existing code:

| Method | Route | Description |
|---|---|---|
| GET | / | Returns a basic healthcheck confirming the service is up |
| GET | /attacks | Returns currently running attacks |
| POST | /attacks | Creates a new attack. A form field named **target** containing a valid US phone number is **required** |
| DELETE | /attacks/:id | Stops the attack referenced by the **id** parameter. Call GET /attacks to see attacks and their Ids |

Note that the API is also secured via Basic Auth, so you will have to supply the correctly formed url in order to call it, e.g.

```curl -X POST https://furball:397degfeug@mysupercatfacts.com/attacks -d target=5103768999```

# Devops
Super CatFacts is a web service designed to be deployed via Kubernetes and to integrate with Twilio via user-supplied Twilio account credentials. From the perspective of Kubernetes, Super Catfacts is a deployment of a Catfacts Docker image and a k8s service that exposes it to the public internet such that Twilio can interact with it. Note that the service is locked down via Basic Auth so that only Twilio and server administrators should be able to access it.

You will also need a domain name to map to your Kubernetes service. Once you have your CatFacts k8s service running successfully, point a DNS A record at the ipv4 address of your loadBalancer and then update config.yml such that FQDN is set to your domain name. Rebuild the image via ```build.sh```, tag and deploy it and then configure your Twilio webhooks to point to your domain name (including basic auth credentials).

**E.G:** If your *domain* is mysupercatfacts.com, and your *catfactsusername* is **furball** and your *catfactspassword* is **397degfeug**, then the URL you'd enter in your Twilio dashboard as your webhooks would be:

```https://furball:397degfeug@mysupercatfacts.com/call/receive```

![Example Twilio configuration](/static/img/twilio-configuration.jpg)

This allows Twilio to reach your service correctly and retrieve the TwiMl it renders to control the phone tree experience, etc.

In order to ensure your service renders the correct URLs within TwiMl, your config.yml's ```server.fqdn``` field should be set to ```mysupercatfacts.com```, in this example.

# Security and authentication

To prevent abuse of your service, Super CatFacts requires HTTP Basic Auth by default. In your config, you define a Basic Auth username and password, and Super CatFacts automatically requires these credentials on all API endpoints.

When configuring your URLs in the Twilio dashboard, you must supply all webhook URLs correctly formatted with the HTTP Basic Auth ```username:password```syntax, e.g: ```https://twiliousername:su3e97r7ehdgdh@mycatfactsservice.com```

# Controlling your CatFacts deployment
Management of the service (starting and stopping attacks, checking system status) is done via SMS messages to your configured Twilio phone number. You must be an admin (as configured via the Server.Admins node in config.yml) to control the service.

# Start an attack
Text a US phone number to your Twilio CatFacts service number. The service will respond confirming the target is under attack

![Commanding your catfacts service](/static/img/initiate.png)

# Get attack status
Text "Status" to your Twilio CatFacts service number. The service will respond with a read out of all current attacks, their message counts, and start times.

![Getting service status](/static/img/status.png)

# Stop an attack
Text a US phone number that is already under attack to your Twilio CatFacts service number. The service will respond confirming the target is no longer being attacked

![Stopping an attack](/static/img/terminate.png)
