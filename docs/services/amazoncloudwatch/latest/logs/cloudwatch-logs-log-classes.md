# Log classes

CloudWatch Logs offers two classes of log groups:

- The _CloudWatch Logs Standard_ log class is a full-featured option for logs that
require real-time monitoring or logs that you access frequently.

- The _CloudWatch Logs Infrequent Access_ log class is a new log class that
you can use to cost-effectively consolidate
your logs. This log class offers a subset of CloudWatch Logs capabilities including managed ingestion,
storage, cross-account log analytics, and encryption with a lower ingestion price per GB. The
Infrequent
Access log class is ideal for ad-hoc querying and after-the-fact forensic analysis on infrequently accessed logs.

###### Note

For charges, the Standard and Infrequent Access log classes differ in ingestion costs only. Storage charges
and CloudWatch Logs Insights charges are the same in each log class.

For more information about CloudWatch Logs pricing, see [Amazon CloudWatch Pricing](https://aws.amazon.com/cloudwatch/pricing).

###### Important

After a log group is created, its log class can't be changed.

## Supported features

The following table lists the features for each log class.

FeatureStandardInfrequent Access

Fully managed log ingestion and storage

Yes ✓

Yes ✓

[Cross-account features](../monitoring/cloudwatch-unified-cross-account.md)

Yes ✓

Yes ✓

[Encryption with AWS KMS](encrypt-log-data-kms.md)

Yes ✓

Yes ✓

[CloudWatch Logs Insights query commands](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AnalyzingLogData.html)

Yes ✓

Yes ✓ (Most commands– see [Logs Insights QL commands supported in log classes](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL_AnalyzeLogData_Classes.html).)

[CloudWatch Logs Insights discovered fields](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL_AnalyzeLogData-discoverable-fields.html)

Yes ✓

Yes ✓

[Facets](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatchLogs-Facets.html)

Yes ✓

No

[Using CloudWatch Pipeline to transform\
logs](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/cloudwatch-pipelines.html)

Yes ✓

Yes ✓

[Export to Amazon S3](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/S3Export.html)

Yes ✓

Yes ✓

[S3 Tables\
Integration](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/s3-tables-integration.html)

Yes ✓

Yes ✓

[Scheduled Queries](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/ScheduledQueries.html)

Yes ✓

Yes ✓

[Using OpenSearch PPL or OpenSearch SQL to query \
in CloudWatch Logs Insights;](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/AnalyzingLogData.html)

Yes ✓

Yes ✓

[Natural language query assist](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatchLogs-Insights-Query-Assist.html)

Yes ✓

No

[CloudWatch Logs Anomaly Detection](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/LogsAnomalyDetection.html)

Yes ✓

No

[Live Tail](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatchLogs_LiveTail.html)

Yes ✓

No

[Field indexing](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CloudWatchLogs-Field-Indexing.html)

Yes ✓

No

[Compare to previous time range](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL_AnalyzeLogData_Compare.html)

Yes ✓

No

[Subscription filters](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/Subscriptions.html)

Yes ✓

No

[GetLogEvents](https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_GetLogEvents.html) and [FilterLogEvents](https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_FilterLogEvents.html) API operations

Yes ✓

Not supported. Use CloudWatch Logs Insights to view log events stored in log groups in the Infrequent
Access log class.

[Metric filters](https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/MonitoringLogData.html)

Yes ✓

No

[Container Insights log ingestion](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/ContainerInsights.html)

Yes ✓

No

[Lambda Insights log ingestion](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Lambda-Insights.html)

Yes ✓

No

[Sensitive data protection with masking](mask-sensitive-log-data.md)

Yes ✓

Yes ✓

[Embedded metrics format](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Embedded_Metric_Format.html)

Yes ✓

No

###### Note

In addition to these two log classes, there is a `Delivery` log class. Use the `Delivery` log class only for delivering AWS Lambda logs to store in Amazon S3 or
Amazon Data Firehose. Log events in log groups in the Delivery class are kept in CloudWatch Logs for two days. This retention period is fixed and cannot be changed. This log class doesn't offer rich CloudWatch Logs capabilities such as
CloudWatch Logs Insights queries.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Billing and costs

Getting started
