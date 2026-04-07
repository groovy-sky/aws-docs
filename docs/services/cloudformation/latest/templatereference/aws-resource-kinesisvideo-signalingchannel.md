This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisVideo::SignalingChannel

Specifies a signaling channel.

`CreateSignalingChannel` is an asynchronous operation.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::KinesisVideo::SignalingChannel",
  "Properties" : {
      "MessageTtlSeconds" : Integer,
      "Name" : String,
      "Tags" : [ Tag, ... ],
      "Type" : String
    }
}

```

### YAML

```yaml

Type: AWS::KinesisVideo::SignalingChannel
Properties:
  MessageTtlSeconds: Integer
  Name: String
  Tags:
    - Tag
  Type: String

```

## Properties

`MessageTtlSeconds`

The period of time (in seconds) a signaling channel retains undelivered messages
before they are discarded. Use API\_UpdateSignalingChannel to update
this value.

_Required_: No

_Type_: Integer

_Minimum_: `5`

_Maximum_: `120`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

A name for the signaling channel that you are creating. It must be unique for each AWS account and AWS Region.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9_.-]+`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kinesisvideo-signalingchannel-tag.html)

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

A type of the signaling channel that you are creating. Currently,
`SINGLE_MASTER` is the only supported channel type.

_Required_: No

_Type_: String

_Allowed values_: `SINGLE_MASTER`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`Arn`

The Amazon Resource Name (ARN) of the signaling channel.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Amazon Kinesis Video Streams

Tag
