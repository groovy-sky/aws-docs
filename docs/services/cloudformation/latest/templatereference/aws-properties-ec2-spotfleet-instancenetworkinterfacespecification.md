This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::SpotFleet InstanceNetworkInterfaceSpecification

Describes a network interface.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AssociatePublicIpAddress" : Boolean,
  "DeleteOnTermination" : Boolean,
  "Description" : String,
  "DeviceIndex" : Integer,
  "Groups" : [ String, ... ],
  "Ipv6AddressCount" : Integer,
  "Ipv6Addresses" : [ InstanceIpv6Address, ... ],
  "NetworkInterfaceId" : String,
  "PrivateIpAddresses" : [ PrivateIpAddressSpecification, ... ],
  "SecondaryPrivateIpAddressCount" : Integer,
  "SubnetId" : String
}

```

### YAML

```yaml

  AssociatePublicIpAddress: Boolean
  DeleteOnTermination: Boolean
  Description: String
  DeviceIndex: Integer
  Groups:
    - String
  Ipv6AddressCount: Integer
  Ipv6Addresses:
    - InstanceIpv6Address
  NetworkInterfaceId: String
  PrivateIpAddresses:
    - PrivateIpAddressSpecification
  SecondaryPrivateIpAddressCount: Integer
  SubnetId: String

```

## Properties

`AssociatePublicIpAddress`

Indicates whether to assign a public IPv4 address to an instance you launch in a VPC.
The public IP address can only be assigned to a network interface for eth0, and can only
be assigned to a new network interface, not an existing one. You cannot specify more
than one network interface in the request. If launching into a default subnet, the
default value is `true`.

AWS charges for all public IPv4 addresses, including public IPv4 addresses
associated with running instances and Elastic IP addresses. For more information, see the _Public IPv4 Address_ tab on the [Amazon VPC pricing page](https://aws.amazon.com/vpc/pricing).

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DeleteOnTermination`

Indicates whether the network interface is deleted when the instance is
terminated.

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

If you specify a network interface when launching an instance, you must specify the
device index.

_Required_: Conditional

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Groups`

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

_Type_: Array of [InstanceIpv6Address](aws-properties-ec2-spotfleet-instanceipv6address.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NetworkInterfaceId`

The ID of the network interface.

If you are creating a Spot Fleet, omit this parameter because you can’t specify a
network interface ID in a launch specification.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PrivateIpAddresses`

The private IPv4 addresses to assign to the network interface. Only one private IPv4
address can be designated as primary. You cannot specify this option if you're launching
more than one instance in a [RunInstances](../../../../reference/awsec2/latest/apireference/api-runinstances.md)
request.

_Required_: No

_Type_: Array of [PrivateIpAddressSpecification](aws-properties-ec2-spotfleet-privateipaddressspecification.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SecondaryPrivateIpAddressCount`

The number of secondary private IPv4 addresses. You can’t specify this parameter and
also specify a secondary private IP address using the `PrivateIpAddress`
parameter.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SubnetId`

The ID of the subnet associated with the network interface.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InstanceIpv6Address

InstanceRequirementsRequest

All content copied from https://docs.aws.amazon.com/.
