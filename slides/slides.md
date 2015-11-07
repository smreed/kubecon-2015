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

# Motivation

![k8s-vs-microservices](./assets/k8s-vs-microservices.png)
<small>Google Searches For <font color=blue>"Kubernetes"</font> and <font color=red>"Microservices"</font></small>

* Everybody should be considering Kubernetes
* Monolithic applications can benefit
* Kubernetes is not "for" microservices

---

# Monoliths <br>These Days

* Application deployed as a single executable or package.
* May still utilize many other external services (DB, cache, SMTP, etc)

Often, they are not very cohesive pieces of software.

---

# 3 Phases

1. Containerize
  * single monolithic pod
2. Decouple
  * extract services from pod where it makes sense
  * refactoring your infrastructure
3. Simplify
  * introduce ambassadors, adapters, and sidecars in order to increase cohesion

---

# Phase 1 <br>Containerize

* Monolithic Pod
  * "Megapod?" "Podolith?"
* Volume mounts where possible
  * configuration
  * persistent storage
  * logs

---

# Monolithic Pod

Example!

```
kind: Pod
apiVersion: v1
metadata:
  name: monolith
  labels:
    app: monolith
    phase: one
spec:
  containers:
  - name: monolith-app
    image: "gcr.io/smreed_kubecon_2015/smreed-kubecon-2015-monolith:37a5bbc"
    ports:
    - containerPort: 8065
    volumeMounts:
    - name: monolith-data
      mountPath: /data/monolith
    - name: monolith-logs
      mountPath: /var/log/monolith
  - name: monolith-db
    image: "gcr.io/smreed_kubecon_2015/smreed-kubecon-2015-monolith-db:78f789f"
    env:
    - name: PGDATA
      value: /var/lib/postgresql/data/monolith
    ports:
    - containerPort: 5432
    volumeMounts:
    - name: monolith-db
      mountPath: /var/lib/postgresql/data
  - name: smtp-server
    image: "gcr.io/smreed_kubecon_2015/smreed-kubecon-2015-smtp-dummy:de60fd7"
    ports:
    - containerPort: 25
  volumes:
  - name: monolith-data
    gcePersistentDisk:
      pdName: monolith-data
      fsType: ext4
  - name: monolith-db
    gcePersistentDisk:
      pdName: monolith-db
      fsType: ext4
  - name: monolith-logs
    emptyDir: {}
```

---

# Monolithic Pod

* Everything's `localhost`
* Dependencies are specified, enforcable
* Now make it better!

[Demo](http://monolith/kubecon-2015)

---

# Phase 2: Decouple

* Database, SMTP
  * Needs to scale independently
  * Might be used by other pods as they are added
  * Want ability to proxy
  * Want dev environment "mocks"

---

# Database Pod

```
kind: Pod
apiVersion: v1
metadata:
  name: monolith-db
  labels:
    app: monolith-db
    phase: two
spec:
  containers:
  - name: monolith-db
    image: "gcr.io/smreed_kubecon_2015/smreed-kubecon-2015-monolith-db:78f789f"
    env:
    - name: PGDATA
      value: /var/lib/postgresql/data/monolith
    ports:
    - containerPort: 5432
    volumeMounts:
    - name: monolith-db
      mountPath: /var/lib/postgresql/data
  volumes:
  - name: monolith-db
    gcePersistentDisk:
      pdName: monolith-db
      fsType: ext4
```

---

# SMTP Pod

```
kind: Pod
apiVersion: v1
metadata:
  name: smtp-server
  labels:
    app: smtp-server
    phase: two
spec:
  containers:
  - name: smtp-server
    image: "gcr.io/smreed_kubecon_2015/smreed-kubecon-2015-smtp-dummy:de60fd7"
    ports:
    - containerPort: 25
```

---

# New "Podolith"

```
kind: Pod
apiVersion: v1
metadata:
  name: monolith-app
  labels:
    app: monolith
    phase: two
spec:
  containers:
  - name: monolith-app
    image: "gcr.io/smreed_kubecon_2015/smreed-kubecon-2015-monolith:c01a060"
    ports:
    - containerPort: 8065
    volumeMounts:
    - name: monolith-data
      mountPath: /data/monolith
    - name: monolith-logs
      mountPath: /var/log/monolith
  volumes:
  - name: monolith-data
    gcePersistentDisk:
      pdName: monolith-data
      fsType: ext4
  - name: monolith-logs
    emptyDir: {}
```

* `monolith-db.default.svc.cluster.local`
* `smtp-server.default.svc.cluster.local`

---

# Phase 3: Simplify

TODO : figure that out

[Demo](https://monolith-ssl/kubecon-2015)

---

# Helmsman <br>Of The Titanic

The metaphor doesn't entirely fit, but it sounds cool.

* Remove this slide if someone else uses the phrase!

---

# That's It!

* Don't let a lack of microservices keep you from trying Kubernetes.
* But Kubernetes gives us a language for specifying our infrastructure needs
  and component interaction. We can use this language to monitor and verify
  deployment environments.

---

# Thank You

* [slides, code, examples](https://github.com/smreed/kubecon-2015)

<!--
# Outline

* Briefly, what is a monolith?
* Phases:
  * Containerize: monolithic pod
  * Decouple: Pull services out of pod where you want to scale differently
  * Simplify: Increase modularity (introduce ambassadors/adapters/sidecars)
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
* Simplify: Move the configubility of interactions with external services
  into modular containers.
* Example: sharded caches. simplify configuration to the simplest case and use
  an ambassador like twemproxy.
* Summary: Don't let a lack of microservices keep you from trying Kubernetes.
* Summary: Because Kubernetes can do something doesn't mean it must.
* Summary: Everything described here is possible with other tools. But Kubernetes
  gives us a language for specifying our infrastructure needs and component
  interaction. We can use this language to monitor and verify deployment
  environments.


misc notes

Maybe I should just write my own fake monolith?

at some point, after adapters/ambassadors/sidecars are deployed you
now have a "monolithic" "Modular Container"

Brendan Burns' "The Distributed System Toolkit" at Dockercon SF 2015
-->
