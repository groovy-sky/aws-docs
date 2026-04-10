This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::ChannelPlacementGroup

The `AWS::MediaLive::ChannelPlacementGroup` resource Property description not available. for MediaLive.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::MediaLive::ChannelPlacementGroup",
  "Properties" : {
      "ClusterId" : String,
      "Name" : String,
      "Nodes" : [ String, ... ],
      "Tags" : [ Tags, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::MediaLive::ChannelPlacementGroup
Properties:
  ClusterId: String
  Name: String
  Nodes:
    - String
  Tags:
    - Tags

```

## Properties

`ClusterId`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Nodes`

Property description not available.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Property description not available.

_Required_: No

_Type_: [Array](aws-properties-medialive-channelplacementgroup-tags.md) of [Tags](aws-properties-medialive-channelplacementgroup-tags.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`Arn`

The ARN of the channel placement group.

`Channels`

List of channel IDs added to the channel placement group.

`Id`

The unique identifier of the channel placement group.

`State`

The current state of the channel placement group.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

WebvttDestinationSettings

Tags

All content copied from https://docs.aws.amazon.com/.
