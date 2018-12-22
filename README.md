# vegeta-operator

Distrubuted vegeta operator for kubernetes. CURRENTLY under active dev.

https://github.com/tsenart/vegeta

![Vegeta](https://media.giphy.com/media/do0xfdH2Ssh56/giphy.gif)

# RFC:
I have implemented a first design architecture. If you are interested, look at https://github.com/MalloZup/vegeta-operator/issues/3

# Quickstart:

1) Deploy vegeta controller in the configured Kubernetes cluster in ~/.kube/config ( you can use kind for dev env.)
`make deploy`

2) Perform a vegeta attack in distributed manner.
`kubectl apply -f config/samples/vegeta_v1beta1_vegeta.yaml`

( NOTE: You will create resources, but it will have **no effect** since the controller is on WIP).

# Development:

For usefull infos, look at : 

1) https://github.com/MalloZup/k8s-crd-controller-the-hard-way
and kubebuilder page.

Deploy for k8s:

I use personally kind https://github.com/kubernetes-sigs/kind for having a k8s cluster in container.

# Extra

This project use `kubebuilder`. https://github.com/kubernetes-sigs/kubebuilder
