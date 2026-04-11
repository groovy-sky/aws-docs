This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::SecurityProfile AlertTarget

A structure containing the alert target ARN and the role ARN.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AlertTargetArn" : String,
  "RoleArn" : String
}

```

### YAML

```yaml

  AlertTargetArn: String
  RoleArn: String

```

## Properties

`AlertTargetArn`

The Amazon Resource Name (ARN) of the notification target to which alerts are sent.

_Required_: Yes

_Type_: String

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The ARN of the role that grants permission to send alerts to the
notification target.

_Required_: Yes

_Type_: String

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::IoT::SecurityProfile

Behavior

All content copied from https://docs.aws.amazon.com/.
