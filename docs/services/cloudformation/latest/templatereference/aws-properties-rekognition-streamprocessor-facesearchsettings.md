This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Rekognition::StreamProcessor FaceSearchSettings

The input parameters used to recognize faces in a streaming video analyzed by a Amazon Rekognition stream processor. `FaceSearchSettings` is a request
parameter for [CreateStreamProcessor](https://docs.aws.amazon.com/rekognition/latest/APIReference/API_CreateStreamProcessor).
For more information, see [FaceSearchSettings](https://docs.aws.amazon.com/rekognition/latest/APIReference/API_FaceSearchSettings).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CollectionId" : String,
  "FaceMatchThreshold" : Number
}

```

### YAML

```yaml

  CollectionId: String
  FaceMatchThreshold: Number

```

## Properties

`CollectionId`

The ID of a collection that contains faces that you want to search for.

_Required_: Yes

_Type_: String

_Pattern_: `\A[a-zA-Z0-9_\.\-]+$`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FaceMatchThreshold`

Minimum face match confidence score that must be met to return a result for a recognized face. The default is 80.
0 is the lowest confidence. 100 is the highest confidence. Values between 0 and 100 are accepted, and values lower than 80 are set to 80.

_Required_: No

_Type_: Number

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DataSharingPreference

KinesisDataStream
