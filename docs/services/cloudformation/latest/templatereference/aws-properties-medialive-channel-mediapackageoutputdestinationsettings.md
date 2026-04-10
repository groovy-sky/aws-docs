This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel MediaPackageOutputDestinationSettings

Destination settings for a MediaPackage output.

The parent of this entity is OutputDestination.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ChannelEndpointId" : String,
  "ChannelGroup" : String,
  "ChannelId" : String,
  "ChannelName" : String,
  "MediaPackageRegionName" : String
}

```

### YAML

```yaml

  ChannelEndpointId: String
  ChannelGroup: String
  ChannelId: String
  ChannelName: String
  MediaPackageRegionName: String

```

## Properties

`ChannelEndpointId`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ChannelGroup`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ChannelId`

The ID of the channel in MediaPackage that is the destination for this output group. You don't need to specify
the individual inputs in MediaPackage; MediaLive handles the connection of the two MediaLive pipelines to the two
MediaPackage inputs. The MediaPackage channel and MediaLive channel must be in the same Region.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ChannelName`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MediaPackageRegionName`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MediaPackageGroupSettings

MediaPackageOutputSettings

All content copied from https://docs.aws.amazon.com/.
