This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel OutputDestination

Configuration information for an output.

This entity is at the top level in the channel.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Id" : String,
  "LogicalInterfaceNames" : [ String, ... ],
  "MediaPackageSettings" : [ MediaPackageOutputDestinationSettings, ... ],
  "MultiplexSettings" : MultiplexProgramChannelDestinationSettings,
  "Settings" : [ OutputDestinationSettings, ... ],
  "SrtSettings" : [ SrtOutputDestinationSettings, ... ]
}

```

### YAML

```yaml

  Id: String
  LogicalInterfaceNames:
    - String
  MediaPackageSettings:
    - MediaPackageOutputDestinationSettings
  MultiplexSettings:
    MultiplexProgramChannelDestinationSettings
  Settings:
    - OutputDestinationSettings
  SrtSettings:
    - SrtOutputDestinationSettings

```

## Properties

`Id`

The ID for this destination.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogicalInterfaceNames`

Property description not available.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MediaPackageSettings`

The destination settings for a MediaPackage output.

_Required_: No

_Type_: Array of [MediaPackageOutputDestinationSettings](aws-properties-medialive-channel-mediapackageoutputdestinationsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MultiplexSettings`

Destination settings for a Multiplex output; one destination for both encoders.

_Required_: No

_Type_: [MultiplexProgramChannelDestinationSettings](aws-properties-medialive-channel-multiplexprogramchanneldestinationsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Settings`

The destination settings for an output.

_Required_: No

_Type_: Array of [OutputDestinationSettings](aws-properties-medialive-channel-outputdestinationsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SrtSettings`

Property description not available.

_Required_: No

_Type_: Array of [SrtOutputDestinationSettings](aws-properties-medialive-channel-srtoutputdestinationsettings.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Output

OutputDestinationSettings

All content copied from https://docs.aws.amazon.com/.
