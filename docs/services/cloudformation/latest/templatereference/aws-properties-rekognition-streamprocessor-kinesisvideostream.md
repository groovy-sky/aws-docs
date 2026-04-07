This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Rekognition::StreamProcessor KinesisVideoStream

The Kinesis video stream that provides the source of the streaming video for an Amazon Rekognition Video stream processor. For more information, see
[KinesisVideoStream](https://docs.aws.amazon.com/rekognition/latest/APIReference/API_KinesisVideoStream).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Arn" : String
}

```

### YAML

```yaml

  Arn: String

```

## Properties

`Arn`

ARN of the Kinesis video stream stream that streams the source video.

_Required_: Yes

_Type_: String

_Pattern_: `(^arn:([a-z\d-]+):kinesisvideo:([a-z\d-]+):\d{12}:.+$)`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

KinesisDataStream

NotificationChannel
