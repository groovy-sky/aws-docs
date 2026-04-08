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

- [Basic monitoring and detailed monitoring in CloudWatch](cloudwatch-metrics-basic-detailed.md)

- [Publish custom metrics](publishingmetrics.md)

- [Send metrics using OpenTelemetry (new)](cloudwatch-opentelemetry-sections.md)

- [Publish AWS vended metrics as OpenTelemetry metrics (new)](publishingmetrics.md)

- [Query metrics with PromQL (new)](cloudwatch-promql.md)

- [Query your CloudWatch metrics with CloudWatch Metrics Insights](query-with-cloudwatch-metrics-insights.md)

- [Use metrics explorer to monitor resources by their tags and properties](cloudwatch-metrics-explorer.md)

- [Use metric streams](cloudwatch-metric-streams.md)

- [View available metrics](viewing-metrics-with-cloudwatch.md)

- [Graphing metrics](graph-metrics.md)

- [Using CloudWatch anomaly detection](cloudwatch-anomaly-detection.md)

- [Using math expressions with CloudWatch metrics](using-metric-math.md)

- [Use search expressions in graphs](using-search-expressions.md)

- [Get statistics for a metric](getting-metric-statistics.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Changing the time range or time zone format

Concepts

All content copied from https://docs.aws.amazon.com/.
