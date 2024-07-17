package controller

import (
	"context"
	"fmt"
	"crd/pkg/crd"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/controller/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/source"
	"time"
)

// app controller
type DSLAppController struct {
	client.Client
	kubeClient kubernetes.Interface
	scheme     *runtime.Scheme
	config     *rest.Config
}

func NewDSLAppController(client client.Client, kubeClient kubernetes.Interface, scheme *runtime.Scheme, config *rest.Config) *DSLAppController {
	return &DSLAppController{
		Client:     client,
		kubeClient: kubeClient,
		scheme:     scheme,
		config:     config,
	}
}

// Reconcile function
func (r *DSLAppController) Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {
	// Get the DSLApp instance
	instance := &crd.DSLApp{}
	err := r.Get(ctx, req.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Object not found
			return reconcile.Result{}, nil
		}
		// Error
		return reconcile.Result{}, err
	}

	// Logic to scale based on minPods, maxPods 
	newReplicas := calculateReplicas(instance.Spec.MinPods, instance.Spec.MaxPods)
	instance.Status.Replicas = newReplicas

	// Update the status of the DSLApp instance
	err = r.Status().Update(ctx, instance)
	if err != nil {
		return reconcile.Result{}, err
	}
	return reconcile.Result{RequeueAfter: time.Minute}, nil
}

// calculates the no of replicas based on minPods and maxPods
func calculateReplicas(minPods, maxPods int32) int32 {
	return (minPods + maxPods) / 2
}


// sets up the controller with the Manager.
func (r *DSLAppController) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&crd.DSLApp{}).
		Complete(r)
}

// informer
func (r *DSLAppController) Informer(ctx context.Context) {
// Informer sets up an informer to watch for DSLApp resources
func (r *DSLAppController) Informer(ctx context.Context) {
	// handle events on DSLApp resources
	handler := func(obj runtime.Object, event event.Event) bool {
		// runtime.Object to a DSLApp instance
		dslApp, ok := obj.(*crd.DSLApp)
		if !ok {
			log.Printf("Expected DSLApp but got %T\n", obj)
			return false
		}

		switch event {
		case event.Added:// added
			log.Printf("[INFO] DSLApp added: %s\n", dslApp.Name)
		case event.Modified:// modified
			log.Printf("[INFO] DSLApp modified: %s\n", dslApp.Name)
		case event.Deleted:// deleted
			log.Printf("[INFO] DSLApp deleted: %s\n", dslApp.Name)
		}

		return false
	}

	// Create informer for DSLApp resources
	informer := cache.NewFilteringInformer(
		r.Client,
		&crd.DSLApp{}, // Type of the resource to watch
		time.Minute,   // Resync period
		cache.ResourceEventHandlerFuncs{
			AddFunc:    handler,
			UpdateFunc: func(oldObj, newObj interface{}) { handler(newObj, event.Modified) },
			DeleteFunc: func(obj interface{}) { handler(obj, event.Deleted) },
		},
		func(obj client.Object) bool {
			return true // Filter function, always return true to watch all DSLApp instances
		},
	)

	// Start informer
	go informer.Run(ctx.Done())
}

}

// Listener
func (r *DSLAppController) Listener(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case event := <-r.recorder.Events():
			switch event.Type {
			case corev1.EventTypeNormal:
				log.Printf("[INFO] Normal event: %s\n", event.Reason)
			case corev1.EventTypeWarning:
				log.Printf("[INFO] Warning event: %s\n", event.Reason)
			}
		}
	}
}


