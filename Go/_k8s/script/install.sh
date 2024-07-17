masternode="<hostname-of-the-master-node>"
controllernode1="<hostname-of-the-controller-node-1>"
controllernode2="<hostname-of-the-controller-node-2>"
gpuworker1="<hostname-of-the-gpu-node-1>"
gpuworker2="<hostname-of-the-gpu-node-2>"
gpuworker3="<hostname-of-the-gpu-node-3>"
gpuworker4="<hostname-of-the-gpu-node-4>"

kubectl label node $controllernode1 dedicated=controller
kubectl taint nodes $controllernode1 node-role.kubernetes.io/controller:NoSchedule
kubectl label node $controllernode2 dedicated=controller
kubectl taint nodes $controllernode2 node-role.kubernetes.io/controller:NoSchedule

kubectl label node $masternode dedicated=master
kubectl label node $gpuworker1 accelerator=nvidia-gpu
kubectl label node $gpuworker2 accelerator=nvidia-gpu
kubectl label node $gpuworker3 accelerator=nvidia-gpu
kubectl label node $gpuworker4 accelerator=nvidia-gpu


kubectl apply -f ./deploy/framework/crd/stateful.yaml
kubectl apply -f ./deploy/framework/crd/stateless.yaml

kubectl apply -f ./deploy/framework/metrics-server.yaml
kubectl apply -f ./deploy/framework/register-gpu-info.yaml
kubectl apply -f ./deploy/framework/priority-class.yaml

sleep 15s

istioctl install -f ./deploy/TestScenario/istio_operator.yml
kubectl label namespace default istio-injection=enabled

#Make sure register-gpu-info.yaml is running before installing mps-server
kubectl apply -f ./deploy/framework/mps-server.yaml

#Scheculer for cluster without istio
#kubectl apply -f ./deploy/framework/scheduler.yaml
#Scheculer for cluster with istio
istioctl kube-inject -f ./deploy/framework/scheduler.yaml | kubectl apply -f -

kubectl apply -f ./deploy/framework/data-manager.yaml
kubectl apply -f ./deploy/framework/resource-monitor.yaml
kubectl apply -f ./deploy/framework/controller.yaml
kubectl apply -f ./deploy/framework/router.yaml

#For cluster with istio
kubectl apply -f ./deploy/TestScenario/istio-gateway.yaml
#For cluster without istio
#kubectl apply -f ./deploy/TestScenario/ingress-controller.yaml
#kubectl apply -f ./deploy/TestScenario/ingress.yaml

sleep 5s

kubectl apply -f ./deploy/observability/grafana
kubectl apply -f ./deploy/observability/loki
kubectl apply -f ./deploy/observability/promtail
#Optional
#kubectl apply -f ./deploy/observability/kibana

#Update kube-scheduler cluster-role
#KUBE_EDITOR=nano kubectl edit clusterrole system:kube-scheduler