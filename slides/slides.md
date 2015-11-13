# Kubernetes Is <br>For Monoliths Too
##### KubeCon 2015, Nov 10, 2015

***

Steve Reed (@_smreed)

---

# About Me

Early Kubernetes adopter:

* v0.5 in GCE ~Nov 2014
* One of the first production workloads
* 12 commits, ~143 lines survive today
* [#3965](https://github.com/kubernetes/kubernetes/issues/3965) "GCE PD Data Loss"
* I have nothing to sell you today :)

Note:
features are but one way to be a contributor to open source

K8S is very open project, even to small contributors
like myself

Honor to speak

not here to sell you anything :)

---

# Motivation

* Everybody should be considering Kubernetes
![k8s-vs-microservices](./assets/k8s-vs-microservices.png)
<small>Google Searches For <font color=blue>"Kubernetes"</font> and <font color=red>"Microservices"</font></small>
* Not just for microservices

Note:
everybody at least consider k8s

k8s and microservices in same breath

monolith? single deployed artifact

monolith? has external infra. dependencies

---

# Helmsman <br>Of The Titanic

* Probably don't do this.
* Remove this slide if someone else uses the title.

Note:
mostly mental exercise

this talk: actually about the parallels
between basic programming concepts, and kubernetes primitives

also analogy for migrating to containers/kubernetes

---

# Object Oriented <br>Infrastructure?

* Kubernetes Primitives
  * they describe your application
  * they also describe requirements

Pods, Services, etc are the "interface" between your infrastructure and application.

Note:
benefit: the k8s primitives

k8s primitives: yes, it's a language for describing your application's interface

BUT: also language for describing required infrastructure interface

monolith is just one single box with needs

my monolith needs persistent storage

my monolith needs a database, etc

---

# Containerize...

* Monolithic Pod
  * "Megapod?" "Podolith?"
* Volume mounts where possible
  * configuration
  * persistent storage
  * logs
* Secrets
* Resource Limits?

Note:
let's get started

your monolithic app interface is now "on paper"

so are it's prereqs

containerization analogy: marking class members private so compiler can show you usages

---

# Monolithic Pod

* Everything's `localhost`
* Dependencies are specified, enforcable
* Now make it better!

[Demo](http://monolith/kubecon-2015)

Note:
show monolithic pod yaml

---

# Refactor

* Database, SMTP
  * Needs to scale independently
  * Might be used by other pods as they are added
  * Want ability to proxy
  * Want dev environment "mocks"

Note:
not code, but infrastructure

---

# Hardcode Stuff!

* Service hostnames are "guaranteed"
  * `monolith-db.default.svc.cluster.local`
  * `smtp-server.default.svc.cluster.local`
* `localhost` for everything inside the pod
* storage paths
  * `/var/log/monolith`

Note:
Your app panics if the filesystem is absent or readonly, why
not do the same for services?

Kubernetes adds degree of freedom

Your app always needs a "db", K8s manages what that "db" actually is

---

# More Ideas 

* "Services without selectors"
* reverse proxy "adapter" container
  * ssl-termination
  * authentication
* `pgpool` "ambassador" to `postgres`
* "semantic pipelines"
* `Secret`s for certs, auth credentials
* Default ports (`80` for `HTTP`, `443` for `HTTPS`)
* replace `crond` with sidecars, `Job`s

Note:
ambassador, adapter, sidecar

can add functionality without increasing size of app

---

# Bonus Content?

* "dueling" replication controllers

---

# Thank You

* Just try Kubernetes.
* If you want to refactor your app, Kubernetes can help you get started!
* [slides, code, examples](https://github.com/smreed/kubecon-2015)
* [@_smreed](https://twitter.com/_smreed)

