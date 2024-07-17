package crd

import (
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/apimachinery/pkg/runtime"
)

// DslappSpec defines the desired state of Dslapp
type DslappSpec struct {
    App            string `json:"app,omitempty"`
    MinPods        int32  `json:"minPods,omitempty"`
    MaxPods        int32  `json:"maxPods,omitempty"`
    PodTemplate    string `json:"podTemplate,omitempty"`
    ServiceTemplate string `json:"serviceTemplate,omitempty"`
    CPULimit       int32  `json:"cpuLimit,omitempty"`
    GPULimit       int32  `json:"gpuLimit,omitempty"`
    FreePods       int32  `json:"freePods,omitempty"`
    HealthAPI      string `json:"healthAPI,omitempty"`
    MemoryLimit    string `json:"memoryLimit,omitempty"`
}

// DslappStatus defines the observed state of Dslapp
type DslappStatus struct {
    // ObservedGeneration is the most recent generation observed for this Dslapp. It corresponds to the CRD's metadata.generation field.
    ObservedGeneration int64 `json:"observedGeneration,omitempty"`

    // Conditions represent the latest available observations of a resource's current state.
    Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="App",type="string",JSONPath=".spec.app",description="The application name"
// +kubebuilder:printcolumn:name="MinPods",type="integer",JSONPath=".spec.minPods",description="Minimum number of pods"
// ... additional annotations

// Dslapp is the Schema for the dslapps API
type Dslapp struct {
    metav1.TypeMeta   `json:",inline"`
    metav1.ObjectMeta `json:"metadata,omitempty"`

    Spec   DslappSpec   `json:"spec,omitempty"`
    Status DslappStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DslappList contains a list of Dslapp
type DslappList struct {
    metav1.TypeMeta `json:",inline"`
    metav1.ListMeta `json:"metadata,omitempty"`
    Items           []Dslapp `json:"items"`
}

func init() {
    SchemeBuilder.Register(&Dslapp{}, &DslappList{})
}
