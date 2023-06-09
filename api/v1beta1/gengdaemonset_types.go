/*
Copyright 2023.

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

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// GengDaemonsetSpec defines the desired state of GengDaemonset
type GengDaemonsetSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// 修改 Foo 属性为 Image 注意 json tag 也要修改，我们要实现一个简单的 daemonset 的 operator 。
	Image string `json:"image,omitempty"`
}

// GengDaemonsetStatus defines the observed state of GengDaemonset
type GengDaemonsetStatus struct {
	// 我们在 status 定义一个可用副本数的属性
	AvailableReplicas int `json:"availableReplicas,omitempty"`
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// GengDaemonset is the Schema for the gengdaemonsets API
type GengDaemonset struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GengDaemonsetSpec   `json:"spec,omitempty"`
	Status GengDaemonsetStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// GengDaemonsetList contains a list of GengDaemonset
type GengDaemonsetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GengDaemonset `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GengDaemonset{}, &GengDaemonsetList{})
}
