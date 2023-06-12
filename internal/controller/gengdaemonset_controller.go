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

package controller

import (
	"context"
	"errors"
	"fmt"
	v1 "k8s.io/api/core/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	appsv1beta1 "github.com/gengqianyu/operator/api/v1beta1"
)

// GengDaemonsetReconciler reconciles a GengDaemonset object
type GengDaemonsetReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=apps.gengqianyu.io,resources=gengdaemonsets,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=apps.gengqianyu.io,resources=gengdaemonsets/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=apps.gengqianyu.io,resources=gengdaemonsets/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the GengDaemonset object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.4/pkg/reconcile
func (r *GengDaemonsetReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)
	// TODO(user): your logic here
	gengDsp := &appsv1beta1.GengDaemonset{}
	// 从 request 中去接收一个 GengDaemonset 对象
	if err := r.Client.Get(ctx, req.NamespacedName, gengDsp); err != nil {
		panic(err)
	}
	// 如果这个 image 存在，我就在每一个 node 节点上，创建一个 pod
	if gengDsp.Spec.Image == "" {
		panic(errors.New("GengDaemonset.spec.iamge 不能为空"))
	}
	nlp := &v1.NodeList{}
	//获取所有 pod
	if err := r.Client.List(ctx, nlp); err != nil {
		panic(err)
	}
	// 遍历所有 node ，在 node 上启动一个 pod
	for _, n := range nlp.Items {
		//初始化 pod 对象
		pp := &v1.Pod{
			TypeMeta: v12.TypeMeta{
				APIVersion: "v1",
				Kind:       "pod",
			},
			ObjectMeta: v12.ObjectMeta{
				GenerateName: fmt.Sprintf("%s-", n.Name),
				Namespace:    gengDsp.Namespace,
			},
			Spec: v1.PodSpec{
				Containers: []v1.Container{
					{
						Image: gengDsp.Spec.Image,
						Name:  "gds-test-container",
					},
				},
				NodeName: n.Name,
			},
		}
		// 创建 pod 对象
		if err := r.Client.Create(ctx, pp); err != nil {
			panic(err)
		}
	}
	//测试 git rebase on origin myfeature 1
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *GengDaemonsetReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&appsv1beta1.GengDaemonset{}).
		Complete(r)
}
