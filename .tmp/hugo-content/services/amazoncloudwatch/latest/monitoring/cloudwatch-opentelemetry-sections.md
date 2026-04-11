# OpenTelemetry

OpenTelemetry is an open-source observability framework that provides vendor-agnostic
instrumentation for collecting metrics, logs, and traces from your applications. Amazon CloudWatch
supports OpenTelemetry natively across all three signal types: metrics queryable with PromQL
for flexible and scalable analytics, logs searchable with Logs Insights, and traces
explorable with [Transaction Search](cloudwatch-transaction-search.md). Because all three signals share a common data model with
consistent attributes, you can correlate across metrics, logs, and traces to move from
detection to root cause faster. OpenTelemetry metrics carry rich semantic labels and support
higher granularity than traditional CloudWatch metrics, enabling precise filtering, aggregation,
and analysis across your applications and infrastructure.

CloudWatch supports OpenTelemetry metrics, allowing you to send custom OTel metrics directly to
CloudWatch and query them using PromQL alongside AWS vended metrics from over 70 services. You
can use PromQL to build dashboards, set CloudWatch Alarms, and explore metrics in Query Studio,
the native PromQL console experience. OTel logs sent to CloudWatch are available in Logs Insights
for interactive querying and in LiveTail for real-time streaming.

For application performance monitoring, CloudWatch [Application Signals](cloudwatch-application-monitoring-intro.md) provides a unified, application-centric view of your
services and dependencies. Application Signals uses OTLP traces to help you monitor
application health, triage issues, and identify the impact on end users. With Transaction
Search, you can explore OTLP spans interactively, find transactions using attributes such as
customer identifiers or order numbers, correlate transactions to business events such as
failed payments, and trace interactions between application components to establish root
cause.

![OpenTelemetry overview](https://docs.aws.amazon.com/images/AmazonCloudWatch/latest/monitoring/images/OpenTelemetry.png)

###### Topics

- [OTLP Endpoints](cloudwatch-otlpendpoint.md)

- [Getting started](cloudwatch-otlpgettingstarted.md)

- [Troubleshooting](cloudwatch-otlptroubleshooting.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Troubleshooting

OTLP Endpoints

All content copied from https://docs.aws.amazon.com/.
