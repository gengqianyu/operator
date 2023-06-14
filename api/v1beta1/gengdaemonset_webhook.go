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
	"errors"

	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var gengdaemonsetlog = logf.Log.WithName("gengdaemonset-resource")

// SetupWebhookWithManager 定义了为 GengDaemonset 对象 SetupWebhook，这个方法会在 kubernetes 集群上启动一个 webhook service。
// apiService 在做 mutating 变形和 validating 校验的时候，会把对象发到这个 webhook service。
// kubernetes 为了安全起见，apiService 访问 webhook service 的时候是一定要使用 https 协议的。
// 因为必须为 webhook service 配置好 https 的证书。
func (r *GengDaemonset) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

// +kubebuilder:webhook:path=/mutate-apps-gengqianyu-io-v1beta1-gengdaemonset,mutating=true,failurePolicy=fail,sideEffects=None,groups=apps.gengqianyu.io,resources=gengdaemonsets,verbs=create;update,versions=v1beta1,name=mgengdaemonset.kb.io,admissionReviewVersions=v1
// 注意下面定义了一个匿名变量，没有任何引用，直接使用 下划线"_" 做为变量名，在编译时会丢弃掉了引用值，但是会对 _ 匿名变量做类型检查，防止忘记实现接口的全部方法。
// 这样做的作用是，提供了一个静态(编译时)检查，GengDaemonset 结构体 必须实现 admission.Defaulter 接口。
// 用 "_" 作为变量名，是要告诉编译器有效地丢弃RHS值，但要对其进行类型检查并评估它是否有任何副作用，但匿名变量本身不占用任何进程空间。
// 在开发时它是一个方便的构造，并且接口的方法集和 / 或由类型实现的方法经 常被改变。
// 该构造用作 防止忘记 匹配类型的方法集 和用于使它们 兼容的接口的方法集。
// 它有效地防止 go install 出现这种遗漏的破碎(中间)版本。
var _ webhook.Defaulter = &GengDaemonset{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
// Default 方法为 GengDaemonset 对象设置默认值，其实就是 mutating 的动作，
func (r *GengDaemonset) Default() {
	gengdaemonsetlog.Info("default", "name", r.Name)

	// TODO(user): fill in your defaulting logic.
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-apps-gengqianyu-io-v1beta1-gengdaemonset,mutating=false,failurePolicy=fail,sideEffects=None,groups=apps.gengqianyu.io,resources=gengdaemonsets,verbs=create;update,versions=v1beta1,name=vgengdaemonset.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &GengDaemonset{} // Validator 就是做校验的 webhook，其实就是对 crd 做 validating 校验的地方

// ValidateCreate 实现了 webhook.Validator 接口，因此将为 GengDaemonset 类型注册一个 webhook
func (r *GengDaemonset) ValidateCreate() (warnings admission.Warnings, err error) {
	gengdaemonsetlog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	// 做一个 image 属性校验，如果为空直接返回错误信息
	if r.Spec.Image == "" {
		return admission.Warnings{}, errors.New("image is required")
	}
	return admission.Warnings{}, nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *GengDaemonset) ValidateUpdate(old runtime.Object) (warnings admission.Warnings, err error) {
	gengdaemonsetlog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return admission.Warnings{}, nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *GengDaemonset) ValidateDelete() (warnings admission.Warnings, err error) {
	gengdaemonsetlog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return admission.Warnings{}, nil
}
