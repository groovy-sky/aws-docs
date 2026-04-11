# Monitoring with CloudWatch metrics

You can use the tables in this section to review the metrics that Amazon CloudWatch Logs sends to Amazon CloudWatch every minute.

## CloudWatch Logs metrics

The `AWS/Logs` namespace includes the following metrics.

MetricDescription

`CallCount`

The number of specified API operations performed in your account.

`CallCount` is a CloudWatch Logs service usage metric. For more information,
see [CloudWatch Logs service usage metrics](#CloudWatchLogs-Usage-Metrics).

Valid Dimensions: Class, Resource, Service, Type

Valid Statistic: Sum

Units: None

`DeliveryErrors`

The number of log events for which CloudWatch Logs received an error when forwarding data to the
subscription destination. If the destination service returns a retryable error such as a
throttling exception or a retryable service exception (HTTP 5xx for example), CloudWatch Logs
continues to retry delivery for up to 24 hours. CloudWatch Logs does not try to re-deliver if
the error is a non-retryable error, such as `AccessDeniedException` or
`ResourceNotFoundException`.

Valid Dimensions: LogGroupName, DestinationType, FilterName, PolicyLevel

Valid Statistic: Sum

Units: None

`DeliveryThrottling`

The number of log events for which CloudWatch Logs was throttled when
forwarding data to the subscription destination.

If the destination service returns a retryable error such as a
throttling exception or a retryable service exception (HTTP 5xx for example), CloudWatch Logs
continues to retry delivery for up to 24 hours. CloudWatch Logs does not try to re-deliver if
the error is a non-retryable error, such as `AccessDeniedException` or
`ResourceNotFoundException`.

Valid Dimensions: LogGroupName, DestinationType, FilterName, PolicyLevel

Valid Statistic: Sum

Units: None

`EMFDisabledErrors`

The number of EMF-formatted log events that were ignored due to an
active account policy of type METRIC\_EXTRACTION\_POLICY that matches
the log group.For more information about the metric extraction
policies, see [PutAccountPolicy](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-putaccountpolicy.md).

Valid Dimensions: `LogGroupName`

Valid Statistic: Sum

Units: None

`EMFParsingErrors`

The number
of parsing errors
encountered
while processing embedded metric format logs.
Such errors happen
when logs are identified
as embedded metric format but do not follow the correct format.
For more information about the embedded metric format, see
[Specification: Embedded metric format](../monitoring/cloudwatch-embedded-metric-format-specification.md).

Valid Dimensions: `LogGroupName`

Valid Statistic: Sum

Units: None

`EMFValidationErrors`

The number
of validation errors
encountered
while processing embedded metric format logs.
These errors occur
when metric definitions
within embedded metric format logs do not adhere
to the embedded metric format and `MetricDatum` specifications.
For information
about the CloudWatch embedded metric format,
see [Specification: Embedded metric format](../monitoring/cloudwatch-embedded-metric-format-specification.md).
For information
about the data type `MetricDatum`,
see [MetricDatum](../../../../reference/amazoncloudwatch/latest/apireference/api-metricdatum.md)
in the _Amazon CloudWatch API Reference_.

###### Note

Certain validation errors can lead
to multiple metrics
within an EMF log not being published.
For example,
all metrics set with an invalid namespace will be dropped.

Valid Dimensions: `LogGroupName`

Valid Statistic: Sum

Units: None

`ErrorCount`

The number of API operations performed in your account that resulted in errors.

`ErrorCount` is a CloudWatch Logs service usage metric. For more information,
see [CloudWatch Logs service usage metrics](#CloudWatchLogs-Usage-Metrics).

Valid Dimensions: Class, Resource, Service, Type

Valid Statistic: Sum

Units: None

`ForwardedBytes`

The volume of log events in compressed bytes forwarded to the
subscription destination.

Valid Dimensions: LogGroupName, DestinationType, FilterName

Valid Statistic: Sum

Units: Bytes

`ForwardedLogEvents`

The number of log events forwarded to the subscription
destination.

Valid Dimensions: LogGroupName, DestinationType, FilterName, PolicyLevel

Valid Statistic: Sum

Units: None

`IncomingBytes`

The volume of log events in uncompressed bytes uploaded to CloudWatch Logs.
When used with the `LogGroupName` dimension, this is the
volume of log events in uncompressed bytes uploaded to the log
group.

Valid Dimensions: LogGroupName

Valid Statistic: Sum

Units: Bytes

`IncomingLogEvents`

The number of log events uploaded to CloudWatch Logs. When used with the
`LogGroupName` dimension, this is the number of log
events uploaded to the log group.

Valid Dimensions: LogGroupName

Valid Statistic: Sum

Units: None

`LogEventsWithFindings`

The number of log events that matched a data string
that you are auditing using the CloudWatch Logs data protection feature. For more
information, see [Help protect sensitive log data with masking](mask-sensitive-log-data.md).

Valid Dimensions: None

Valid Statistic: Sum

Units: None

`ThrottleCount`

The number of API operations performed in your account that
were throttled because of usage quotas.

`ThrottleCount` is a CloudWatch Logs service usage metric. For more information,
see [CloudWatch Logs service usage metrics](#CloudWatchLogs-Usage-Metrics).

Valid Dimensions: Class, Resource, Service, Type

Valid Statistic: Sum

Units: None

## Dimensions for CloudWatch Logs metrics

The dimensions that you can use with most CloudWatch Logs metrics are listed in the following table.

Dimension

Description
`LogGroupName`

The name of the CloudWatch Logs log group for which to display
metrics.

`DestinationType`

The subscription destination for the CloudWatch Logs data, which can be
AWS Lambda, Amazon Kinesis Data Streams, or Amazon Data Firehose.

`FilterName`

The name of the subscription filter that is forwarding data from
the log group to the destination. The subscription filter name is
automatically converted by CloudWatch to ASCII and any unsupported
characters get replaced with a question mark (?).

**Subscription filter metric dimensions**

The dimensions for metrics related to account-level subscription filters are listed in the following table.

Dimension

Description
`PolicyLevel`

The level where the policy applies. Currently, the only valid value for this dimension
is `AccountPolicy`

`DestinationType`

The subscription destination for the CloudWatch Logs data, which can be
AWS Lambda, Amazon Kinesis Data Streams, or Amazon Data Firehose.

`FilterName`

The name of the subscription filter that is forwarding data from
the log group to the destination. The subscription filter name is
automatically converted by CloudWatch to ASCII and any unsupported
characters get replaced with a question mark (?).

## Log transformer metrics and dimensions

CloudWatch Logs publishes the following log transformer metrics to CloudWatch in the `AWS/Logs` namespace.

Metric

Description
`TransformationErrors`

The number of errors encountered while transforming log events with the specified transformer.

Unit: None

Valid statistic: Sum

`TransformedBytes`

The volume of the output of transformed log events, in uncompressed bytes.

Unit: Bytes

Valid statistic: Sum

`TransformedLogEvents`

The number of transformed log events.

Unit: None

Valid statistic: Sum

The following dimensions are used by transformer metrics.

Dimension

Description
`LogGroupname`

This dimension is used only for log-group-level transformers.

`PolicyLevel`

This dimension is used only for account-level transformers. Currently the only valid value for this dimension is `AccountPolicy`

## Centralization metrics and dimensions

For centralized monitoring across multiple accounts and regions, you can use CloudWatch Logs Centralization
to consolidate log data and metrics in a central location. For more information, see
[Cross-account cross-Region log centralization](cloudwatchlogs-centralization.md).

CloudWatch Logs publishes the following centralization metrics to CloudWatch in the `AWS/Logs` namespace.
These metrics help you monitor the replication of log data from source accounts to destination accounts
when using CloudWatch Logs centralization rules.

###### Note

CloudWatch begins reporting centralization metrics shortly after you create a centralization rule.
Metrics are published on a best-effort basis and track only new log events that arrive after
the centralization rule is created.

Metric

Description

Published in
`IncomingCopiedBytes`

The volume of log data in uncompressed bytes that was replicated into the destination account.
This metric applies only to new log events that arrive after the centralization rule is created.

Valid Dimensions: SourceRegion, SourceAccount

Valid Statistic: Sum

Units: Bytes

Destination account

`IncomingCopiedLogEvents`

The number of log events that were replicated into the destination account.
This metric applies only to new log events that arrive after the centralization rule is created.

Valid Dimensions: SourceRegion, SourceAccount

Valid Statistic: Sum

Units: None

Destination account

`OutgoingCopiedBytes`

The volume of log data in uncompressed bytes that was sent from source accounts
to the destination account. This metric applies only to new log events that are copied to the destination
after the centralization rule is created.

Valid Dimensions: DestinationRegion

Valid Statistic: Sum

Units: Bytes

Source account

`OutgoingCopiedLogEvents`

The number of log events that were sent from source accounts
to the destination account. This metric applies only to new log events that are copied to the destination
after the centralization rule is created.

Valid Dimensions: DestinationRegion

Valid Statistic: Sum

Units: None

Source account

`CentralizationError`

The number of errors encountered while replicating log data.
Errors can occur due to various reasons such as KMS key permission issues, log group quota limits,
or log tier mismatches. To identify specific log groups or log streams that failed replication
and their failure reasons, monitor the `ErrorType` dimension.

Valid Dimensions: ErrorType

Valid Statistic: Sum

Units: None

Destination account

`CentralizationThrottled`

The number of times centralization processing was throttled.
Throttling results in slower processing of centralization but does not prevent
log data from being replicated.

Valid Dimensions: DestinationRegion

Valid Statistic: Sum

Units: None

Source account

The following dimensions are used by centralization metrics.

Dimension

Description
`SourceRegion`

The AWS Region where the source log data originated.

`SourceAccount`

The AWS account ID where the source log data originated.

`DestinationRegion`

The AWS Region where the log data is being replicated to.

`ErrorType`

The type of error encountered during centralization. Possible values include:

- `LogGroupQuotaExceeded`: The destination account has reached its log group quota.

- `InvalidKMS`: The KMS key is missing, deleted, not enabled, or is asymmetric instead of symmetric.

- `AccessDenied`: Access was denied when attempting to replicate log data.
This can occur due to insufficient permissions, such as incorrect KMS key permissions,
missing IAM permissions, or restrictive resource policies.

- `LogTierMismatch`: The log tier of the source and destination log groups do not match.

- `InvalidLogStream`: Invalid parameter encountered when creating or writing to a log stream,
such as when the log stream name exceeds the maximum length.

- `InvalidLogGroup`: Invalid parameter encountered when creating or configuring a log group
in the destination account.

- `DestinationEncryptionMismatch`: The destination log group already exists with a different
KMS encryption configuration than what the centralization rule specifies.

## CloudWatch Logs service usage metrics

CloudWatch Logs sends metrics to CloudWatch that track the usage of CloudWatch Logs API operations.
These metrics correspond to AWS service quotas.
Tracking these metrics can help you proactively manage your quotas. For more information, see
[Service Quotas Integration and Usage Metrics](../monitoring/cloudwatch-service-quota-integration.md).

For example, you could track the `ThrottleCount` metric or set an alarm
on that metric. If the value of this metric rises, you should consider requesting
a quota increase for the API operation that is getting throttled. For more information
about CloudWatch Logs service quotas, see [CloudWatch Logs quotas](cloudwatch-limits-cwl.md).

CloudWatch Logs publishes service quota usage metrics every minute in both the `AWS/Usage`
and `AWS/Logs`
namespaces.

The following table lists the service usage metrics published by CloudWatch Logs. These
metrics do not have a specified unit. The most useful statistic
for these metrics is `SUM`, which represents the total operation count for the 1-minute
period.

Each of these metrics is published with values for all of the `Service`,
`Class`, `Type`, and `Resource` dimensions. They are
also published with a single dimension called `Account Metrics`. Use the
`Account Metrics` dimension to see the sum of metrics for all
API operations in your account. Use the other dimensions and specify the name of
an API operation for the `Resource` dimension to find the metrics for
that particular API.

**Metrics**

MetricDescription

`CallCount`

The number of specified operations performed in your account.

`CallCount` is published in both the `AWS/Usage`
and `AWS/Logs`
namespaces.

`ErrorCount`

The number of API operations performed in your account that resulted in errors.

`ErrorCount` is published in only the `AWS/Logs`.

`ThrottleCount`

The number of API operations performed in your account that
were throttled because of usage quotas.

`ThrottleCount` is published in only the `AWS/Logs`.

**Dimensions**

DimensionDescription

`Account metrics`

Use this dimension to get a sum of the metric across all of
the CloudWatch Logs APIs.

If you want to see the metrics for one particular API,
use the other dimensions listed in this table and specify the
API name as the value of `Resource`.

`Service`

The name of the AWS service containing the resource. For CloudWatch Logs usage metrics, the value
for this dimension is `Logs`.

`Class`

The class of resource being tracked. CloudWatch Logs API usage metrics use this
dimension with a value of `None`.

`Type`

The type of resource being tracked. Currently, when the `Service` dimension
is `Logs`,
the only valid value for `Type` is `API`.

`Resource`

The name of the API operation. Valid values include all of the API operation names
that are listed in
[Actions](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-operations.md). For example, `PutLogEvents`

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Logging API and console operations with AWS CloudTrail

Service quotas

All content copied from https://docs.aws.amazon.com/.
