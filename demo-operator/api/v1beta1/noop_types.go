/*
Copyright 2022.

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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NoopSpec defines the desired state of Noop
type NoopSpec struct {
	// Foo is an example field of Noop.
	Foo string `json:"foo"`
	// Bar is an example field of Noop.
	Bar string `json:"bar,omitempty"`
}

// NoopStatus defines the observed state of Noop
type NoopStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Noop is the Schema for the noops API
type Noop struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NoopSpec   `json:"spec,omitempty"`
	Status NoopStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// NoopList contains a list of Noop
type NoopList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Noop `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Noop{}, &NoopList{})
}
