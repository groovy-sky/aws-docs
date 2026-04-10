This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::VPCBlockPublicAccessExclusion

Create a VPC Block Public Access (BPA) exclusion. A VPC BPA exclusion is a mode that can be applied to a single VPC or subnet that exempts it from the account’s BPA mode and will allow bidirectional or egress-only access. You can create BPA exclusions for VPCs and subnets even when BPA is not enabled on the account to ensure that there is no traffic disruption to the exclusions when VPC BPA is turned on. To learn more about VPC BPA, see [Block public access to VPCs and subnets](../../../vpc/latest/userguide/security-vpc-bpa.md) in the _Amazon VPC User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::VPCBlockPublicAccessExclusion",
  "Properties" : {
      "InternetGatewayExclusionMode" : String,
      "SubnetId" : String,
      "Tags" : [ Tag, ... ],
      "VpcId" : String
    }
}

```

### YAML

```yaml

Type: AWS::EC2::VPCBlockPublicAccessExclusion
Properties:
  InternetGatewayExclusionMode: String
  SubnetId: String
  Tags:
    - Tag
  VpcId: String

```

## Properties

`InternetGatewayExclusionMode`

The desired VPC Block Public Access mode for a specific VPC or subnet exclusion.

- `allow-bidirectional`: Allow all internet traffic to and from the excluded VPCs and subnets.

- `allow-egress`: Allow outbound internet traffic from the excluded VPCs and subnets. Block inbound internet traffic to the excluded VPCs and subnets. Only applies when VPC Block Public Access is set to `block-bidirectional`.

_Required_: Yes

_Type_: String

_Allowed values_: `allow-bidirectional | allow-egress`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubnetId`

The ID of the subnet you want to exclude. Required only if you don't specify VpcId.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

An array of key-value pairs to apply to this resource.

_Required_: No

_Type_: Array of [Tag](aws-properties-ec2-vpcblockpublicaccessexclusion-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcId`

The ID of the VPC you want to exclude. Required only if you don't specify SubnetId.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the exclusion.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ExclusionId`

The ID of the exclusion.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Tag

All content copied from https://docs.aws.amazon.com/.
