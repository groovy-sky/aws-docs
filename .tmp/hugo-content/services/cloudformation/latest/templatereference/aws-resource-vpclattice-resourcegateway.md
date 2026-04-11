This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::VpcLattice::ResourceGateway

A resource gateway is a point of ingress into the VPC where a resource resides. It spans
multiple Availability Zones. For your resource to be accessible from all Availability Zones, you
should create your resource gateways to span as many Availability Zones as possible. A VPC can
have multiple resource gateways.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::VpcLattice::ResourceGateway",
  "Properties" : {
      "IpAddressType" : String,
      "Ipv4AddressesPerEni" : Integer,
      "Name" : String,
      "SecurityGroupIds" : [ String, ... ],
      "SubnetIds" : [ String, ... ],
      "Tags" : [ Tag, ... ],
      "VpcIdentifier" : String
    }
}

```

### YAML

```yaml

Type: AWS::VpcLattice::ResourceGateway
Properties:
  IpAddressType: String
  Ipv4AddressesPerEni: Integer
  Name: String
  SecurityGroupIds:
    - String
  SubnetIds:
    - String
  Tags:
    - Tag
  VpcIdentifier: String

```

## Properties

`IpAddressType`

The type of IP address used by the resource gateway.

_Required_: No

_Type_: String

_Allowed values_: `IPV4 | IPV6 | DUALSTACK`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Ipv4AddressesPerEni`

The number of IPv4 addresses in each ENI for the resource gateway.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the resource gateway.

_Required_: Yes

_Type_: String

_Pattern_: `^(?!rgw-)(?![-])(?!.*[-]$)(?!.*[-]{2})[a-z0-9-]+$`

_Minimum_: `3`

_Maximum_: `40`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SecurityGroupIds`

The IDs of the security groups applied to the resource gateway.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubnetIds`

The IDs of the VPC subnets for the resource gateway.

_Required_: Yes

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags for the resource gateway.

_Required_: No

_Type_: Array of [Tag](aws-properties-vpclattice-resourcegateway-tag.md)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcIdentifier`

The ID of the VPC for the resource gateway.

_Required_: Yes

_Type_: String

_Minimum_: `5`

_Maximum_: `50`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the resource gateway.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the resource gateway.

`Id`

The ID of the resource gateway.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Tag

All content copied from https://docs.aws.amazon.com/.
