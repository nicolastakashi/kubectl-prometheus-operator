package k8sclient

import (
	"context"
	"os"
	"path/filepath"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	monitoringv1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	monitoringv1alpha1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1alpha1"
	opClientv1 "github.com/prometheus-operator/prometheus-operator/pkg/client/versioned/typed/monitoring/v1"
	opClientv1alpha1 "github.com/prometheus-operator/prometheus-operator/pkg/client/versioned/typed/monitoring/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func homeDir() string {
	if h := homedir.HomeDir(); h != "" {
		return h
	}
	return os.Getenv("HOME")
}

type Clientv1 struct {
	*opClientv1.MonitoringV1Client
}

type Clientv1Alpha1 struct {
	*opClientv1alpha1.MonitoringV1alpha1Client
}

func NewClientv1(logger log.Logger) (*Clientv1, error) {
	kubeconfig := filepath.Join(homeDir(), ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		level.Error(logger).Log("msg", "Error building kubeconfig", "err", err)
		return nil, err
	}

	opClientv1, err := opClientv1.NewForConfig(config)
	if err != nil {
		level.Error(logger).Log("msg", "Error creating Prometheus Operator clientset", "err", err)
		return nil, err
	}

	return &Clientv1{opClientv1}, nil
}

func (c *Clientv1) CreatePrometheus(
	ctx context.Context,
	prometheus *monitoringv1.Prometheus,
	dryRun bool) error {
	opts := metav1.CreateOptions{}
	if dryRun {
		opts.DryRun = []string{"All"}
	}

	_, err := c.Prometheuses(prometheus.Namespace).Create(ctx, prometheus, opts)
	if err != nil {
		return err
	}

	return nil
}

func NewClientv1alpha1(logger log.Logger) (*Clientv1Alpha1, error) {
	kubeconfig := filepath.Join(homeDir(), ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		level.Error(logger).Log("msg", "Error building kubeconfig", "err", err)
		return nil, err
	}

	opClientv1alpha1, err := opClientv1alpha1.NewForConfig(config)
	if err != nil {
		level.Error(logger).Log("msg", "Error creating Prometheus Operator clientset", "err", err)
		return nil, err
	}

	return &Clientv1Alpha1{opClientv1alpha1}, nil
}

func (c *Clientv1Alpha1) CreatePrometheusAgent(
	ctx context.Context,
	prometheus *monitoringv1alpha1.PrometheusAgent,
	dryRun bool) error {
	opts := metav1.CreateOptions{}
	if dryRun {
		opts.DryRun = []string{"All"}
	}

	_, err := c.PrometheusAgents(prometheus.Namespace).Create(ctx, prometheus, opts)
	if err != nil {
		return err
	}

	return nil
}
