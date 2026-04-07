This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Rekognition::StreamProcessor DataSharingPreference

Allows you to opt in or opt out to share data with Rekognition to improve model performance.
You can choose this option at the account level or on a per-stream basis. Note that if you opt out at the account level, this setting is ignored on individual streams.
For more information, see [StreamProcessorDataSharingPreference](https://docs.aws.amazon.com/rekognition/latest/APIReference/API_StreamProcessorDataSharingPreference).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "OptIn" : Boolean
}

```

### YAML

```yaml

  OptIn: Boolean

```

## Properties

`OptIn`

Describes the opt-in status applied to a stream processor's data sharing policy.

_Required_: Yes

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ConnectedHomeSettings

FaceSearchSettings
