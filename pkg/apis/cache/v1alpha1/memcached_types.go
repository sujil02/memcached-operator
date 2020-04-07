package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// MemcachedSpec defines the desired state of Memcached
// +k8s:openapi-gen=true
type MemcachedSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html

	// The desired number of member Pods for the deployment
	// +operator-sdk:gen-csv:customresourcedefinitions.specDescriptors=true
	// +operator-sdk:gen-csv:customresourcedefinitions.specDescriptors.displayName="Pod Count"
	// +operator-sdk:gen-csv:customresourcedefinitions.specDescriptors.x-descriptors="urn:alm:descriptor:com.tectonic.ui:podCount,urn:alm:descriptor:io.kubernetes:custom"
	Size int32 `json:"size"`
	// Memcached version
	// +operator-sdk:gen-csv:customresourcedefinitions.specDescriptors=true
	Version string `json:"version,omitempty"`
}

// MemcachedStatus defines the observed state of Memcached
// +k8s:openapi-gen=true
type MemcachedStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html

	// Nodes are the names of the memcached pods
	// +listType=set
	Nodes []string `json:"nodes"`
	// Phase is the memcached phase
	Phase MemcachedPhase `json:"phase,omitempty"`
}

type MemcachedPhase string

const (
	MemcachedPhaseCreating MemcachedPhase = "Creating"
	MemcachedPhaseRunning  MemcachedPhase = "Running"
	MemcachedPhaseFailed   MemcachedPhase = "Failed"
	MemcachedPhaseDeleting MemcachedPhase = "Deleting"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Represents a cluster of Memcached apps
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=memcacheds,scope=Namespaced
// +operator-sdk:gen-csv:customresourcedefinitions.displayName="Memcached App"
// +operator-sdk:gen-csv:customresourcedefinitions.resources=`Deployment,v1,""`
// +operator-sdk:gen-csv:customresourcedefinitions.resources=`Service,v1,""`
type Memcached struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MemcachedSpec   `json:"spec,omitempty"`
	Status MemcachedStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MemcachedList contains a list of Memcached
type MemcachedList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Memcached `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Memcached{}, &MemcachedList{})
}
