This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaPackageV2::ChannelPolicy

Specifies the configuration parameters of a MediaPackage V2 channel policy.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::MediaPackageV2::ChannelPolicy",
  "Properties" : {
      "ChannelGroupName" : String,
      "ChannelName" : String,
      "Policy" : Json
    }
}

```

### YAML

```yaml

Type: AWS::MediaPackageV2::ChannelPolicy
Properties:
  ChannelGroupName: String
  ChannelName: String
  Policy: Json

```

## Properties

`ChannelGroupName`

The name of the channel group associated with the channel policy.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ChannelName`

The name of the channel associated with the channel policy.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Policy`

The policy associated with the channel.

_Required_: Yes

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns `ChannelGroupName/ChannelName`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWS::MediaPackageV2::OriginEndpoint

All content copied from https://docs.aws.amazon.com/.
