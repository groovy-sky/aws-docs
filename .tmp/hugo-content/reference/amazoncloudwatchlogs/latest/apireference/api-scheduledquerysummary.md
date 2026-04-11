# ScheduledQuerySummary

Summary information about a scheduled query, including basic configuration and execution
status.

## Contents

**creationTime**

The timestamp when the scheduled query was created.

Type: Long

Valid Range: Minimum value of 0.

Required: No

**destinationConfiguration**

Configuration for where query results are delivered.

Type: [DestinationConfiguration](api-destinationconfiguration.md) object

Required: No

**lastExecutionStatus**

The status of the most recent execution.

Type: String

Valid Values: `Running | InvalidQuery | Complete | Failed | Timeout`

Required: No

**lastTriggeredTime**

The timestamp when the scheduled query was last executed.

Type: Long

Valid Range: Minimum value of 0.

Required: No

**lastUpdatedTime**

The timestamp when the scheduled query was last updated.

Type: Long

Valid Range: Minimum value of 0.

Required: No

**name**

The name of the scheduled query.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `^[a-zA-Z0-9_\-/.#]+$`

Required: No

**scheduledQueryArn**

The ARN of the scheduled query.

Type: String

Required: No

**scheduleExpression**

The cron expression that defines when the scheduled query runs.

Type: String

Length Constraints: Maximum length of 256.

Required: No

**state**

The current state of the scheduled query.

Type: String

Valid Values: `ENABLED | DISABLED`

Required: No

**timezone**

The timezone used for evaluating the schedule expression.

Type: String

Length Constraints: Minimum length of 1.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/scheduledquerysummary.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/scheduledquerysummary.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/scheduledquerysummary.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ScheduledQueryDestination

SearchedLogStream

All content copied from https://docs.aws.amazon.com/.
