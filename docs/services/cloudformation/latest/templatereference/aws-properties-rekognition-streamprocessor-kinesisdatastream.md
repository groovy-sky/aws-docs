This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Rekognition::StreamProcessor KinesisDataStream

Amazon Rekognition Video Stream Processor take as input a Kinesis video stream (Input) and a Kinesis data stream (Output).
This is the Amazon Kinesis Data Streams instance to which the Amazon Rekognition stream processor streams the analysis results.
This must be created within the constraints specified at
[KinesisDataStream](https://docs.aws.amazon.com/rekognition/latest/APIReference/API_KinesisDataStream).

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

ARN of the output Amazon Kinesis Data Streams stream.

_Required_: Yes

_Type_: String

_Pattern_: `(^arn:([a-z\d-]+):kinesis:([a-z\d-]+):\d{12}:.+$)`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

FaceSearchSettings

KinesisVideoStream
