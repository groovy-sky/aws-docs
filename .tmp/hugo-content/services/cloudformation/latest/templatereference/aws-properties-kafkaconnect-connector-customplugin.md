This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KafkaConnect::Connector CustomPlugin

A plugin is an AWS resource that contains the code that defines a connector's logic.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CustomPluginArn" : String,
  "Revision" : Integer
}

```

### YAML

```yaml

  CustomPluginArn: String
  Revision: Integer

```

## Properties

`CustomPluginArn`

The Amazon Resource Name (ARN) of the custom plugin.

_Required_: Yes

_Type_: String

_Pattern_: `arn:(aws|aws-us-gov|aws-cn):kafkaconnect:.*`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Revision`

The revision of the custom plugin.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CloudWatchLogsLogDelivery

FirehoseLogDelivery

All content copied from https://docs.aws.amazon.com/.
