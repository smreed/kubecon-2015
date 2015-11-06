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

* Briefly, what is a monolith?
* Phases:
  * Containerize: monolithic pod
  * Decouple: Pull services out of pod where you want to scale differently
  * Refactor: Increase modularity (introduce ambassadors/adapters/sidecars)
    in order to simplify the monolith and increase cohesion
* Containerize: app, database, configuration, logs, everything in a single
  pod. Can it work? Is it worth it?
* Decouple: Kubernetes allows you to define the interface that your
  infrastructure must provide for your app.
* Pod and Service specs define infrastructure requirements. Any environment that
  implements the "spec" should be able to run, in some fashion, your application.
  This is true for all apps in Kubernetes, including monolithic ones.
* Development environment can implement the infrastructure interface with
  in-memory databases, temporary filesystems, self-signed certs, etc.
* One way to go about this refactoring is to try to optimize for simplicity
  of the development environment.
* Example: dummy smtp server
* Refactor: Move the configubility of interactions with external services
  into modular containers.
* Example: sharded caches. simplify configuration to the simplest case and use
  an ambassador like twemproxy.
* Summary: Don't let a lack of microservices keep you from trying Kubernetes.
* Summary: Because Kubernetes can do something doesn't mean it must.
* Summary: Everything described here is possible with other tools. But Kubernetes
  gives us a language for specifying our infrastructure needs and component
  interaction. We can use this language to monitor and verify deployment
  environments.


<!--
misc notes

Maybe I should just write my own fake monolith?

at some point, after adapters/ambassadors/sidecars are deployed you
now have a "monolithic" "Modular Container"
-->

---

# Motivation

![k8s-vs-microservices](./assets/k8s-vs-microservices.png)
<small>Google Searches For "Kubernetes" and "Microservices"</small>

* Everybody should be considering Kubernetes
* Monolithic applications can benefit
* Kubernetes is not "for" microservices

<!--
* (?) Kubernetes Isn't just for microservices
* There are reasons to move your monolith into Kubernetes
-->

---

# Monoliths <br>These Days

* Application deployed as a single executable or package.
* May still utilize many other external services (DB, cache, SMTP, etc)

Often, they are not very cohesive pieces of software.

<!--
* An amalgam of orthogonal concerns.

Come back to the point about cohesion. Once in K8S, you have a way to
start to easily externalize a lot of the bloat and improve cohesion.
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

