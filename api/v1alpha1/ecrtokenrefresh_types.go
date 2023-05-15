/*
Copyright 2023 @apanasiuk-el edenlabllc.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ECRTokenRefreshSpec defines the desired state of ECRTokenRefresh
type ECRTokenRefreshSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	ECRRegistry string           `json:"ecrRegistry"`
	Region      string           `json:"region"`
	Frequency   *metav1.Duration `json:"frequency"`
}

// ECRTokenRefreshStatus defines the observed state of ECRTokenRefresh
type ECRTokenRefreshStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Error           string       `json:"error,omitempty"`
	LastUpdatedTime *metav1.Time `json:"lastUpdatedTime,omitempty"`
	Phase           string       `json:"phase,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:printcolumn:name="PHASE",type=string,JSONPath=`.status.phase`
//+kubebuilder:printcolumn:name="LAST-UPDATED-TIME",type=string,JSONPath=".status.lastUpdatedTime"

// ECRTokenRefresh is the Schema for the ecrtokenrefreshes API
type ECRTokenRefresh struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ECRTokenRefreshSpec   `json:"spec,omitempty"`
	Status ECRTokenRefreshStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ECRTokenRefreshList contains a list of ECRTokenRefresh
type ECRTokenRefreshList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ECRTokenRefresh `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ECRTokenRefresh{}, &ECRTokenRefreshList{})
}
