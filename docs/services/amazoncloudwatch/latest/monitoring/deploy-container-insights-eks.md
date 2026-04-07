# Setting up Container Insights on Amazon EKS and Kubernetes

Container Insights is supported on Amazon EKS versions 1.23 and later. The quick start method
of installation is supported only on versions 1.24 and later.

The overall process for setting up Container Insights on Amazon EKS or Kubernetes is as
follows:

1. Verify that you have the necessary prerequisites.

2. Set up the Amazon CloudWatch Observability EKS add-on, the CloudWatch agent, or AWS Distro for
    OpenTelemetry on your cluster to send metrics to CloudWatch.

###### Note

To use Container Insights with enhanced observability for Amazon EKS, you must use the
Amazon CloudWatch Observability EKS add-on or the CloudWatch agent. For more information about
this version of Container Insights, see [Container Insights with enhanced observability for Amazon EKS](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/container-insights-detailed-metrics.html).

To use Container Insights with Fargate, you must use AWS Distro for
OpenTelemetry. Container Insights with enhanced observability for Amazon EKS is not
supported on Fargate.

###### Note

Container Insights now supports Windows worker nodes in an Amazon EKS cluster.
Container Insights with enhanced observability for Amazon EKS is also supported on Windows.
For information about enabling Container Insights on Windows, see [Using the CloudWatch agent with Container Insights enhanced observability enabled](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Container-Insights-EKS-agent.html).

To use Container Insights with OpenTelemetry metrics, install the Amazon CloudWatch
    Observability EKS add-on version `v6.0.1-eksbuild.1` or later. For more
    information, see [Container Insights with OpenTelemetry metrics for Amazon EKS](container-insights-otel-metrics.md).

Set up Fluent Bit or Fluentd to send logs to CloudWatch Logs. (This is enabled by default if
    you install the Amazon CloudWatch Observability EKS add-on.)

You can perform these steps at once as part of the quick start setup if you are
    using the CloudWatch agent, or do them separately.

3. (Optional) Set up Amazon EKS control plane logging.

4. (Optional) Set up the CloudWatch agent as a StatsD endpoint on the cluster to send StatsD
    metrics to CloudWatch.

5. (Optional) Enable App Mesh Envoy Access Logs.

With the original version of Container Insights, metrics collected and logs ingested are
charged as custom metrics. With Container Insights with enhanced observability for Amazon EKS,
Container Insights metrics and logs are charged per observation instead of being charged per
metric stored or log ingested. For more information about CloudWatch pricing, see [Amazon CloudWatch Pricing](https://aws.amazon.com/cloudwatch/pricing).

###### Topics

- [Verifying prerequisites for Container Insights in CloudWatch](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Container-Insights-prerequisites.html)

- [Using the CloudWatch agent with Container Insights enhanced observability enabled](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Container-Insights-EKS-agent.html)

- [Using AWS Distro for OpenTelemetry](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Container-Insights-EKS-otel.html)

- [Send logs to CloudWatch Logs](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Container-Insights-EKS-logs.html)

- [Updating or deleting Container Insights on Amazon EKS and Kubernetes](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/ContainerInsights-update-delete.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Set up FireLens to send logs to CloudWatch Logs

Verifying prerequisites
