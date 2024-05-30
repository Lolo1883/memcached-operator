package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ImmortalDBSpec defines the desired state of ImmortalDB
type ImmortalDBSpec struct {
	Image    string `json:"image,omitempty"`
	Replicas int32  `json:"replicas,omitempty"`
}

// ImmortalDBStatus defines the observed state of ImmortalDB
type ImmortalDBStatus struct {
	Nodes []string `json:"nodes,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ImmortalDB is the Schema for the immortaldbs API
type ImmortalDB struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ImmortalDBSpec   `json:"spec,omitempty"`
	Status ImmortalDBStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ImmortalDBList contains a list of ImmortalDB
type ImmortalDBList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ImmortalDB `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ImmortalDB{}, &ImmortalDBList{})
}
