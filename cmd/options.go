package cmd

type createPrometheusOptions struct {
	name      string
	namespace string
	mode      string
	replicas  int
	shards    int
}
