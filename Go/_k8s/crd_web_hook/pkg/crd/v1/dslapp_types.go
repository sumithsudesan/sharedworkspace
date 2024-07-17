// +groupName=test.com
// +versionName=v1

package v1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// DSLAppSpec defines the desired state of DSLApp
type DSLAppSpec struct {
	App            string `json:"app"`
	MinPods        int32  `json:"minPods"`
	MaxPods        int32  `json:"maxPods"`
	PodTemplate    string `json:"podTemplate"`
	ServiceTemplate string `json:"serviceTemplate"`
	CPULimit       int32  `json:"cpuLimit"`
	GPULimit       int32  `json:"gpuLimit"`
	FreePods       int32  `json:"freePods"`
	HealthAPI      string `json:"healthAPI"`
	MemoryLimit    string `json:"memoryLimit"`
}

// DSLAppStatus defines the observed state of DSLApp
type DSLAppStatus struct {
	Replicas int32 `json:"replicas"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="App",type="string",JSONPath=".spec.app"
// +kubebuilder:printcolumn:name="MinPods",type="integer",JSONPath=".spec.minPods"
// +kubebuilder:printcolumn:name="MaxPods",type="integer",JSONPath=".spec.maxPods"
// +kubebuilder:printcolumn:name="Replicas",type="integer",JSONPath=".status.replicas"

// DSLApp is the Schema for the dslapps API
type DSLApp struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DSLAppSpec   `json:"spec,omitempty"`
	Status DSLAppStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DSLAppList contains a list of DSLApp
type DSLAppList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DSLApp `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DSLApp{}, &DSLAppList{})
}
