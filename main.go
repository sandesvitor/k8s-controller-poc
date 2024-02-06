package main

import (
	"flag"
	"fmt"
	"time"

	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"
	"k8s.io/sample-controller/pkg/signals"
)

var (
	masterURL  string
	kubeconfig string
)

func main() {
	klog.InitFlags(nil)
	flag.Parse()

	ctx := signals.SetupSignalHandler()
	logger := klog.FromContext(ctx)

	cfg, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
	if err != nil {
		logger.Error(err, "Error building kubeconfig")
		klog.FlushAndExit(klog.ExitFlushTimeout, 1)
	}

	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		logger.Error(err, "Error building kubernetes clientset")
		klog.FlushAndExit(klog.ExitFlushTimeout, 1)
	}

	test := kubeClient.AppsV1()
	deployList, err := test.Deployments("default").List(ctx, meta.ListOptions{})
	if err != nil {
		logger.Error(err, "Error getting Depoy list")
		klog.FlushAndExit(klog.ExitFlushTimeout, 1)
	}

	for _, item := range deployList.Items {
		logger.Info(fmt.Sprintf("DEPLOY [%s] have [%d] REPLICAS", item.GetName(), item.Status.Replicas))
	}

	// SIMULANDO UM WAIT PARA USAR DEPLOYMENT
	for {
		// PARA NÃO QUEBRAR O POD (NÃO QUERO USAR JOB)
		time.Sleep(10 * time.Second)
	}
}

func init() {
	flag.StringVar(&kubeconfig, "kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
	flag.StringVar(&masterURL, "master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
}
