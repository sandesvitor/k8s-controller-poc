FROM alpine

COPY ./k8s-controller-poc /usr/local/bin/k8s-controller-poc
COPY ./kubeconfig.yaml /root/.kube/config


CMD ["k8s-controller-poc", "--kubeconfig", "/root/.kube/config"]

