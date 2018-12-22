# vegeta-operator

Distrubuted vegeta operator for kubernetes. CURRENTLY under active dev.

https://github.com/tsenart/vegeta

![Vegeta](https://media.giphy.com/media/do0xfdH2Ssh56/giphy.gif)

# RFC:
I have implemented a first design architecture. If you are interested, look at https://github.com/MalloZup/vegeta-operator/issues/3

# Quickstart:

0) Install the CRDs into the cluster

make install

1) Run the command locally against the remote cluster.

make run

2) In a new terminal - create an instance and expect the Controller to pick it up

kubectl apply -f config/samples/*.yaml


# Development:

For usefull infos, look at : 

1) https://github.com/MalloZup/k8s-crd-controller-the-hard-way
and kubebuilder page.

Deploy for k8s:

I use personally kind https://github.com/kubernetes-sigs/kind for having a k8s cluster in container.

# Extra

This project use `kubebuilder`. https://github.com/kubernetes-sigs/kubebuilder
