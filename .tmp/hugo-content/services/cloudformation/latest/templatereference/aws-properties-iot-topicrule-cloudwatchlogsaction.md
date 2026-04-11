This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::TopicRule CloudwatchLogsAction

Describes an action that updates a CloudWatch log.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BatchMode" : Boolean,
  "LogGroupName" : String,
  "RoleArn" : String
}

```

### YAML

```yaml

  BatchMode: Boolean
  LogGroupName: String
  RoleArn: String

```

## Properties

`BatchMode`

Indicates whether batches of log records will be extracted and uploaded into
CloudWatch.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogGroupName`

The CloudWatch log name.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The IAM role that allows access to the CloudWatch log.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CloudwatchAlarmAction

CloudwatchMetricAction

All content copied from https://docs.aws.amazon.com/.
