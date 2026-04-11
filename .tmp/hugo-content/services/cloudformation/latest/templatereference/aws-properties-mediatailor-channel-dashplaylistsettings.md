This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaTailor::Channel DashPlaylistSettings

Dash manifest configuration parameters.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ManifestWindowSeconds" : Number,
  "MinBufferTimeSeconds" : Number,
  "MinUpdatePeriodSeconds" : Number,
  "SuggestedPresentationDelaySeconds" : Number
}

```

### YAML

```yaml

  ManifestWindowSeconds: Number
  MinBufferTimeSeconds: Number
  MinUpdatePeriodSeconds: Number
  SuggestedPresentationDelaySeconds: Number

```

## Properties

`ManifestWindowSeconds`

The total duration (in seconds) of each manifest. Minimum value: `30` seconds. Maximum value: `3600` seconds.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinBufferTimeSeconds`

Minimum amount of content (measured in seconds) that a player must keep available in the buffer. Minimum value: `2` seconds. Maximum value: `60` seconds.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinUpdatePeriodSeconds`

Minimum amount of time (in seconds) that the player should wait before requesting updates to the manifest. Minimum value: `2` seconds. Maximum value: `60` seconds.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SuggestedPresentationDelaySeconds`

Amount of time (in seconds) that the player should be from the live point at the end of the manifest. Minimum value: `2` seconds. Maximum value: `60` seconds.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::MediaTailor::Channel

HlsPlaylistSettings

All content copied from https://docs.aws.amazon.com/.
