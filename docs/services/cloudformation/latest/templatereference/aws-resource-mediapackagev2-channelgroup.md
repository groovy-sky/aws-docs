This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaPackageV2::ChannelGroup

Specifies the configuration for a MediaPackage V2 channel group.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::MediaPackageV2::ChannelGroup",
  "Properties" : {
      "ChannelGroupName" : String,
      "Description" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::MediaPackageV2::ChannelGroup
Properties:
  ChannelGroupName: String
  Description: String
  Tags:
    - Tag

```

## Properties

`ChannelGroupName`

The name of the channel group.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

The configuration for a MediaPackage V2 channel group.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags associated with the channel group.

_Required_: No

_Type_: Array of [Tag](aws-properties-mediapackagev2-channelgroup-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns `arn:aws:mediapackagev2:region:AccountId:ChannelGroup/ChannelGroupName`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The attributes of the channel group. You can only use the `GetAtt` function for `readOnly` properties. For example, you can use the `GetAtt` function to retrieve the read-only `CreatedAt` property.

`Arn`

The Amazon Resource Name (ARN) of the channel group.

`CreatedAt`

The timestamp of the creation of the channel group.

`EgressDomain`

The egress domain of the channel group.

`ModifiedAt`

The timestamp of the modification of the channel group.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Tag

All content copied from https://docs.aws.amazon.com/.
