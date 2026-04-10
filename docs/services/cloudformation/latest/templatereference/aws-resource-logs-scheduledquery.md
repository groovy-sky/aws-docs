This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Logs::ScheduledQuery

Creates a scheduled query that runs CloudWatch Logs Insights queries at regular intervals.
Scheduled queries enable proactive monitoring by automatically executing queries to detect
patterns and anomalies in your log data. Query results can be delivered to Amazon S3 for analysis
or further processing.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Logs::ScheduledQuery",
  "Properties" : {
      "Description" : String,
      "DestinationConfiguration" : DestinationConfiguration,
      "ExecutionRoleArn" : String,
      "LogGroupIdentifiers" : [ String, ... ],
      "Name" : String,
      "QueryLanguage" : String,
      "QueryString" : String,
      "ScheduleEndTime" : Number,
      "ScheduleExpression" : String,
      "ScheduleStartTime" : Number,
      "StartTimeOffset" : Integer,
      "State" : String,
      "Tags" : [ TagsItems, ... ],
      "Timezone" : String
    }
}

```

### YAML

```yaml

Type: AWS::Logs::ScheduledQuery
Properties:
  Description: String
  DestinationConfiguration:
    DestinationConfiguration
  ExecutionRoleArn: String
  LogGroupIdentifiers:
    - String
  Name: String
  QueryLanguage: String
  QueryString:
    String
  ScheduleEndTime: Number
  ScheduleExpression: String
  ScheduleStartTime: Number
  StartTimeOffset: Integer
  State: String
  Tags:
    - TagsItems
  Timezone: String

```

## Properties

`Description`

Property description not available.

_Required_: No

_Type_: String

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DestinationConfiguration`

Configuration for where query results are delivered.

_Required_: No

_Type_: [DestinationConfiguration](aws-properties-logs-scheduledquery-destinationconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExecutionRoleArn`

Property description not available.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogGroupIdentifiers`

Property description not available.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the scheduled query.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_\-/.#]+$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`QueryLanguage`

Property description not available.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`QueryString`

Property description not available.

_Required_: Yes

_Type_: String

_Minimum_: `0`

_Maximum_: `10000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScheduleEndTime`

Property description not available.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScheduleExpression`

The cron expression that defines when the scheduled query runs.

_Required_: Yes

_Type_: String

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScheduleStartTime`

Property description not available.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StartTimeOffset`

Property description not available.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`State`

The current state of the scheduled query.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Property description not available.

_Required_: No

_Type_: Array of [TagsItems](aws-properties-logs-scheduledquery-tagsitems.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Timezone`

The timezone used for evaluating the schedule expression.

_Required_: No

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`CreationTime`

The timestamp when the scheduled query was created.

`LastExecutionStatus`

The status of the most recent execution.

`LastTriggeredTime`

The timestamp when the scheduled query was last executed.

`LastUpdatedTime`

The timestamp when the scheduled query was last updated.

`ScheduledQueryArn`

The ARN of the scheduled query.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Logs::ResourcePolicy

DestinationConfiguration

All content copied from https://docs.aws.amazon.com/.
