This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::TopicRule SnsAction

Describes an action to publish to an Amazon SNS topic.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MessageFormat" : String,
  "RoleArn" : String,
  "TargetArn" : String
}

```

### YAML

```yaml

  MessageFormat: String
  RoleArn: String
  TargetArn: String

```

## Properties

`MessageFormat`

(Optional) The message format of the message to publish. Accepted values are "JSON"
and "RAW". The default value of the attribute is "RAW". SNS uses this setting to determine
if the payload should be parsed and relevant platform-specific bits of the payload should
be extracted. For more information, see [Amazon SNS Message and JSON Formats](../../../sns/latest/dg/json-formats.md) in the
_Amazon Simple Notification Service Developer Guide_.

_Required_: No

_Type_: String

_Allowed values_: `RAW | JSON`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The ARN of the IAM role that grants access.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetArn`

The ARN of the SNS topic.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SigV4Authorization

SqsAction

All content copied from https://docs.aws.amazon.com/.
