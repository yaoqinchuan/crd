/*
Copyright 2023 yqc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ServiceIngressControllerSpec defines the desired state of ServiceIngressController
type ServiceIngressControllerSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of ServiceIngressController. Edit serviceingresscontroller_types.go to remove/update
	Foo string `json:"foo,omitempty"`
}

// ServiceIngressControllerStatus defines the observed state of ServiceIngressController
type ServiceIngressControllerStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ServiceIngressController is the Schema for the serviceingresscontrollers API
type ServiceIngressController struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ServiceIngressControllerSpec   `json:"spec,omitempty"`
	Status ServiceIngressControllerStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ServiceIngressControllerList contains a list of ServiceIngressController
type ServiceIngressControllerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceIngressController `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ServiceIngressController{}, &ServiceIngressControllerList{})
}
