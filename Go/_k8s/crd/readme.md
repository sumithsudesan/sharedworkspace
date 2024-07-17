CRD and Operator for DSLApp
Overview
This repository contains the custom resource definition (CRD) and an operator implemented in Go that manages resources of type DSLApp in Kubernetes clusters. The DSLApp resource allows users to define application configurations with specific attributes such as app, minPods, maxPods,cpuLimit, gpuLimit, freePods,  and memoryLimit.

The operator ensures that DSLApp resources are reconciled according to the defined specifications, scaling the application pods based on minPods and maxPods attributes.

Operator
    The operator manages the lifecycle of DSLApp resources.

    Reconcile Logic: Implements the reconcile function to ensure DSLApp resources are in the desired state, scaling pods based on minPods and maxPods.

RBAC: Uses RBAC (rbac.yaml) to define roles and permissions required for the operator to manage DSLApp resources within Kubernetes.

Note:
Steps to Generate CRD Types:
1. Install controller-gen
   go get sigs.k8s.io/controller-tools/cmd/controller-gen@v0.7.0
2. Define CRD: Create a YAML file that defines your Custom Resource Definition
3. Generate CRD Types: controller-gen crd paths=./... output:crd:pkg/crd
