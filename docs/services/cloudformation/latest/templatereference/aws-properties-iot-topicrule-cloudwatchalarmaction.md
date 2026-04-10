This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::TopicRule CloudwatchAlarmAction

Describes an action that updates a CloudWatch alarm.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AlarmName" : String,
  "RoleArn" : String,
  "StateReason" : String,
  "StateValue" : String
}

```

### YAML

```yaml

  AlarmName: String
  RoleArn: String
  StateReason: String
  StateValue: String

```

## Properties

`AlarmName`

The CloudWatch alarm name.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The IAM role that allows access to the CloudWatch alarm.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StateReason`

The reason for the alarm change.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StateValue`

The value of the alarm state. Acceptable values are: OK, ALARM,
INSUFFICIENT\_DATA.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BatchConfig

CloudwatchLogsAction

All content copied from https://docs.aws.amazon.com/.
