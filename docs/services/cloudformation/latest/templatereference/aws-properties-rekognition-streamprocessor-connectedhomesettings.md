This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Rekognition::StreamProcessor ConnectedHomeSettings

Connected home settings to use on a streaming video. Defining the settings is required in the request parameter for `CreateStreamProcessor`.
Including this setting in the CreateStreamProcessor request lets you use the stream processor for connected home features. You can then select
what you want the stream processor to detect, such as people or pets.

When the stream processor has started, one notification is sent
for each object class specified. For example, if packages and pets are selected, one SNS notification is published the first time a package is
detected and one SNS notification is published the first time a pet is detected. An end-of-session summary is also published.
For more information, see the ConnectedHome section of [StreamProcessorSettings](https://docs.aws.amazon.com/rekognition/latest/APIReference/API_StreamProcessorSettings).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Labels" : [ String, ... ],
  "MinConfidence" : Number
}

```

### YAML

```yaml

  Labels:
    - String
  MinConfidence: Number

```

## Properties

`Labels`

Specifies what you want to detect in the video, such as people, packages, or pets.
The current valid labels you can include in this list are: "PERSON", "PET", "PACKAGE", and "ALL".

_Required_: Yes

_Type_: Array of String

_Maximum_: `128`

_Minimum_: `1 | 1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MinConfidence`

The minimum confidence required to label an object in the video.

_Required_: No

_Type_: Number

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

BoundingBox

DataSharingPreference
