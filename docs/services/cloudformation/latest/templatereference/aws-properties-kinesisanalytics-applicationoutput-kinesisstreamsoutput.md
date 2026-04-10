This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisAnalytics::ApplicationOutput KinesisStreamsOutput

When configuring application output, identifies an Amazon Kinesis stream as the
destination. You provide the stream Amazon Resource Name (ARN) and also an IAM role ARN
that Amazon Kinesis Analytics can use to write to the stream on your behalf.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ResourceARN" : String,
  "RoleARN" : String
}

```

### YAML

```yaml

  ResourceARN: String
  RoleARN: String

```

## Properties

`ResourceARN`

ARN of the destination Amazon Kinesis stream to write to.

_Required_: Yes

_Type_: String

_Pattern_: `arn:.*`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleARN`

ARN of the IAM role that Amazon Kinesis Analytics can assume to write to the
destination stream on your behalf. You need to grant the necessary permissions to this
role.

_Required_: Yes

_Type_: String

_Pattern_: `arn:.*`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

KinesisFirehoseOutput

LambdaOutput

All content copied from https://docs.aws.amazon.com/.
