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

package main

import (
	"flag"
	"os"

	// Import all Kubernetes client auth plugins (e.g. Azure, GCP, OIDC, etc.)
	// to ensure that exec-entrypoint and run can make use of them.
	_ "k8s.io/client-go/plugin/pkg/client/auth"

	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/healthz"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	appsv1beta1 "github.com/gengqianyu/operator/api/v1beta1"
	"github.com/gengqianyu/operator/internal/controller"
	//+kubebuilder:scaffold:imports
)

var (
	// Scheme 指的是一种提供了 Kinds 与对应 Go 类型映射的机制。
	// 它能够实现给定 Go 类型就知道其 GVK（Group、Version、Kind），
	// 给定 GVK 就知道其对应的 Go 类型
	scheme   = runtime.NewScheme()
	setupLog = ctrl.Log.WithName("setup")
)

func init() {
	// 在 scheme 中注册 GengDaemonset 类型的序列化和反序列化规则。
	// 这样在创建或读取 GengDaemonset 类型的对象时，GORM 就知道如何处理它们了。
	utilruntime.Must(
		// clientgoscheme 是 Kubernetes 提供的一个包，它包含了所有内置资源类型的序列化和反序列化逻辑。
		// 通过调用 AddToScheme 方法，将这些内置资源类型的序列化规则 添加到 schem 中，
		// 这样在创建或读取这些资源时，GORM 就知道如何处理它们了。
		clientgoscheme.AddToScheme(scheme), // 将 k8s 内置资源类型的序列化规则 添加到 scheme 中
	)

	utilruntime.Must(
		// 结合 api/v1beta1/gengdaemonset_types.go 中的 init 函数就明白了，注册的流程。

		appsv1beta1.AddToScheme(scheme), // 将 GengDaemonset 类型的序列化规则 添加到 scheme 中
	)
	//+kubebuilder:scaffold:scheme
	// 这样完成以上两步后，scheme 中就包含了 GengDaemonset 类型 和 k8s 内置类型的序列化和反序列化规则了。
}

// main 函数是程序的入口点，它初始化并启动一个 Kubernetes 控制器管理器
// 该控制器管理器用于协调 GengDaemonset 自定义资源的生命周期
// 它设置了日志记录器，解析了命令行参数，并配置了控制器管理器的各种选项
// 包括 metrics 绑定地址、健康检查绑定地址、是否启用 leader 选举等
func main() {
	// 定义 metrics 绑定地址的命令行参数，默认值为 ":8080"
	var metricsAddr string
	// 定义是否启用 leader 选举的命令行参数，默认值为 false
	var enableLeaderElection bool
	// 定义健康检查绑定地址的命令行参数，默认值为 ":8081"
	var probeAddr string
	// 使用 flag 包解析命令行参数
	flag.StringVar(&metricsAddr, "metrics-bind-address", ":8080", "指标端点绑定到的地址.")
	flag.StringVar(&probeAddr, "health-probe-bind-address", ":8081", "探针端点绑定到的地址.")
	flag.BoolVar(&enableLeaderElection, "leader-elect", false, "为 Controller Manager 启用 leader 选举.启用此选项将确保只有一个活动的控制器管理器.")
	// 使用 zap 库的 Options 结构体设置日志记录器选项
	opts := zap.Options{
		Development: true,
	}
	// 将 flag 包的命令行参数与 zap 库的选项绑定
	opts.BindFlags(flag.CommandLine)
	// 解析命令行参数
	flag.Parse()

	// 使用 zap 库创建一个新的日志记录器，并将其设置为全局日志记录器
	ctrl.SetLogger(zap.New(zap.UseFlagOptions(&opts)))

	// 创建一个新的控制器管理器实例，并配置其选项
	// Manager 具有 3 个职责：
	// 负责运行所有的 Controllers；
	// 初始化共享 caches，包含 listAndWatch 功能；
	// 初始化 clients 用于与 Api Server 通信。
	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:                 scheme,
		MetricsBindAddress:     metricsAddr,
		Port:                   9443,
		HealthProbeBindAddress: probeAddr,
		LeaderElection:         enableLeaderElection,
		LeaderElectionID:       "a17b38b3.gengqianyu.io",
		// LeaderElectionReleaseOnCancel defines if the leader should step down voluntarily
		// when the Manager ends. This requires the binary to immediately end when the
		// Manager is stopped, otherwise, this setting is unsafe. Setting this significantly
		// speeds up voluntary leader transitions as the new leader don't have to wait
		// LeaseDuration time first.
		//
		// In the default scaffold provided, the program ends immediately after
		// the manager stops, so would be fine to enable this option. However,
		// if you are doing or is intended to do any operation such as perform cleanups
		// after the manager stops then its usage might be unsafe.
		// LeaderElectionReleaseOnCancel: true,
	})
	// 如果创建控制器管理器时发生错误，记录错误并退出程序
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

	// 创建 GengDaemonset 控制器的协调器，并将其设置到管理器中
	if err = (&controller.GengDaemonsetReconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		// 如果设置控制器时发生错误，记录错误并退出程序
		setupLog.Error(err, "unable to create controller", "controller", "GengDaemonset")
		os.Exit(1)
	}

	// 设置 GengDaemonset 自定义资源的 Webhook
	if err = (&appsv1beta1.GengDaemonset{}).SetupWebhookWithManager(mgr); err != nil {
		// 如果设置 Webhook 时发生错误，记录错误并退出程序
		setupLog.Error(err, "unable to create webhook", "webhook", "GengDaemonset")
		os.Exit(1)
	}
	//+kubebuilder:scaffold:builder

	// 为控制器管理器添加健康检查
	if err := mgr.AddHealthzCheck("healthz", healthz.Ping); err != nil {
		// 如果添加健康检查时发生错误，记录错误并退出程序
		setupLog.Error(err, "unable to set up health check")
		os.Exit(1)
	}
	// 为控制器管理器添加就绪检查
	if err := mgr.AddReadyzCheck("readyz", healthz.Ping); err != nil {
		// 如果添加就绪检查时发生错误，记录错误并退出程序
		setupLog.Error(err, "unable to set up ready check")
		os.Exit(1)
	}

	// 记录控制器管理器启动的信息
	setupLog.Info("starting manager")
	// 启动控制器管理器，并处理来自 Kubernetes API 的请求
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		// 如果在运行管理器时发生错误，记录错误并退出程序
		setupLog.Error(err, "problem running manager")
		os.Exit(1)
	}
}
