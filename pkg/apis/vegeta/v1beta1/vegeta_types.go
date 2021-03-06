/*
Copyright 2018 Dario Maiocchi.

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

// json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// VegetaSpec defines the desired state of Vegeta
type VegetaSpec struct {
	// Attack contains the options for vegeta CLI/lib attack
	Attack string `json:"attack,omitempty"`
	// numAttack, with this you can run e.g. 3 times attack
	NumAttack int32 `json:"numAttack,omitempty"`
}

// VegetaStatus defines the observed state of Vegeta
type VegetaStatus struct {
	// CountAttack will count how many attack the same type are running
	CountAttack int32 `json:"countAttack,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Vegeta is the Schema for the vegeta API
// +k8s:openapi-gen=true
type Vegeta struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VegetaSpec   `json:"spec,omitempty"`
	Status VegetaStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VegetaList contains a list of Vegeta
type VegetaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Vegeta `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Vegeta{}, &VegetaList{})
}
