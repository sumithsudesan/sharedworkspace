package webhook

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	admissionv1 "k8s.io/api/admission/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// Handles validation and mutation of DSLApp resources
type DSLAppValidator struct {
	Client  kubernetes.Interface
	Decoder *admission.Decoder
}

// Validates and mutates DSLApp resources
func (v *DSLAppValidator) Handle(ctx context.Context, req admission.Request) admission.Response {
	dslApp := &v1.DSLApp{}

	err := v.Decoder.Decode(req, dslApp)
	if err != nil {
		return admission.Errored(http.StatusBadRequest, err)
	}

	var errs field.ErrorList

	switch req.Operation {
	case admissionv1.Create:
		errs = v.validateCreate(dslApp)
	case admissionv1.Update:
		errs = v.validateUpdate(dslApp)
	}

	if len(errs) > 0 {
		return admission.ValidationResponse(false, strings.Join(errs.ToAggregate().Errors(), ", "))
	}

	// DSLApp resource
	v.mutate(dslApp)

	marshaledDSLApp, err := json.Marshal(dslApp)
	if err != nil {
		return admission.Errored(http.StatusInternalServerError, err)
	}

	return admission.PatchResponseFromRaw(req.Object.Raw, marshaledDSLApp)
}

// Rregisters the webhook with the Kubernetes server
func SetupWebhook(mgr *admission.WebhookServer, cfg *rest.Config) {
	wh := &DSLAppValidator{}

	mgr.Register("/v1/dslapp", &admission.Webhook{
		Handler: wh,
	})

	mgr.Register("/v2/dslapp", &admission.Webhook{
		Handler: wh,
	})
}

// Validates the creation of DSLApp resources
func (v *DSLAppValidator) validateCreate(dslApp *v1.DSLApp) field.ErrorList {
	var errs field.ErrorList
	if dslApp.Spec.MinPods < 1 {
		errs = append(errs, field.Invalid(field.NewPath("spec", "minPods"), dslApp.Spec.MinPods, "must be at least 1"))
	}
	return errs
}

// Validates the update of DSLApp resources
func (v *DSLAppValidator) validateUpdate(dslApp *v1.DSLApp) field.ErrorList {
	var errs field.ErrorList
	if dslApp.Spec.MaxPods < dslApp.Spec.MinPods {
		errs = append(errs, field.Invalid(field.NewPath("spec", "maxPods"), dslApp.Spec.MaxPods, "must be greater than minPods"))
	}
	return errs
}

// Mutates the DSLApp resource
func (v *DSLAppValidator) mutate(dslApp *v1.DSLApp) {
	// CPU limit is set
	if dslApp.Spec.CPULimit == 0 {
		dslApp.Spec.CPULimit = 1
	}
}
