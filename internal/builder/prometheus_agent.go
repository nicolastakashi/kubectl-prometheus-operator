package builder

import (
	monitoringv1alpha1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PrometheusAgentBuilder struct {
	monitoringv1alpha1.PrometheusAgent
}

func NewPrometheusAgent() *PrometheusAgentBuilder {
	return &PrometheusAgentBuilder{
		PrometheusAgent: monitoringv1alpha1.PrometheusAgent{
			TypeMeta: metav1.TypeMeta{
				Kind:       monitoringv1alpha1.PrometheusAgentKindKey,
				APIVersion: monitoringv1alpha1.SchemeGroupVersion.String(),
			},
			ObjectMeta: metav1.ObjectMeta{},
		},
	}
}

func (psb *PrometheusAgentBuilder) WithName(name string) *PrometheusAgentBuilder {
	psb.PrometheusAgent.ObjectMeta.Name = name
	return psb
}

func (psb *PrometheusAgentBuilder) WithNamespace(namespace string) *PrometheusAgentBuilder {
	psb.PrometheusAgent.ObjectMeta.Namespace = namespace
	return psb
}

func (psb *PrometheusAgentBuilder) WithReplicas(replicas int32) *PrometheusAgentBuilder {
	psb.PrometheusAgent.Spec.Replicas = &replicas
	return psb
}

func (psb *PrometheusAgentBuilder) WithShards(shards int32) *PrometheusAgentBuilder {
	psb.PrometheusAgent.Spec.Shards = &shards
	return psb
}

func (psb *PrometheusAgentBuilder) Build() *monitoringv1alpha1.PrometheusAgent {
	return &psb.PrometheusAgent
}
