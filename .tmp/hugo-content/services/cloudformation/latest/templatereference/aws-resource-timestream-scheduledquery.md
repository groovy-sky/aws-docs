This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Timestream::ScheduledQuery

Create a scheduled query that will be run on your behalf at the configured schedule.
Timestream assumes the execution role provided as part of the
`ScheduledQueryExecutionRoleArn` parameter to run the query. You can use the
`NotificationConfiguration` parameter to configure notification for your
scheduled query operations.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Timestream::ScheduledQuery",
  "Properties" : {
      "ClientToken" : String,
      "ErrorReportConfiguration" : ErrorReportConfiguration,
      "KmsKeyId" : String,
      "NotificationConfiguration" : NotificationConfiguration,
      "QueryString" : String,
      "ScheduleConfiguration" : ScheduleConfiguration,
      "ScheduledQueryExecutionRoleArn" : String,
      "ScheduledQueryName" : String,
      "Tags" : [ Tag, ... ],
      "TargetConfiguration" : TargetConfiguration
    }
}

```

### YAML

```yaml

Type: AWS::Timestream::ScheduledQuery
Properties:
  ClientToken: String
  ErrorReportConfiguration:
    ErrorReportConfiguration
  KmsKeyId: String
  NotificationConfiguration:
    NotificationConfiguration
  QueryString:
    String
  ScheduleConfiguration:
    ScheduleConfiguration
  ScheduledQueryExecutionRoleArn: String
  ScheduledQueryName: String
  Tags:
    - Tag
  TargetConfiguration:
    TargetConfiguration

```

## Properties

`ClientToken`

Using a ClientToken makes the call to CreateScheduledQuery idempotent, in other words,
making the same request repeatedly will produce the same result. Making multiple identical
CreateScheduledQuery requests has the same effect as making a single request.

- If CreateScheduledQuery is called without a `ClientToken`, the Query SDK
generates a `ClientToken` on your behalf.

- After 8 hours, any request with the same `ClientToken` is treated as a new
request.

_Required_: No

_Type_: String

_Minimum_: `32`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ErrorReportConfiguration`

Configuration for error reporting. Error reports will be generated when a problem is
encountered when writing the query results.

_Required_: Yes

_Type_: [ErrorReportConfiguration](aws-properties-timestream-scheduledquery-errorreportconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`KmsKeyId`

The Amazon KMS key used to encrypt the scheduled query resource, at-rest. If the Amazon
KMS key is not specified, the scheduled query resource will be encrypted with a Timestream
owned Amazon KMS key. To specify a KMS key, use the key ID, key ARN, alias name, or alias ARN.
When using an alias name, prefix the name with _alias/_

If ErrorReportConfiguration uses `SSE_KMS` as encryption type, the same
KmsKeyId is used to encrypt the error report at rest.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NotificationConfiguration`

Notification configuration for the scheduled query. A notification is sent by Timestream
when a query run finishes, when the state is updated or when you delete it.

_Required_: Yes

_Type_: [NotificationConfiguration](aws-properties-timestream-scheduledquery-notificationconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`QueryString`

The query string to run. Parameter names can be specified in the query string
`@` character followed by an identifier. The named Parameter
`@scheduled_runtime` is reserved and can be used in the query to get the time at
which the query is scheduled to run.

The timestamp calculated according to the ScheduleConfiguration parameter, will be the
value of `@scheduled_runtime` paramater for each query run. For example, consider
an instance of a scheduled query executing on 2021-12-01 00:00:00. For this instance, the
`@scheduled_runtime` parameter is initialized to the timestamp 2021-12-01
00:00:00 when invoking the query.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `262144`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ScheduleConfiguration`

Schedule configuration.

_Required_: Yes

_Type_: [ScheduleConfiguration](aws-properties-timestream-scheduledquery-scheduleconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ScheduledQueryExecutionRoleArn`

The ARN for the IAM role that Timestream will assume when running the scheduled query.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ScheduledQueryName`

A name for the query. Scheduled query names must be unique within each Region.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9_.-]+`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

A list of key-value pairs to label the scheduled query.

_Required_: No

_Type_: Array of [Tag](aws-properties-timestream-scheduledquery-tag.md)

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetConfiguration`

Scheduled query target store configuration.

_Required_: No

_Type_: [TargetConfiguration](aws-properties-timestream-scheduledquery-targetconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the scheduled query ARN.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` returns a value for the specified attribute of this type. The
following are the available attributes:

`Arn`

The `ARN` of the scheduled query.

`SQErrorReportConfiguration`

The scheduled query error reporting configuration.

`SQKmsKeyId`

The KMS key used to encrypt the query resource, if a customer managed KMS key was
provided.

`SQName`

The scheduled query name.

`SQNotificationConfiguration`

The scheduled query notification configuration.

`SQQueryString`

The scheduled query string..

`SQScheduleConfiguration`

The scheduled query schedule configuration.

`SQScheduledQueryExecutionRoleArn`

The ARN of the IAM role that will be used by Timestream to run the query.

`SQTargetConfiguration`

The configuration for query output.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

DimensionMapping

All content copied from https://docs.aws.amazon.com/.
