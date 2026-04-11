# Container Insights with OpenTelemetry metrics for Amazon EKS

###### Preview

Container Insights with OpenTelemetry metrics provides visibility into the operational
health of your Amazon EKS cluster infrastructure. It is available in public preview at no
additional charge in US East (N. Virginia), US West (Oregon), Europe (Ireland), Asia Pacific
(Singapore), and Asia Pacific (Sydney).

The Amazon CloudWatch Observability EKS add-on collects open source metrics from your Amazon EKS
clusters and sends them to CloudWatch using the OpenTelemetry Protocol (OTLP) at 30 second
granularity. These metrics use metric names from their original sources, including cAdvisor,
Prometheus Node Exporter, NVIDIA DCGM, Kube State Metrics, and AWS Neuron Monitor. You can
query these metrics using PromQL in CloudWatch Query Studio or through the Prometheus compatible
query API.

Each metric is automatically enriched with up to 150 labels, including OpenTelemetry
semantic convention attributes and Kubernetes pod and node labels. PromQL handles aggregation
at query time, so each metric is published once per resource rather than at multiple
aggregation levels. The add-on also correlates accelerator metrics from AWS Neuron and AWS
Elastic Fabric Adapter with the specific pods and containers using them, providing visibility
that is not available from the metric sources alone.

To enable OTel Container Insights on an Amazon EKS cluster, install the Amazon CloudWatch Observability
EKS add-on version `v6.0.1-eksbuild.1` or later through the Amazon EKS console or
through infrastructure as code.

For more information about setting up OTel Container Insights, see [Setting up Container Insights](deploy-container-insights.md).

For more information about querying these metrics with PromQL, see [PromQL querying](cloudwatch-promql-querying.md).

## How OTel Container Insights compares to the Container Insights (enhanced)

The following table summarizes the differences between Container Insights (enhanced) and
OTel Container Insights.

FeatureContainer Insights (enhanced)OTel Container InsightsMetric namesCloudWatch-format metrics (for example, `pod_cpu_utilization`)Open-source native (for example,
`container_cpu_usage_seconds_total`)Labels per metric3–6 predefined dimensions per metricUp to 150 labels, including all Kubernetes pod and node labelsAggregationPre-aggregated at multiple levels (cluster, namespace, workload, pod)Raw per-resource metrics; aggregate at query time with PromQLQuery languageCloudWatch Metrics APIPromQL (Prometheus-compatible)Metric ingestionCloudWatch Logs in EMF formatOTLP endpoint

## How metrics are labeled

Each metric collected by OTel Container Insights carries labels from three
sources.

Telemetry source native labels

Labels from the original metric source (for example, cAdvisor provides labels such
as `pod`, `namespace`, and `container`). These are
preserved as datapoint attributes.

OpenTelemetry resource attributes

The add-on appends resource attributes following OpenTelemetry semantic conventions
for [Kubernetes](https://opentelemetry.io/docs/specs/semconv/resource/k8s), [Host](https://opentelemetry.io/docs/specs/semconv/resource/host), and
[Cloud](https://opentelemetry.io/docs/specs/semconv/resource/cloud), such as `k8s.pod.name`,
`k8s.namespace.name`, `k8s.node.name`,
`host.name`, and `cloud.region`. These attributes are
consistent across all metric sources.

Kubernetes pod and node labels

All pod labels and node labels discovered from the Kubernetes API are appended as
resource attributes with the prefixes `k8s.pod.label` and
`k8s.node.label`.

For more information about how to query these attributes using PromQL, see [PromQL querying](cloudwatch-promql-querying.md).

## Supported metrics

The following table lists the metric sources and categories collected by OTel Container
Insights.

Metric sourceMetric categoryPrerequisitescAdvisorCPU metrics-cAdvisorMemory metrics-cAdvisorNetwork metrics-cAdvisorDisk and filesystem metrics-Prometheus Node ExporterCPU metrics-Prometheus Node ExporterMemory metrics-Prometheus Node ExporterDisk metrics-Prometheus Node ExporterFilesystem metrics-Prometheus Node ExporterNetwork metrics-Prometheus Node ExporterSystem metrics-Prometheus Node ExporterVMStat metrics-Prometheus Node ExporterNetstat and socket metrics-NVIDIA DCGMGPU utilization and performance metrics[NVIDIA device\
plugin](https://github.com/NVIDIA/k8s-device-plugin) and [NVIDIA container\
toolkit](https://github.com/NVIDIA/nvidia-container-toolkit) must be installed.NVIDIA DCGMGPU memory metrics[NVIDIA device\
plugin](https://github.com/NVIDIA/k8s-device-plugin) and [NVIDIA container\
toolkit](https://github.com/NVIDIA/nvidia-container-toolkit) must be installed.NVIDIA DCGMGPU power and thermal metrics[NVIDIA device\
plugin](https://github.com/NVIDIA/k8s-device-plugin) and [NVIDIA container\
toolkit](https://github.com/NVIDIA/nvidia-container-toolkit) must be installed.NVIDIA DCGMGPU throttling metrics[NVIDIA device\
plugin](https://github.com/NVIDIA/k8s-device-plugin) and [NVIDIA container\
toolkit](https://github.com/NVIDIA/nvidia-container-toolkit) must be installed.NVIDIA DCGMGPU error and reliability metrics[NVIDIA device\
plugin](https://github.com/NVIDIA/k8s-device-plugin) and [NVIDIA container\
toolkit](https://github.com/NVIDIA/nvidia-container-toolkit) must be installed.NVIDIA DCGMGPU NVLink metrics[NVIDIA device\
plugin](https://github.com/NVIDIA/k8s-device-plugin) and [NVIDIA container\
toolkit](https://github.com/NVIDIA/nvidia-container-toolkit) must be installed.NVIDIA DCGMGPU informational metrics[NVIDIA device\
plugin](https://github.com/NVIDIA/k8s-device-plugin) and [NVIDIA container\
toolkit](https://github.com/NVIDIA/nvidia-container-toolkit) must be installed.AWS Neuron MonitorNeuronCore metrics[Neuron driver](https://awsdocs-neuron.readthedocs-hosted.com/en/latest/general/setup/neuron-setup/pytorch/neuronx/ubuntu/torch-neuronx-ubuntu22.html) and [Neuron device plugin](https://awsdocs-neuron.readthedocs-hosted.com/en/latest/containers/kubernetes-getting-started.html) must be installed.AWS Neuron MonitorNeuronDevice metrics[Neuron driver](https://awsdocs-neuron.readthedocs-hosted.com/en/latest/general/setup/neuron-setup/pytorch/neuronx/ubuntu/torch-neuronx-ubuntu22.html) and [Neuron device plugin](https://awsdocs-neuron.readthedocs-hosted.com/en/latest/containers/kubernetes-getting-started.html) must be installed.AWS Neuron MonitorNeuron system metrics[Neuron driver](https://awsdocs-neuron.readthedocs-hosted.com/en/latest/general/setup/neuron-setup/pytorch/neuronx/ubuntu/torch-neuronx-ubuntu22.html) and [Neuron device plugin](https://awsdocs-neuron.readthedocs-hosted.com/en/latest/containers/kubernetes-getting-started.html) must be installed.AWS Elastic Fabric AdapterEFA metrics[EFA device plugin](https://github.com/aws/eks-charts/tree/master/stable/aws-efa-k8s-device-plugin) must be installed.NVMeNVMe SMART metrics-Kube State MetricsPod, node, Deployment, DaemonSet, StatefulSet, ReplicaSet, Job, CronJob,
Service, Namespace, PersistentVolume, PersistentVolumeClaim metrics-Kubernetes API serverAPI server and etcd metrics-

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Container Insights with enhanced observability for Amazon EKS

Setting up Container Insights

All content copied from https://docs.aws.amazon.com/.
