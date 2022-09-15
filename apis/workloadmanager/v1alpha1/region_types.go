

package v1alpha1

import (
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// RegionParameters are the configurable fields of a Region.
type RegionParameters struct {
	ConfigurableField string `json:"configurableField"`
}

// RegionObservation are the observable fields of a Region.
type RegionObservation struct {
	ObservableField string `json:"observableField,omitempty"`
}

// An RegionSpec defines the desired state of a Region.
type RegionSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       RegionParameters `json:"forProvider"`
}

// An RegionStatus represents the observed state of a Region.
type RegionStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          RegionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A Region is an example API type.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,template}
type Region struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RegionSpec   `json:"spec"`
	Status RegionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RegionList contains a list of Region
type RegionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Region `json:"items"`
}

// Region type metadata.
var (
	RegionKind             = reflect.TypeOf(Region{}).Name()
	RegionGroupKind        = schema.GroupKind{Group: Group, Kind: RegionKind}.String()
	RegionKindAPIVersion   = RegionKind + "." + SchemeGroupVersion.String()
	RegionGroupVersionKind = SchemeGroupVersion.WithKind(RegionKind)
)