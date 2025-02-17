/*
Copyright 2021.

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

// HardwareEventSpec defines the desired state of HardwareEvent
type HardwareEventSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	// TransportHost is amq host url  e.g.amqp://amq-router-service-name.amq-namespace.svc.cluster.local"
	TransportHost string `json:"transportHost"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:default=debug
	LogLevel string `json:"logLevel,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:default=10
	MsgParserTimeout int `json:"msgParserTimeout,omitempty"`

	// +kubebuilder:validation:Required
	NodeSelector map[string]string `json:"nodeSelector"`
}

// HardwareEventStatus defines the observed state of HardwareEvent
type HardwareEventStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	LastSynced *metav1.Time `json:"lastSyncTimestamp,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// HardwareEvent is the Schema for the hardwareevents API
type HardwareEvent struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HardwareEventSpec   `json:"spec,omitempty"`
	Status HardwareEventStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// HardwareEventList contains a list of HardwareEvent
type HardwareEventList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HardwareEvent `json:"items"`
}

func init() {
	SchemeBuilder.Register(&HardwareEvent{}, &HardwareEventList{})
}
