This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisFirehose::DeliveryStream KinesisStreamSourceConfiguration

The `KinesisStreamSourceConfiguration` property type specifies the stream and
role Amazon Resource Names (ARNs) for a Kinesis stream used as the source for a delivery
stream.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "KinesisStreamARN" : String,
  "RoleARN" : String
}

```

### YAML

```yaml

  KinesisStreamARN: String
  RoleARN: String

```

## Properties

`KinesisStreamARN`

The ARN of the source Kinesis data stream.

_Required_: Yes

_Type_: String

_Pattern_: `arn:.*`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RoleARN`

The ARN of the role that provides access to the source Kinesis data stream.

_Required_: Yes

_Type_: String

_Pattern_: `arn:.*`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InputFormatConfiguration

KMSEncryptionConfig

All content copied from https://docs.aws.amazon.com/.
