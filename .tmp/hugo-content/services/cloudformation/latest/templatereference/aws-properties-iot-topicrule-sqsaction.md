This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::TopicRule SqsAction

Describes an action to publish data to an Amazon SQS queue.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "QueueUrl" : String,
  "RoleArn" : String,
  "UseBase64" : Boolean
}

```

### YAML

```yaml

  QueueUrl: String
  RoleArn: String
  UseBase64: Boolean

```

## Properties

`QueueUrl`

The URL of the Amazon SQS queue.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The ARN of the IAM role that grants access.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseBase64`

Specifies whether to use Base64 encoding.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SnsAction

StepFunctionsAction

All content copied from https://docs.aws.amazon.com/.
