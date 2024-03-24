package builder

import (
	monitoringv1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PrometheusServerBuilder struct {
	monitoringv1.Prometheus
}

func NewPrometheusServer() *PrometheusServerBuilder {
	return &PrometheusServerBuilder{
		Prometheus: monitoringv1.Prometheus{
			TypeMeta: metav1.TypeMeta{
				Kind:       monitoringv1.PrometheusKindKey,
				APIVersion: monitoringv1.SchemeGroupVersion.String(),
			},
			ObjectMeta: metav1.ObjectMeta{},
		},
	}
}

func (psb *PrometheusServerBuilder) WithName(name string) *PrometheusServerBuilder {
	psb.Prometheus.ObjectMeta.Name = name
	return psb
}

func (psb *PrometheusServerBuilder) WithNamespace(namespace string) *PrometheusServerBuilder {
	psb.Prometheus.ObjectMeta.Namespace = namespace
	return psb
}

func (psb *PrometheusServerBuilder) WithReplicas(replicas int32) *PrometheusServerBuilder {
	psb.Prometheus.Spec.Replicas = &replicas
	return psb
}

func (psb *PrometheusServerBuilder) WithShards(shards int32) *PrometheusServerBuilder {
	psb.Prometheus.Spec.Shards = &shards
	return psb
}

func (psb *PrometheusServerBuilder) Build() *monitoringv1.Prometheus {
	return &psb.Prometheus
}
