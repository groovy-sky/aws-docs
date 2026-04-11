This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cognito::LogDeliveryConfiguration FirehoseConfiguration

Configuration for the Amazon Data Firehose stream destination of user activity log export with
threat protection.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "StreamArn" : String
}

```

### YAML

```yaml

  StreamArn: String

```

## Properties

`StreamArn`

The ARN of an Amazon Data Firehose stream that's the destination for threat protection log
export.

_Required_: No

_Type_: String

_Pattern_: `arn:[\w+=/,.@-]+:[\w+=/,.@-]+:([\w+=/,.@-]*)?:[0-9]+:[\w+=/,.@-]+(:[\w+=/,.@-]+)?(:[\w+=/,.@-]+)?`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CloudWatchLogsConfiguration

LogConfiguration

All content copied from https://docs.aws.amazon.com/.
