This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTAnalytics::Dataset IotEventsDestinationConfiguration

Configuration information for delivery of dataset contents to AWS IoT Events.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "InputName" : String,
  "RoleArn" : String
}

```

### YAML

```yaml

  InputName: String
  RoleArn: String

```

## Properties

`InputName`

The name of the AWS IoT Events input to which dataset contents are delivered.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z][a-zA-Z0-9_]*$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The ARN of the role that grants AWS IoT Analytics permission to deliver dataset contents to an AWS IoT Events
input.

_Required_: Yes

_Type_: String

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GlueConfiguration

LateDataRule

All content copied from https://docs.aws.amazon.com/.
