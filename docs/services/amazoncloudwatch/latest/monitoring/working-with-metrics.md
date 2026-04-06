# Metrics in Amazon CloudWatch

Metrics are data about the performance of your systems. Amazon CloudWatch collects metrics through
two paths: AWS vended metrics from services such as Amazon EC2, Amazon EBS, and Amazon RDS, and custom
metrics that you publish using the OpenTelemetry Protocol (OTLP) or the CloudWatch API. Many AWS
services provide free metrics for resources by default. You can also enable detailed monitoring
for some resources, such as Amazon EC2 instances.

CloudWatch supports two query languages for metrics. The Prometheus Query Language (PromQL)
provides flexible analytics for OpenTelemetry metrics and AWS vended metrics published in
OpenTelemetry format. CloudWatch Metrics Insights provides a SQL-based query engine for traditional CloudWatch metrics. Both
support graphing, dashboards, and alarms.

Key Topics:

- [Metrics concepts](cloudwatch-concepts.md)

- [Basic monitoring and detailed monitoring in CloudWatch](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/cloudwatch-metrics-basic-detailed.html)

- [Publish custom metrics](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/publishingMetrics.html)

- [Send metrics using OpenTelemetry (new)](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-OpenTelemetry-Sections.html)

- [Publish AWS vended metrics as OpenTelemetry metrics (new)](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/publishingMetrics.html)

- [Query metrics with PromQL (new)](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-PromQL.html)

- [Query your CloudWatch metrics with CloudWatch Metrics Insights](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/query_with_cloudwatch-metrics-insights.html)

- [Use metrics explorer to monitor resources by their tags and properties](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Metrics-Explorer.html)

- [Use metric streams](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Metric-Streams.html)

- [View available metrics](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/viewing_metrics_with_cloudwatch.html)

- [Graphing metrics](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/graph_metrics.html)

- [Using CloudWatch anomaly detection](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Anomaly_Detection.html)

- [Using math expressions with CloudWatch metrics](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/using-metric-math.html)

- [Use search expressions in graphs](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/using-search-expressions.html)

- [Get statistics for a metric](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/getting-metric-statistics.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Changing the time range or time zone format

Concepts
