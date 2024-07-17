package main

import (
	"flag"
	"log"

	"crd/pkg/controller"

	"github.com/operator-framework/operator-sdk/pkg/k8sutil"
	"github.com/operator-framework/operator-sdk/pkg/leader"
	"github.com/operator-framework/operator-sdk/pkg/manager"
	"github.com/operator-framework/operator-sdk/pkg/sdk"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

func main() {
	namespace, err := k8sutil.GetWatchNamespace()
	if err != nil {
		log.Fatalf("Failed to get watch namespace: %v", err)
	}

	// Setup the Operator SDK manager
	mgr, err := manager.New(config.Impl{
		Namespace: namespace,
	})
	if err != nil {
		log.Fatalf("Failed to create manager: %v", err)
	}

	// Setup the DSLApp controller
	reconciler := &controller.DSLAppReconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),
	}

	err = reconciler.SetupWithManager(mgr)
	if err != nil {
		log.Fatalf("Failed to setup controller: %v", err)
	}

	// Start the Operator SDK manager
	err = mgr.Start(signals.SetupSignalHandler())
	if err != nil {
		log.Fatalf("Failed to start manager: %v", err)
	}
}
