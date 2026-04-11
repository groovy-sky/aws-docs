This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Input Smpte2110ReceiverGroupSdpSettings

The `Smpte2110ReceiverGroupSdpSettings` property type specifies Property description not available. for an [AWS::MediaLive::Input](aws-resource-medialive-input.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AncillarySdps" : [ InputSdpLocation, ... ],
  "AudioSdps" : [ InputSdpLocation, ... ],
  "VideoSdp" : InputSdpLocation
}

```

### YAML

```yaml

  AncillarySdps:
    - InputSdpLocation
  AudioSdps:
    - InputSdpLocation
  VideoSdp:
    InputSdpLocation

```

## Properties

`AncillarySdps`

Property description not available.

_Required_: No

_Type_: Array of [InputSdpLocation](aws-properties-medialive-input-inputsdplocation.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AudioSdps`

Property description not available.

_Required_: No

_Type_: Array of [InputSdpLocation](aws-properties-medialive-input-inputsdplocation.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VideoSdp`

Property description not available.

_Required_: No

_Type_: [InputSdpLocation](aws-properties-medialive-input-inputsdplocation.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Smpte2110ReceiverGroup

Smpte2110ReceiverGroupSettings

All content copied from https://docs.aws.amazon.com/.
