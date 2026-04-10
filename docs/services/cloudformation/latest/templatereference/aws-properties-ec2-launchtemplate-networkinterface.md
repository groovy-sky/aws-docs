This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::LaunchTemplate NetworkInterface

Specifies the parameters for a network interface.

`NetworkInterface` is a property of [AWS::EC2::LaunchTemplate LaunchTemplateData](../userguide/aws-properties-ec2-launchtemplate-launchtemplatedata.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AssociateCarrierIpAddress" : Boolean,
  "AssociatePublicIpAddress" : Boolean,
  "ConnectionTrackingSpecification" : ConnectionTrackingSpecification,
  "DeleteOnTermination" : Boolean,
  "Description" : String,
  "DeviceIndex" : Integer,
  "EnaQueueCount" : Integer,
  "EnaSrdSpecification" : EnaSrdSpecification,
  "Groups" : [ String, ... ],
  "InterfaceType" : String,
  "Ipv4PrefixCount" : Integer,
  "Ipv4Prefixes" : [ Ipv4PrefixSpecification, ... ],
  "Ipv6AddressCount" : Integer,
  "Ipv6Addresses" : [ Ipv6Add, ... ],
  "Ipv6PrefixCount" : Integer,
  "Ipv6Prefixes" : [ Ipv6PrefixSpecification, ... ],
  "NetworkCardIndex" : Integer,
  "NetworkInterfaceId" : String,
  "PrimaryIpv6" : Boolean,
  "PrivateIpAddress" : String,
  "PrivateIpAddresses" : [ PrivateIpAdd, ... ],
  "SecondaryPrivateIpAddressCount" : Integer,
  "SubnetId" : String
}

```

### YAML

```yaml

  AssociateCarrierIpAddress: Boolean
  AssociatePublicIpAddress: Boolean
  ConnectionTrackingSpecification:
    ConnectionTrackingSpecification
  DeleteOnTermination: Boolean
  Description: String
  DeviceIndex: Integer
  EnaQueueCount: Integer
  EnaSrdSpecification:
    EnaSrdSpecification
  Groups:
    - String
  InterfaceType: String
  Ipv4PrefixCount: Integer
  Ipv4Prefixes:
    - Ipv4PrefixSpecification
  Ipv6AddressCount: Integer
  Ipv6Addresses:
    - Ipv6Add
  Ipv6PrefixCount: Integer
  Ipv6Prefixes:
    - Ipv6PrefixSpecification
  NetworkCardIndex: Integer
  NetworkInterfaceId: String
  PrimaryIpv6: Boolean
  PrivateIpAddress: String
  PrivateIpAddresses:
    - PrivateIpAdd
  SecondaryPrivateIpAddressCount: Integer
  SubnetId: String

```

## Properties

`AssociateCarrierIpAddress`

Associates a Carrier IP address with eth0 for a new network interface.

Use this option when you launch an instance in a Wavelength Zone and want to associate
a Carrier IP address with the network interface. For more information about Carrier IP
addresses, see [Carrier IP addresses](../../../wavelength/latest/developerguide/how-wavelengths-work.md#provider-owned-ip) in the _AWS Wavelength Developer_
_Guide_.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AssociatePublicIpAddress`

Associates a public IPv4 address with eth0 for a new network interface.

AWS charges for all public IPv4 addresses, including public IPv4 addresses
associated with running instances and Elastic IP addresses. For more information, see the _Public IPv4 Address_ tab on the [Amazon VPC pricing page](https://aws.amazon.com/vpc/pricing).

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConnectionTrackingSpecification`

A connection tracking specification for the network interface.

_Required_: No

_Type_: [ConnectionTrackingSpecification](aws-properties-ec2-launchtemplate-connectiontrackingspecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeleteOnTermination`

Indicates whether the network interface is deleted when the instance is
terminated.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description for the network interface.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeviceIndex`

The device index for the network interface attachment. The primary network interface
has a device index of 0. If the network interface is of type `interface`, you
must specify a device index.

If you create a launch template that includes secondary network interfaces but no
primary network interface, and you specify it using the `LaunchTemplate` property
of `AWS::EC2::Instance`, then you must include a primary network interface using
the `NetworkInterfaces` property of `AWS::EC2::Instance`.

_Required_: Conditional

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnaQueueCount`

The number of ENA queues to be created with the instance.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnaSrdSpecification`

The ENA Express configuration for the network interface.

_Required_: No

_Type_: [EnaSrdSpecification](aws-properties-ec2-launchtemplate-enasrdspecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Groups`

The IDs of one or more security groups.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InterfaceType`

The type of network interface. To create an Elastic Fabric Adapter (EFA), specify
`efa` or `efa`. For more information, see [Elastic Fabric Adapter for AI/ML\
and HPC workloads on Amazon EC2](../../../ec2/latest/userguide/efa.md) in the
_Amazon EC2 User Guide_.

If you are not creating an EFA, specify `interface` or omit this
parameter.

If you specify `efa-only`, do not assign any IP addresses to the network
interface. EFA-only network interfaces do not support IP addresses.

Valid values: `interface` \| `efa` \| `efa-only`

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Ipv4PrefixCount`

The number of IPv4 prefixes to be automatically assigned to the network interface. You
cannot use this option if you use the `Ipv4Prefix` option.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Ipv4Prefixes`

One or more IPv4 prefixes to be assigned to the network interface. You cannot use this
option if you use the `Ipv4PrefixCount` option.

_Required_: No

_Type_: Array of [Ipv4PrefixSpecification](aws-properties-ec2-launchtemplate-ipv4prefixspecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Ipv6AddressCount`

The number of IPv6 addresses to assign to a network interface. Amazon EC2
automatically selects the IPv6 addresses from the subnet range. You can't use this
option if specifying specific IPv6 addresses.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Ipv6Addresses`

One or more specific IPv6 addresses from the IPv6 CIDR block range of your subnet. You
can't use this option if you're specifying a number of IPv6 addresses.

_Required_: No

_Type_: Array of [Ipv6Add](aws-properties-ec2-launchtemplate-ipv6add.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Ipv6PrefixCount`

The number of IPv6 prefixes to be automatically assigned to the network interface. You
cannot use this option if you use the `Ipv6Prefix` option.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Ipv6Prefixes`

One or more IPv6 prefixes to be assigned to the network interface. You cannot use this
option if you use the `Ipv6PrefixCount` option.

_Required_: No

_Type_: Array of [Ipv6PrefixSpecification](aws-properties-ec2-launchtemplate-ipv6prefixspecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NetworkCardIndex`

The index of the network card. Some instance types support multiple network cards. The
primary network interface must be assigned to network card index 0. The default is
network card index 0.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NetworkInterfaceId`

The ID of the network interface.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PrimaryIpv6`

The primary IPv6 address of the network interface. When you enable an IPv6 GUA address
to be a primary IPv6, the first IPv6 GUA will be made the primary IPv6 address until the
instance is terminated or the network interface is detached. For more information about
primary IPv6 addresses, see [RunInstances](../../../../reference/awsec2/latest/apireference/api-runinstances.md).

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PrivateIpAddress`

The primary private IPv4 address of the network interface.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PrivateIpAddresses`

One or more private IPv4 addresses.

_Required_: No

_Type_: Array of [PrivateIpAdd](aws-properties-ec2-launchtemplate-privateipadd.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecondaryPrivateIpAddressCount`

The number of secondary private IPv4 addresses to assign to a network
interface.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubnetId`

The ID of the subnet for the network interface.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [LaunchTemplateInstanceNetworkInterfaceSpecificationRequest](../../../../reference/awsec2/latest/apireference/api-launchtemplateinstancenetworkinterfacespecificationrequest.md) in the
_Amazon EC2 API Reference_

- [Create a launch template for an Auto Scaling group](../../../autoscaling/ec2/userguide/create-launch-template.md) in the _Amazon EC2 Auto Scaling User Guide_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NetworkBandwidthGbps

NetworkInterfaceCount

All content copied from https://docs.aws.amazon.com/.
