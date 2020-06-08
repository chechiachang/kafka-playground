#!/bin/bash
#
# https://github.com/helm/charts/tree/master/incubator/kafka

#HELM_NAME=kafka
HELM_NAME=kafka-1

helm repo add incubator http://storage.googleapis.com/kubernetes-charts-incubator

# Stable: chart version: redis-ha-3.6.1	app version: 5.0.1
helm upgrade --install ${HELM_NAME} incubator/kafka --version 0.16.2 -f values-staging.yaml
