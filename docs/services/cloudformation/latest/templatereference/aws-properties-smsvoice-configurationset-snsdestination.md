This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SMSVOICE::ConfigurationSet SnsDestination

An object that defines an Amazon SNS destination for events. You can use
Amazon SNS to send notification when certain events occur.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "TopicArn" : String
}

```

### YAML

```yaml

  TopicArn: String

```

## Properties

`TopicArn`

The Amazon Resource Name (ARN) of the Amazon SNS topic that you want to
publish events to.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:\S+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

KinesisFirehoseDestination

Tag

All content copied from https://docs.aws.amazon.com/.
