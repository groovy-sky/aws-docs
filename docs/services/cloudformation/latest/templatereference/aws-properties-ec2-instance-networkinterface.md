This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::Instance NetworkInterface

Specifies a network interface that is to be attached to an instance.

You can create a network interface when launching an instance. For an example, see the
[AWS::EC2::Instance examples](../userguide/aws-properties-ec2-instance.md#aws-properties-ec2-instance--examples--Automatically_assign_a_public_IP_address).

Alternatively, you can attach an existing network interface when launching an instance.
For an example, see the [AWS::EC2:NetworkInterface examples](../userguide/aws-resource-ec2-networkinterface.md#aws-resource-ec2-networkinterface--examples).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AssociateCarrierIpAddress" : Boolean,
  "AssociatePublicIpAddress" : Boolean,
  "DeleteOnTermination" : Boolean,
  "Description" : String,
  "DeviceIndex" : String,
  "EnaSrdSpecification" : EnaSrdSpecification,
  "GroupSet" : [ String, ... ],
  "Ipv6AddressCount" : Integer,
  "Ipv6Addresses" : [ InstanceIpv6Address, ... ],
  "NetworkInterfaceId" : String,
  "PrivateIpAddress" : String,
  "PrivateIpAddresses" : [ PrivateIpAddressSpecification, ... ],
  "SecondaryPrivateIpAddressCount" : Integer,
  "SubnetId" : String
}

```

### YAML

```yaml

  AssociateCarrierIpAddress: Boolean
  AssociatePublicIpAddress: Boolean
  DeleteOnTermination: Boolean
  Description: String
  DeviceIndex: String
  EnaSrdSpecification:
    EnaSrdSpecification
  GroupSet:
    - String
  Ipv6AddressCount: Integer
  Ipv6Addresses:
    - InstanceIpv6Address
  NetworkInterfaceId: String
  PrivateIpAddress: String
  PrivateIpAddresses:
    - PrivateIpAddressSpecification
  SecondaryPrivateIpAddressCount: Integer
  SubnetId: String

```

## Properties

`AssociateCarrierIpAddress`

Indicates whether to assign a carrier IP address to the network interface.

You can only assign a carrier IP address to a network interface that is in a subnet in
a Wavelength Zone. For more information about carrier IP addresses, see [Carrier IP address](../../../wavelength/latest/developerguide/how-wavelengths-work.md#provider-owned-ip) in the _AWS Wavelength Developer_
_Guide_.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`AssociatePublicIpAddress`

Indicates whether to assign a public IPv4 address to an instance. Applies only if
creating a network interface when launching an instance. The network interface must be the
primary network interface. If launching into a default subnet, the default value is
`true`.

AWS charges for all public IPv4 addresses, including public IPv4 addresses
associated with running instances and Elastic IP addresses. For more information, see the _Public IPv4 Address_ tab
on the [VPC pricing page](https://aws.amazon.com/vpc/pricing).

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DeleteOnTermination`

Indicates whether the network interface is deleted when the instance is terminated.
Applies only if creating a network interface when launching an instance.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

The description of the network interface. Applies only if creating a network interface
when launching an instance.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DeviceIndex`

The position of the network interface in the attachment order. A primary network
interface has a device index of 0.

If you create a network interface when launching an instance, you must specify the
device index.

_Required_: Conditional

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EnaSrdSpecification`

Configures ENA Express for UDP network traffic.

_Required_: No

_Type_: [EnaSrdSpecification](aws-properties-ec2-instance-enasrdspecification.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`GroupSet`

The IDs of the security groups for the network interface. Applies only if creating a
network interface when launching an instance.

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Ipv6AddressCount`

A number of IPv6 addresses to assign to the network interface. Amazon EC2 chooses the
IPv6 addresses from the range of the subnet. You cannot specify this option and the
option to assign specific IPv6 addresses in the same request. You can specify this
option if you've specified a minimum number of instances to launch.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Ipv6Addresses`

The IPv6 addresses to assign to the network interface. You cannot specify this option
and the option to assign a number of IPv6 addresses in the same request. You cannot
specify this option if you've specified a minimum number of instances to launch.

_Required_: No

_Type_: Array of [InstanceIpv6Address](aws-properties-ec2-instance-instanceipv6address.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NetworkInterfaceId`

The ID of the network interface, when attaching an existing network interface.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PrivateIpAddress`

The private IPv4 address of the network interface. Applies only if creating a network
interface when launching an instance.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PrivateIpAddresses`

One or more private IPv4 addresses to assign to the network interface. Only one private
IPv4 address can be designated as primary.

_Required_: No

_Type_: Array of [PrivateIpAddressSpecification](aws-properties-ec2-instance-privateipaddressspecification.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SecondaryPrivateIpAddressCount`

The number of secondary private IPv4 addresses. You can't specify this option and
specify more than one private IP address using the private IP addresses option.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SubnetId`

The ID of the subnet associated with the network interface.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MetadataOptions

PrivateDnsNameOptions

All content copied from https://docs.aws.amazon.com/.
