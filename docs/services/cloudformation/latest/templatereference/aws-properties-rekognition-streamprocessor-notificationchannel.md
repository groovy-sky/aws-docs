This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Rekognition::StreamProcessor NotificationChannel

The Amazon Simple Notification Service topic to which Amazon Rekognition publishes the object detection results and completion status of a video analysis operation.
Amazon Rekognition publishes a notification the first time an object of interest or a person is detected in the video stream.
Amazon Rekognition also publishes an an end-of-session notification with a summary when the stream processing session is complete.
For more information, see [StreamProcessorNotificationChannel](../../../../reference/rekognition/latest/apireference/api-streamprocessornotificationchannel.md).

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

The ARN of the SNS topic that receives notifications.

_Required_: Yes

_Type_: String

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

KinesisVideoStream

Point

All content copied from https://docs.aws.amazon.com/.
