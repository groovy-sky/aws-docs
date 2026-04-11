This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel MultiplexOutputSettings

Configuration of a Multiplex output.

The parent of this entity is OutputSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ContainerSettings" : MultiplexContainerSettings,
  "Destination" : OutputLocationRef
}

```

### YAML

```yaml

  ContainerSettings:
    MultiplexContainerSettings
  Destination:
    OutputLocationRef

```

## Properties

`ContainerSettings`

Property description not available.

_Required_: No

_Type_: [MultiplexContainerSettings](aws-properties-medialive-channel-multiplexcontainersettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Destination`

Destination is a Multiplex.

_Required_: No

_Type_: [OutputLocationRef](aws-properties-medialive-channel-outputlocationref.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MultiplexM2tsSettings

MultiplexProgramChannelDestinationSettings

All content copied from https://docs.aws.amazon.com/.
