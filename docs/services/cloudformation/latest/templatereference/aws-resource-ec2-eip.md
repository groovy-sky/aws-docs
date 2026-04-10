This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::EIP

Specifies an Elastic IP (EIP) address and can, optionally, associate it with an Amazon
EC2 instance.

You can allocate an Elastic IP address from an address pool owned by AWS or from an address pool created from a public IPv4 address range that you have brought
to AWS for use with your AWS resources using bring your
own IP addresses (BYOIP). For more information, see [Bring Your Own IP Addresses (BYOIP)](../../../ec2/latest/userguide/ec2-byoip.md)
in the _Amazon EC2 User Guide_.

For more information, see [Elastic IP Addresses](../../../ec2/latest/userguide/elastic-ip-addresses-eip.md)
in the _Amazon EC2 User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::EIP",
  "Properties" : {
      "Address" : String,
      "Domain" : String,
      "InstanceId" : String,
      "IpamPoolId" : String,
      "NetworkBorderGroup" : String,
      "PublicIpv4Pool" : String,
      "Tags" : [ Tag, ... ],
      "TransferAddress" : String
    }
}

```

### YAML

```yaml

Type: AWS::EC2::EIP
Properties:
  Address: String
  Domain: String
  InstanceId: String
  IpamPoolId: String
  NetworkBorderGroup: String
  PublicIpv4Pool: String
  Tags:
    - Tag
  TransferAddress: String

```

## Properties

`Address`

An Elastic IP address or a carrier IP address in a Wavelength Zone.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Domain`

The network ( `vpc`).

If you define an Elastic IP address and associate it with a VPC that is defined in the
same template, you must declare a dependency on the VPC-gateway attachment by using the
[DependsOn\
Attribute](../userguide/aws-attribute-dependson.md) on this resource.

_Required_: No

_Type_: String

_Allowed values_: `vpc | standard`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceId`

The ID of the instance.

###### Important

Updates to the `InstanceId` property may require _some_
_interruptions_. Updates on an EIP reassociates the address on its
associated resource.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IpamPoolId`

The ID of an IPAM pool which has an Amazon-provided or BYOIP public IPv4 CIDR provisioned to it. For more information, see [Allocate sequential Elastic IP addresses from an IPAM pool](../../../vpc/latest/ipam/tutorials-eip-pool.md) in the _Amazon VPC IPAM User Guide_.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NetworkBorderGroup`

A unique set of Availability Zones, Local Zones, or Wavelength Zones from which AWS
advertises IP addresses. Use this parameter to limit the IP address to this location. IP
addresses cannot move between network border groups.

Use [DescribeAvailabilityZones](../../../../reference/awsec2/latest/apireference/api-describeavailabilityzones.md) to view the network border groups.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PublicIpv4Pool`

The ID of an address pool that you own. Use this parameter to let Amazon EC2 select an
address from the address pool.

###### Important

Updates to the `PublicIpv4Pool` property may require _some_
_interruptions_. Updates on an EIP reassociates the address on its
associated resource.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Any tags assigned to the Elastic IP address.

###### Important

Updates to the `Tags` property may require _some_
_interruptions_. Updates on an EIP reassociates the address on its
associated resource.

_Required_: No

_Type_: Array of [Tag](aws-properties-ec2-eip-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TransferAddress`

The Elastic IP address you are accepting for transfer. You can only accept one transferred address. For more information on Elastic IP address transfers, see [Transfer Elastic IP addresses](../../../vpc/latest/userguide/vpc-eips.md#transfer-EIPs-intro) in the _Amazon Virtual Private Cloud User Guide_.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Elastic IP address.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AllocationId`

The ID that AWS assigns to represent the allocation of the address for
use with Amazon VPC. This is returned only for VPC elastic IP addresses. For example,
`eipalloc-5723d13e`.

`PublicIp`

The Elastic IP address.

## Examples

### Allocate an Elastic IP address

This example shows how to allocate an Elastic IP address and assign it
to an Amazon EC2 instance with the logical name `myInstance`.

#### JSON

```json

"Resources": {
  "myEIP" : {
      "Type" : "AWS::EC2::EIP",
      "Properties" : {
          "InstanceId" : { "Ref" : "myInstance" }
      }
  }
}
```

#### YAML

```yaml

Resources:
  myEIP:
    Type: AWS::EC2::EIP
    Properties:
      InstanceId: !Ref myInstance
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Tag

All content copied from https://docs.aws.amazon.com/.
