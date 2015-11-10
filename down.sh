#!/bin/sh
kubectl config use-context gke_smreed-kubecon-2015_us-central1-c_monolith
docker rm -f kube-proxy
docker rm -f kubelet
docker rm -f etcd
