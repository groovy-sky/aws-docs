This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KafkaConnect::Connector CloudWatchLogsLogDelivery

The settings for delivering connector logs to Amazon CloudWatch Logs.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Enabled" : Boolean,
  "LogGroup" : String
}

```

### YAML

```yaml

  Enabled: Boolean
  LogGroup: String

```

## Properties

`Enabled`

Whether log delivery to Amazon CloudWatch Logs is enabled.

_Required_: Yes

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LogGroup`

The name of the CloudWatch log group that is the destination for log delivery.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Capacity

CustomPlugin

All content copied from https://docs.aws.amazon.com/.
