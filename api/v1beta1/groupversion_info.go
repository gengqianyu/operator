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

// 程序包 v1beta1 包含应用程序 v1beta1 API 组的 API 架构定义
// +kubebuilder:object:generate=true
// +groupName=apps.gengqianyu.io
package v1beta1

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

var (
	// GroupVersion 是用于注册这些对象的组版本
	GroupVersion = schema.GroupVersion{
		Group:   "apps.gengqianyu.io",
		Version: "v1beta1",
	}

	// SchemeBuilder 用于将 go 类型 添加到 GroupVersionKind scheme 中
	SchemeBuilder = &scheme.Builder{
		GroupVersion: GroupVersion,
	}

	// AddToScheme 将此 group-version 中的类型添加到给定的 scheme 中。
	AddToScheme = SchemeBuilder.AddToScheme

	// “Scheme” 指的是一种提供了 Kinds 与对应 Go 类型映射的机制。
	// 它能够实现给定 Go 类型就知道其 GVK（Group、Version、Kind），给定 GVK 就知道其对应的 Go 类型。
)
