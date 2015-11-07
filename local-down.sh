#!/bin/sh
docker rm -f kube-proxy
docker rm -f kubelet
docker rm -f etcd
