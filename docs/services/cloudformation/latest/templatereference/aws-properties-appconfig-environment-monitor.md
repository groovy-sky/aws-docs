This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppConfig::Environment Monitor

Amazon CloudWatch alarms to monitor during the deployment process.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AlarmArn" : String,
  "AlarmRoleArn" : String
}

```

### YAML

```yaml

  AlarmArn: String
  AlarmRoleArn: String

```

## Properties

`AlarmArn`

Amazon Resource Name (ARN) of the Amazon CloudWatch alarm.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AlarmRoleArn`

ARN of an AWS Identity and Access Management (IAM) role for AWS AppConfig to monitor
`AlarmArn`.

_Required_: No

_Type_: String

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::AppConfig::Environment

Tag

All content copied from https://docs.aws.amazon.com/.
