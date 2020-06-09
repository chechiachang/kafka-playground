REPO := git@github.com:strimzi/strimzi-kafka-operator.git
VERSION := 0.18.0

get:
	kpt pkg get ${REPO}/install/cluster-operator@${VERSION} cluster-operator

kpt:
	kpt pkg update cluster-operator@${VERSION}

NAMESPACE := kafka

sed:
	sed -i '' 's!namespace: .*!namespace: $(NAMESPACE)!g' cluster-operator/*RoleBinding*.yaml
	sed -i '' 's!value: my-kafka-project!value: $(NAMESPACE)!g' cluster-operator/050-Deployment-strimzi-cluster-operator.yaml

apply:
	kubectl apply --namespace $(NAMESPACE) -f cluster-operator
