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

Note:
"Paid my dues"

K8S is very open OSS, even to small contributors

Honor to speak

---

# Motivation

* Everybody should be considering Kubernetes
![k8s-vs-microservices](./assets/k8s-vs-microservices.png)
<small>Google Searches For <font color=blue>"Kubernetes"</font> and <font color=red>"Microservices"</font></small>
* Monolithic applications can benefit
  * dependencies documentated
  * `kubectl`

Note:
everybody at least consider k8s

k8s and microservices in same breath

monolith? single deployed artifact

monolith? has external infra. dependencies

---

# Helmsman <br>Of The Titanic

The metaphor doesn't entirely fit, but it sounds cool.

* Remove this slide if someone else uses the phrase!

---

# Containerize <br>Then Refactor

1. Containerize
  * single monolithic pod
  * prerequisites "on paper"
2. Refactor
  * not code, but infrastructure
  * extract services from pod where it makes sense
  * ambassadors, adapters, and sidecars 
    simplify app and increase cohesion

---

# Containerize

* Monolithic Pod
  * "Megapod?" "Podolith?"
* Resources?
* Volume mounts where possible
  * configuration
  * persistent storage
  * logs

Note:
containerization analogy: marking class members private so compiler can show you usages

---

# Example: Mattermost

* Slack clone
* Single app with many dependencies
  * PostgreSQL
  * SMTP

---

# Monolithic Pod

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

[Demo](http://monolith/kubecon-2015)

Note:
`kubectl create -f spec/phase1`

`kubectl get pods -w`

`kubectl logs monolith monolith-app`

if necessary

`./local-up.sh`

---

# Monolithic Pod

* Everything's `localhost`
* Dependencies are specified, enforcable
* Now make it better!

---

# Refactor

* Database, SMTP
  * Needs to scale independently
  * Might be used by other pods as they are added
  * Want ability to proxy
  * Want dev environment "mocks"

Note:

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

---

# Hardcode Stuff

* Service hostnames are "guaranteed"
  * `monolith-db.default.svc.cluster.local`
  * `smtp-server.default.svc.cluster.local`
* `localhost` for everything inside the pod

Note:
Kubernetes adds degree of freedom

Your app always needs a "db", K8s manages what that "db" actually is

---

# More Refactoring 

* `nginx` adapter container for ssl-termination
  * and auth too?
* `Secret`s for certs, auth credentials
* `pgpool` as ambassador to `postgres`
* Default ports (`80` for `HTTP`, `443` for `HTTPS`)

[Demo](https://monolith-ssl/kubecon-2015)

Note:
* ip per pod makes port mapping easy

`kubectl create -f spec/phase2`

`kubectl get pods -w`

if necessary

`./local-up.sh`

* use k8s api to create pods for batch work

---

# Thank You

* Don't let a lack of microservices keep you from trying Kubernetes.
* If you want microservices, Kubernetes can help you get started!
* [slides, code, examples](https://github.com/smreed/kubecon-2015)
* [@_smreed](https://twitter.com/_smreed)

