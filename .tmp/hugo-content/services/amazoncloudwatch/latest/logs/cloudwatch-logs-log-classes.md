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

[CloudWatch Logs Insights query commands](analyzinglogdata.md)

Yes ✓

Yes ✓ (Most commands– see [Logs Insights QL commands supported in log classes](cwl-analyzelogdata-classes.md).)

[CloudWatch Logs Insights discovered fields](cwl-analyzelogdata-discoverable-fields.md)

Yes ✓

Yes ✓

[Facets](cloudwatchlogs-facets.md)

Yes ✓

No

[Using CloudWatch Pipeline to transform\
logs](../monitoring/cloudwatch-pipelines.md)

Yes ✓

Yes ✓

[Export to Amazon S3](s3export.md)

Yes ✓

Yes ✓

[S3 Tables\
Integration](s3-tables-integration.md)

Yes ✓

Yes ✓

[Scheduled Queries](scheduledqueries.md)

Yes ✓

Yes ✓

[Using OpenSearch PPL or OpenSearch SQL to query \
in CloudWatch Logs Insights;](analyzinglogdata.md)

Yes ✓

Yes ✓

[Natural language query assist](cloudwatchlogs-insights-query-assist.md)

Yes ✓

No

[CloudWatch Logs Anomaly Detection](logsanomalydetection.md)

Yes ✓

No

[Live Tail](cloudwatchlogs-livetail.md)

Yes ✓

No

[Field indexing](cloudwatchlogs-field-indexing.md)

Yes ✓

No

[Compare to previous time range](cwl-analyzelogdata-compare.md)

Yes ✓

No

[Subscription filters](subscriptions.md)

Yes ✓

No

[GetLogEvents](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-getlogevents.md) and [FilterLogEvents](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-filterlogevents.md) API operations

Yes ✓

Not supported. Use CloudWatch Logs Insights to view log events stored in log groups in the Infrequent
Access log class.

[Metric filters](monitoringlogdata.md)

Yes ✓

No

[Container Insights log ingestion](../monitoring/containerinsights.md)

Yes ✓

No

[Lambda Insights log ingestion](../monitoring/lambda-insights.md)

Yes ✓

No

[Sensitive data protection with masking](mask-sensitive-log-data.md)

Yes ✓

Yes ✓

[Embedded metrics format](../monitoring/cloudwatch-embedded-metric-format.md)

Yes ✓

No

###### Note

In addition to these two log classes, there is a `Delivery` log class. Use the `Delivery` log class only for delivering AWS Lambda logs to store in Amazon S3 or
Amazon Data Firehose. Log events in log groups in the Delivery class are kept in CloudWatch Logs for two days. This retention period is fixed and cannot be changed. This log class doesn't offer rich CloudWatch Logs capabilities such as
CloudWatch Logs Insights queries.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Billing and costs

Getting started

All content copied from https://docs.aws.amazon.com/.
