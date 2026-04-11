# Spans

Spans sent to X-Ray are ingested and managed in a log group called `aws/spans`.
This topic describes which CloudWatch Logs features are available for transaction spans.

###### Available features

The following CloudWatch Logs features are available for transaction spans.

- [Metric filters](../logs/monitoringlogdata.md) – Use metric filters to extract custom metrics from spans.

- [Subscriptions](../logs/subscriptions.md) – Use subscriptions to access a real-time feed of span events from CloudWatch Logs.

- [Log anomaly detection](../logs/logsanomalydetection.md) – Use log anomaly detection to establish a baseline for spans sent to the `aws/spans` log group.

- [Contributor Insights](contributorinsights.md) – Use Contributor Insights to analyze span data and create a time series displaying contributor data.

###### Unsupported features

The following are features not supported for transaction spans.

- Spans cannot be sent to CloudWatch Logs with the `PutLogEvents` API.

- Span data cannot be [enriched or transformed](../logs/cloudwatch-logs-transformation.md).

###### Note

Span ingestion is charged separately from log ingestion.
For information about pricing, see [Amazon CloudWatch Pricing](https://aws.amazon.com/cloudwatch/pricing).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Using Transaction Search with CloudFormation

Searching and analyzing spans

All content copied from https://docs.aws.amazon.com/.
