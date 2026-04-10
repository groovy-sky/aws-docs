This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::VPCEncryptionControl

Describes the configuration and state of VPC encryption controls.

For more information, see [Enforce VPC encryption in transit](../../../vpc/latest/userguide/vpc-encryption-controls.md) in the _Amazon VPC User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::VPCEncryptionControl",
  "Properties" : {
      "EgressOnlyInternetGatewayExclusionInput" : String,
      "ElasticFileSystemExclusionInput" : String,
      "InternetGatewayExclusionInput" : String,
      "LambdaExclusionInput" : String,
      "Mode" : String,
      "NatGatewayExclusionInput" : String,
      "Tags" : [ Tag, ... ],
      "VirtualPrivateGatewayExclusionInput" : String,
      "VpcId" : String,
      "VpcLatticeExclusionInput" : String,
      "VpcPeeringExclusionInput" : String
    }
}

```

### YAML

```yaml

Type: AWS::EC2::VPCEncryptionControl
Properties:
  EgressOnlyInternetGatewayExclusionInput: String
  ElasticFileSystemExclusionInput: String
  InternetGatewayExclusionInput: String
  LambdaExclusionInput: String
  Mode: String
  NatGatewayExclusionInput: String
  Tags:
    - Tag
  VirtualPrivateGatewayExclusionInput: String
  VpcId: String
  VpcLatticeExclusionInput: String
  VpcPeeringExclusionInput: String

```

## Properties

`EgressOnlyInternetGatewayExclusionInput`

Specifies whether to exclude egress-only internet gateway traffic from encryption enforcement.

_Required_: No

_Type_: String

_Allowed values_: `enable | disable`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ElasticFileSystemExclusionInput`

Specifies whether to exclude Elastic File System traffic from encryption enforcement.

_Required_: No

_Type_: String

_Allowed values_: `enable | disable`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InternetGatewayExclusionInput`

Specifies whether to exclude internet gateway traffic from encryption enforcement.

_Required_: No

_Type_: String

_Allowed values_: `enable | disable`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LambdaExclusionInput`

Specifies whether to exclude Lambda function traffic from encryption enforcement.

_Required_: No

_Type_: String

_Allowed values_: `enable | disable`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Mode`

The encryption mode for the VPC Encryption Control configuration.

_Required_: No

_Type_: String

_Allowed values_: `monitor | enforce`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NatGatewayExclusionInput`

Specifies whether to exclude NAT gateway traffic from encryption enforcement.

_Required_: No

_Type_: String

_Allowed values_: `enable | disable`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags assigned to the VPC Encryption Control configuration.

_Required_: No

_Type_: Array of [Tag](aws-properties-ec2-vpcencryptioncontrol-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VirtualPrivateGatewayExclusionInput`

Specifies whether to exclude virtual private gateway traffic from encryption enforcement.

_Required_: No

_Type_: String

_Allowed values_: `enable | disable`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcId`

The ID of the VPC for which to create the encryption control configuration.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VpcLatticeExclusionInput`

Specifies whether to exclude VPC Lattice traffic from encryption enforcement.

_Required_: No

_Type_: String

_Allowed values_: `enable | disable`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcPeeringExclusionInput`

Specifies whether to exclude VPC peering connection traffic from encryption enforcement.

_Required_: No

_Type_: String

_Allowed values_: `enable | disable`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the VPC Encryption Control ID.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

Describes the configuration and state of VPC encryption controls.

For more information, see [Enforce VPC encryption in transit](../../../vpc/latest/userguide/vpc-encryption-controls.md) in the _Amazon VPC User Guide_.

`ResourceExclusions.EgressOnlyInternetGateway.State`

The current state of the exclusion configuration.

`ResourceExclusions.EgressOnlyInternetGateway.StateMessage`

A message providing additional information about the exclusion state.

`ResourceExclusions.ElasticFileSystem.State`

The current state of the exclusion configuration.

`ResourceExclusions.ElasticFileSystem.StateMessage`

A message providing additional information about the exclusion state.

`ResourceExclusions.InternetGateway.State`

The current state of the exclusion configuration.

`ResourceExclusions.InternetGateway.StateMessage`

A message providing additional information about the exclusion state.

`ResourceExclusions.Lambda.State`

The current state of the exclusion configuration.

`ResourceExclusions.Lambda.StateMessage`

A message providing additional information about the exclusion state.

`ResourceExclusions.NatGateway.State`

The current state of the exclusion configuration.

`ResourceExclusions.NatGateway.StateMessage`

A message providing additional information about the exclusion state.

`ResourceExclusions.VirtualPrivateGateway.State`

The current state of the exclusion configuration.

`ResourceExclusions.VirtualPrivateGateway.StateMessage`

A message providing additional information about the exclusion state.

`ResourceExclusions.VpcLattice.State`

The current state of the exclusion configuration.

`ResourceExclusions.VpcLattice.StateMessage`

A message providing additional information about the exclusion state.

`ResourceExclusions.VpcPeering.State`

The current state of the exclusion configuration.

`ResourceExclusions.VpcPeering.StateMessage`

A message providing additional information about the exclusion state.

`State`

The current state of the VPC Encryption Control configuration.

`StateMessage`

A message providing additional information about the encryption control state.

`VpcEncryptionControlId`

The ID of the VPC Encryption Control configuration.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::EC2::VPCDHCPOptionsAssociation

ResourceExclusions

All content copied from https://docs.aws.amazon.com/.
