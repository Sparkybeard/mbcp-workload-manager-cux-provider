

package v1alpha1

import (
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// SolutionParameters are the configurable fields of a Solution.
type SolutionParameters struct {
	ConfigurableField string `json:"configurableField"`
}

// SolutionObservation are the observable fields of a Solution.
type SolutionObservation struct {
	ObservableField string `json:"observableField,omitempty"`
}

// An SolutionSpec defines the desired state of a Solution.
type SolutionSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       SolutionParameters `json:"forProvider"`
}

// An SolutionStatus represents the observed state of a Solution.
type SolutionStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          SolutionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A Solution is an example API type.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,template}
type Solution struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SolutionSpec   `json:"spec"`
	Status SolutionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SolutionList contains a list of Solution
type SolutionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Solution `json:"items"`
}

// Solution type metadata.
var (
	SolutionKind             = reflect.TypeOf(Solution{}).Name()
	SolutionGroupKind        = schema.GroupKind{Group: Group, Kind: SolutionKind}.String()
	SolutionKindAPIVersion   = SolutionKind + "." + SchemeGroupVersion.String()
	SolutionGroupVersionKind = SchemeGroupVersion.WithKind(SolutionKind)
)