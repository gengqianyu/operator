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

// 这些注释标签，kubebuilder 在生成 CRD 的时候，根据这些标签可以生成一些 CRD 的属性，
// 比如定义一个+kubebuilder:scope=Namespace/Cluster 标签就是告诉 kubebuilder 你帮我生成的这个 GengDaemonset crd 对象，它是 namespace 级别的还是 cluster 集群级别的。
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
	// 其实下面这段代码是，向 SchemeBuilder 中注册了一个如下函数，
	// 然后在 main 执行 appsv1beta1.AddToScheme(scheme) 的时候，执行如下函数，从而将 GengDaemonset 类型的序列化规则 添加到 scheme 中。
	// func(scheme *runtime.Scheme) error {
	//		scheme.AddKnownTypes(bld.GroupVersion, object...)
	//		metav1.AddToGroupVersion(scheme, bld.GroupVersion)
	//		return nil
	//	}
	SchemeBuilder.Register(
		&GengDaemonset{},
		&GengDaemonsetList{},
	)
}
