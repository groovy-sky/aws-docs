This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::IPAMPrefixListResolverTarget

An IPAM prefix list resolver target is an association between a specific customer-managed prefix list and an IPAM prefix list resolver. The target enables the resolver to synchronize CIDRs selected by its rules into the specified prefix list, which can then be referenced in AWS resources.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::IPAMPrefixListResolverTarget",
  "Properties" : {
      "DesiredVersion" : Integer,
      "IpamPrefixListResolverId" : String,
      "PrefixListId" : String,
      "PrefixListRegion" : String,
      "Tags" : [ Tag, ... ],
      "TrackLatestVersion" : Boolean
    }
}

```

### YAML

```yaml

Type: AWS::EC2::IPAMPrefixListResolverTarget
Properties:
  DesiredVersion: Integer
  IpamPrefixListResolverId: String
  PrefixListId: String
  PrefixListRegion: String
  Tags:
    - Tag
  TrackLatestVersion: Boolean

```

## Properties

`DesiredVersion`

The desired version of the prefix list that this target should synchronize with.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IpamPrefixListResolverId`

The ID of the IPAM prefix list resolver associated with this target.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PrefixListId`

The ID of the managed prefix list associated with this target.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PrefixListRegion`

The AWS Region where the prefix list associated with this target is located.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags assigned to the IPAM prefix list resolver target.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ec2-ipamprefixlistresolvertarget-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TrackLatestVersion`

Indicates whether this target automatically tracks the latest version of the prefix list.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the prefix list resolver target ID. For example: `ipam-plrt-111122223333`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`IpamPrefixListResolverTargetArn`

The Amazon Resource Name (ARN) of the IPAM prefix list resolver target.

`IpamPrefixListResolverTargetId`

The ID of the IPAM prefix list resolver target.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Tag
