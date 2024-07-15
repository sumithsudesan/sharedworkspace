kubectl delete -f ./deploy/framework/crd/stateful.yaml
kubectl delete -f ./deploy/framework/crd/stateless.yaml

kubectl delete -f ./deploy/framework/metrics-server.yaml
kubectl delete -f ./deploy/framework/register-gpu-info.yaml
kubectl delete -f ./deploy/framework/priority-class.yaml


kubectl delete -f ./deploy/framework/mps-server.yaml
kubectl delete -f ./deploy/framework/scheduler.yaml
kubectl delete -f ./deploy/framework/data-manager.yaml
kubectl delete -f ./deploy/framework/resource-monitor.yaml
kubectl delete -f ./deploy/framework/controller.yaml
kubectl delete -f ./deploy/framework/router.yaml

#kubectl delete -f ./deploy/TestScenario/ingress-controller.yaml
#kubectl delete -f ./deploy/TestScenario/ingress.yaml
kubectl delete -f ./deploy/TestScenario/istio-gateway.yaml

kubectl delete -f ./deploy/observability/grafana
kubectl delete -f ./deploy/observability/loki
kubectl delete -f ./deploy/observability/promtail

#kubectl delete -f ./deploy/observability/kibana
#kubectl delete pvc data-es-cluster-0