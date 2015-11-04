# Kubernetes Is <br>For Monoliths Too
##### KubeCon 2015, Nov 10, 2015

***

Steve Reed (@_smreed)

---

# About Me

Early Kubernetes adopter:

* v0.5 in GCE ~Nov 2014
* [#3965](https://github.com/kubernetes/kubernetes/issues/3965) "GCE PD Data Loss"
* 12 commits, ~143 lines survive today
* One of the first production workloads

---

# Outline

* Its service discovery can eliminate a lot of configuration. From
  dev/test/staging/prod copies of everything, to port selection nightmares.
* Example: dev/test/staging/prod can be namespaces, and you can otherwise
  "hardcode" service names. Required dependent services are first class
  citizens in your infrastructure, this in my opinion isn't a bad thing to do.
* Even if you have a database you don't want to containerize, it can still
  be a service in Kubernetes, and utilize the same benefits.
* When your configuration absorbs too many details about your surrounding
  infrastructure, you can wrap that complexity up in an ambassador container.
  For example, wrap up smtp auth w/ an open relay. Run twemproxy along-side
  your app.
* Replace/add features using adapters (nginx for ssl termination)
* Show a sample monolithic application configuration in its simplest form.
  Show how Kubernetes still allows for flexibility, app can be run in a 
  dev or prod mode w/ no change. Walk through changes made to mattermost,
  for example.

<!--
misc notes

at some point, after adapters/ambassadors/sidecars are deployed you
now have a "monolithic" "Modular Container"
-->

---

# Motivation

![k8s-vs-microservices](./assets/k8s-vs-microservices.png)
<small>Google Searches For "Kubernetes" and "Microservices"</small>

* Everybody should be thinking of Kubernetes
* Microservices are not for everybody
* Monolithic applications can benefit
* Kubernetes is not (just) "for" microservices

<!--
* (?) Kubernetes Isn't just for microservices
* There are reasons to move your monolith into Kubernetes
-->

---

# Monoliths <br>These Days

* Application deployed as a single executable or package.
* An amalgam of orthogonal concerns.
* May still utilize many other external services (DB, cache, SMTP, etc)

Often, they are not very cohesive pieces of software.

<!--
* Briefly, What's a monolith? 
-->

---

# Step #1: Containerize

Get it in Kubernetes.

* Expose ports
* Mount filesystems:
  * configuration
  * persistent storage
  * logs

TODO: snippet of monolith Dockerfile?

<!--
Brendan Burns' "The Distributed System Toolkit" at Dockercon SF 2015
-->

---

# Step #1: Containerize

<pre><code class="dockerfile">FROM ubuntu:14.04

EXPOSE 8065
VOLUME /mattermost/data
# TODO logs

# Install

ADD config.json /mattermost/config/

WORKDIR /mattermost/bin
ENTRYPOINT /mattermost/bin/platform</code></pre>

---

# Step #1: Containerize

1. `docker build -t my-monolith .`
2. `docker push my-monolith`
3. <code>kubectl run my-monolith \<br>--image=my-monolith</code>

<!--
even if your container is the only thing running in the cluster, you still
get the advantages of the kubernetes substrate
-->

---

# Step #2: <br>Services

* Service for every external dependency
  * databases
  * webhooks, APIs
* treat these as "prereqs"

These services must exist, hardcode them!

Kubernetes Service becomes the "interface" to these external dependencies.

Interfaces allow decoupling, help reason about cohesion of system components.


---

# Step #2: <br>More Services

* Convert libraries to services
  * As helpers in same pod
  * Or as "Services"

Some libraries exist in monolithic apps just to enable interaction with an
external service. Often times, details of these systems leak into your
monolith in the form of boilerplate code, or configuration that needs to be
injected into the library.

Examples:

* smtp: host, port, auth credentials, supported encryption
* cache shards: names, how many, how are keys hashed
* TLS configuration: host certs, client certs

This increases coupling, decreases cohesion.

---

# Helmsman <br>Of The Titanic

The metaphor doesn't entirely fit, but it sounds cool.

* `Services` elevate your external dependencies to 1st class status
* 

Can Kubernetes help with

* day to day development
* staging and testing

---

# Title

A monolithic application with many external dependencies can lead to complicated
development environments.

---

# Spaghetti.Conf (1/n)

Common to have "dev", "test", and "prod" configuration for the same application.

Kubernetes can provide the same flexibility with namespaces.

---

# Spaghetti.Conf (2/n)

```
# prod-app.properties
db.host = database.prod
db.port = 3306
```

---

# Sandbox Slide

<pre><code class="gooooo">package main

import "fmt"

func main() {
  fmt.Println("Hello!")
}</code></pre>

