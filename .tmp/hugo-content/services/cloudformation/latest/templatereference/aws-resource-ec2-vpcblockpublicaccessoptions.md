This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::VPCBlockPublicAccessOptions

VPC Block Public Access (BPA) enables you to block resources in VPCs and subnets that you own in a Region from reaching or being reached from the internet through internet gateways and egress-only internet gateways. To learn more about VPC BPA, see [Block public access to VPCs and subnets](../../../vpc/latest/userguide/security-vpc-bpa.md) in the _Amazon VPC User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::VPCBlockPublicAccessOptions",
  "Properties" : {
      "InternetGatewayBlockMode" : String
    }
}

```

### YAML

```yaml

Type: AWS::EC2::VPCBlockPublicAccessOptions
Properties:
  InternetGatewayBlockMode: String

```

## Properties

`InternetGatewayBlockMode`

The desired VPC Block Public Access mode for internet gateways in your account. We do not allow you to create this resource type in an "off" mode since off is the default value.

- `block-bidirectional`: Block all traffic to and from internet gateways and egress-only internet gateways in this Region (except for excluded VPCs and subnets).

- `block-ingress`: Block all internet traffic to the VPCs in this Region (except for VPCs or subnets which are excluded). Only traffic to and from NAT gateways and egress-only internet gateways is allowed because these gateways only allow outbound connections to be established.

_Required_: Yes

_Type_: String

_Allowed values_: `block-bidirectional | block-ingress`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns your account ID.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AccountId`

The ID of the AWS account.

`ExclusionsAllowed`

Property description not available.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWS::EC2::VPCCidrBlock

All content copied from https://docs.aws.amazon.com/.
