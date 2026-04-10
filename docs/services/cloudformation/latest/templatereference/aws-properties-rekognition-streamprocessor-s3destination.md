This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Rekognition::StreamProcessor S3Destination

The Amazon S3 bucket location to which Amazon Rekognition publishes the detailed inference results of a video analysis operation.
These results include the name of the stream processor resource, the session ID of the stream processing session,
and labeled timestamps and bounding boxes for detected labels. For more information, see
[S3Destination](../../../../reference/rekognition/latest/apireference/api-s3destination.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BucketName" : String,
  "ObjectKeyPrefix" : String
}

```

### YAML

```yaml

  BucketName: String
  ObjectKeyPrefix: String

```

## Properties

`BucketName`

Describes the destination Amazon Simple Storage Service (Amazon S3) bucket name of a stream processor's exports.

_Required_: Yes

_Type_: String

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ObjectKeyPrefix`

Describes the destination Amazon Simple Storage Service (Amazon S3) object keys of a stream processor's exports.

_Required_: No

_Type_: String

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Point

Tag

All content copied from https://docs.aws.amazon.com/.
