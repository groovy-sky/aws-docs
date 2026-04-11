This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel SrtOutputSettings

The `SrtOutputSettings` property type specifies Property description not available. for an [AWS::MediaLive::Channel](aws-resource-medialive-channel.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BufferMsec" : Integer,
  "ContainerSettings" : UdpContainerSettings,
  "Destination" : OutputLocationRef,
  "EncryptionType" : String,
  "Latency" : Integer
}

```

### YAML

```yaml

  BufferMsec: Integer
  ContainerSettings:
    UdpContainerSettings
  Destination:
    OutputLocationRef
  EncryptionType: String
  Latency: Integer

```

## Properties

`BufferMsec`

Property description not available.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ContainerSettings`

Property description not available.

_Required_: No

_Type_: [UdpContainerSettings](aws-properties-medialive-channel-udpcontainersettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Destination`

Property description not available.

_Required_: No

_Type_: [OutputLocationRef](aws-properties-medialive-channel-outputlocationref.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EncryptionType`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Latency`

Property description not available.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SrtOutputDestinationSettings

StandardHlsSettings

All content copied from https://docs.aws.amazon.com/.
