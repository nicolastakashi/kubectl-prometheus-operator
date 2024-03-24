/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/go-kit/log/level"
	"github.com/prometheus-operator/kubectl-prometheus-operator/internal/builder"
	"github.com/prometheus-operator/kubectl-prometheus-operator/internal/k8sclient"
	"github.com/spf13/cobra"
	"github.com/thanos-io/thanos/pkg/logging"
	"sigs.k8s.io/yaml"
)

var options = createPrometheusOptions{}

// prometheusCmd represents the prometheus command
var prometheusCmd = &cobra.Command{
	Use:   "prometheus",
	Short: "Command to create prometheus and prometheus agent resource",
	Long:  `This command is used to create prometheus and prometheus agent resources`,
	Run: func(cmd *cobra.Command, args []string) {

		//TODO:(nicolas) Change the logger implementation
		logger := logging.NewLogger("info", "json", "")

		dryRun, err := cmd.Flags().GetBool("dry-run")
		if err != nil {
			level.Error(logger).Log("msg", "Error getting dry-run flag", "err", err)
			os.Exit(1)
		}

		output, err := cmd.Flags().GetString("output")
		if err != nil {
			level.Error(logger).Log("msg", "Error getting output flag", "err", err)
			os.Exit(1)
		}

		if options.mode == "server" {
			client, err := k8sclient.NewClientv1(logger)
			if err != nil {
				level.Error(logger).Log("msg", "Error creating Prometheus Operator clientset", "err", err)
				os.Exit(1)
			}

			prometheus := builder.NewPrometheusServer().
				WithName(options.name).
				WithNamespace(options.namespace).
				WithReplicas(int32(options.replicas)).
				WithShards(int32(options.shards)).
				Build()

			err = client.CreatePrometheus(cmd.Context(), prometheus, dryRun)
			if err != nil {
				level.Error(logger).Log("msg", "Error creating Prometheus instance", "err", err)
				os.Exit(1)
			}

			if output == "yaml" {
				yaml, err := yaml.Marshal(prometheus)
				if err != nil {
					level.Error(logger).Log("msg", "Error marshalling Prometheus object", "err", err)
					os.Exit(1)
				}
				fmt.Print(string(yaml))
			}
		}

		if options.mode == "agent" {
			client, err := k8sclient.NewClientv1alpha1(logger)
			if err != nil {
				level.Error(logger).Log("msg", "Error creating Prometheus Operator clientset", "err", err)
				os.Exit(1)
			}
			prometheus := builder.NewPrometheusAgent().
				WithName(options.name).
				WithNamespace(options.namespace).
				WithReplicas(int32(options.replicas)).
				WithShards(int32(options.shards)).
				Build()

			err = client.CreatePrometheusAgent(cmd.Context(), prometheus, dryRun)
			if err != nil {
				level.Error(logger).Log("msg", "Error creating Prometheus instance", "err", err)
				os.Exit(1)
			}

			if output == "yaml" {
				yaml, err := yaml.Marshal(prometheus)
				if err != nil {
					level.Error(logger).Log("msg", "Error marshalling Prometheus object", "err", err)
					os.Exit(1)
				}
				fmt.Print(string(yaml))
			}

		}
	},
}

func init() {
	createCmd.AddCommand(prometheusCmd)
	prometheusCmd.Flags().StringVar(&options.name, "name", "example", "The name of the Prometheus instance")
	prometheusCmd.Flags().StringVar(&options.namespace, "namespace", "default", "The namespace to create the Prometheus instance in")
	prometheusCmd.Flags().StringVar(&options.mode, "mode", "server", "The mode to create the Prometheus instance in (server or agent)")
	prometheusCmd.Flags().IntVar(&options.replicas, "replicas", 1, "The number of Prometheus instances to create")
	prometheusCmd.Flags().IntVar(&options.shards, "shards", 1, "The number of shards to create the Prometheus instance with")
}
