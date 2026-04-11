This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::NetworkInterface

Describes a network interface in an Amazon EC2 instance for AWS CloudFormation.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::NetworkInterface",
  "Properties" : {
      "ConnectionTrackingSpecification" : ConnectionTrackingSpecification,
      "Description" : String,
      "EnablePrimaryIpv6" : Boolean,
      "GroupSet" : [ String, ... ],
      "InterfaceType" : String,
      "Ipv4PrefixCount" : Integer,
      "Ipv4Prefixes" : [ Ipv4PrefixSpecification, ... ],
      "Ipv6AddressCount" : Integer,
      "Ipv6Addresses" : [ InstanceIpv6Address, ... ],
      "Ipv6PrefixCount" : Integer,
      "Ipv6Prefixes" : [ Ipv6PrefixSpecification, ... ],
      "PrivateIpAddress" : String,
      "PrivateIpAddresses" : [ PrivateIpAddressSpecification, ... ],
      "PublicIpDnsHostnameTypeSpecification" : String,
      "SecondaryPrivateIpAddressCount" : Integer,
      "SourceDestCheck" : Boolean,
      "SubnetId" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::EC2::NetworkInterface
Properties:
  ConnectionTrackingSpecification:
    ConnectionTrackingSpecification
  Description: String
  EnablePrimaryIpv6: Boolean
  GroupSet:
    - String
  InterfaceType: String
  Ipv4PrefixCount: Integer
  Ipv4Prefixes:
    - Ipv4PrefixSpecification
  Ipv6AddressCount: Integer
  Ipv6Addresses:
    - InstanceIpv6Address
  Ipv6PrefixCount: Integer
  Ipv6Prefixes:
    - Ipv6PrefixSpecification
  PrivateIpAddress: String
  PrivateIpAddresses:
    - PrivateIpAddressSpecification
  PublicIpDnsHostnameTypeSpecification: String
  SecondaryPrivateIpAddressCount: Integer
  SourceDestCheck: Boolean
  SubnetId: String
  Tags:
    - Tag

```

## Properties

`ConnectionTrackingSpecification`

A connection tracking specification for the network interface.

_Required_: No

_Type_: [ConnectionTrackingSpecification](aws-properties-ec2-networkinterface-connectiontrackingspecification.md)

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`Description`

A description for the network interface.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnablePrimaryIpv6`

If you’re modifying a network interface in a dual-stack or IPv6-only subnet, you have
the option to assign a primary IPv6 IP address. A primary IPv6 address is an IPv6 GUA
address associated with an ENI that you have enabled to use a primary IPv6 address. Use
this option if the instance that this ENI will be attached to relies on its IPv6 address
not changing. AWS will automatically assign an IPv6 address associated
with the ENI attached to your instance to be the primary IPv6 address. Once you enable
an IPv6 GUA address to be a primary IPv6, you cannot disable it. When you enable an IPv6
GUA address to be a primary IPv6, the first IPv6 GUA will be made the primary IPv6
address until the instance is terminated or the network interface is detached. If you
have multiple IPv6 addresses associated with an ENI attached to your instance and you
enable a primary IPv6 address, the first IPv6 GUA address associated with the ENI
becomes the primary IPv6 address.

_Required_: No

_Type_: Boolean

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`GroupSet`

The IDs of the security groups associated with this network interface.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InterfaceType`

The type of network interface. The default is `interface`. The supported values
are `efa` and `trunk`.

_Required_: No

_Type_: String

_Allowed values_: `efa | efa-only | branch | trunk`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Ipv4PrefixCount`

The number of IPv4 prefixes to be automatically assigned to the network interface.

When creating a network interface, you can't specify a count of IPv4 prefixes if you've specified one of the following: specific IPv4 prefixes, specific private IPv4 addresses, or a count of private IPv4 addresses.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Ipv4Prefixes`

The IPv4 delegated prefixes that are assigned to the network interface.

When creating a network interface, you can't specify IPv4 prefixes if you've specified one of the following: a count of IPv4 prefixes, specific private IPv4 addresses, or a count of private IPv4 addresses.

_Required_: No

_Type_: Array of [Ipv4PrefixSpecification](aws-properties-ec2-networkinterface-ipv4prefixspecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Ipv6AddressCount`

The number of IPv6 addresses to assign to the network interface. Amazon EC2 automatically
selects the IPv6 addresses from the subnet range. To specify specific IPv6 addresses, use
the `Ipv6Addresses` property and don't specify this property.

When creating a network interface, you can't specify a count of IPv6 addresses if you've specified one of the following: specific IPv6 addresses, specific IPv6 prefixes, or a count of IPv6 prefixes.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Ipv6Addresses`

The IPv6 addresses from the IPv6 CIDR block range of your subnet to assign to the network interface.
If you're specifying a number of IPv6 addresses, use the `Ipv6AddressCount` property and don't
specify this property.

When creating a network interface, you can't specify IPv6 addresses if you've specified one of the following:
a count of IPv6 addresses, specific IPv6 prefixes, or a count of IPv6 prefixes.

_Required_: No

_Type_: Array of [InstanceIpv6Address](aws-properties-ec2-networkinterface-instanceipv6address.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Ipv6PrefixCount`

The number of IPv6 prefixes to be automatically assigned to the network interface.

When creating a network interface, you can't specify a count of IPv6 prefixes if you've specified one of the following:
specific IPv6 prefixes, specific IPv6 addresses, or a count of IPv6 addresses.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Ipv6Prefixes`

The IPv6 delegated prefixes that are assigned to the network interface.

When creating a network interface, you can't specify IPv6 prefixes if you've specified one of the following:
a count of IPv6 prefixes, specific IPv6 addresses, or a count of IPv6 addresses.

_Required_: No

_Type_: Array of [Ipv6PrefixSpecification](aws-properties-ec2-networkinterface-ipv6prefixspecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PrivateIpAddress`

The private IPv4 address to assign to the network interface as the primary private IP address.
If you want to specify multiple private IP addresses, use the `PrivateIpAddresses` property.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PrivateIpAddresses`

The private IPv4 addresses to assign to the network interface. You can specify a primary private
IP address by setting the value of the `Primary` property to `true`
in the `PrivateIpAddressSpecification` property. If you want EC2 to
automatically assign private IP addresses, use the
`SecondaryPrivateIpAddressCount` property and do not specify this
property.

When creating a network interface, you can't specify private IPv4 addresses if you've specified one of the following:
a count of private IPv4 addresses, specific IPv4 prefixes, or a count of IPv4 prefixes.

_Required_: No

_Type_: Array of [PrivateIpAddressSpecification](aws-properties-ec2-networkinterface-privateipaddressspecification.md)

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`PublicIpDnsHostnameTypeSpecification`

Property description not available.

_Required_: No

_Type_: String

_Allowed values_: `public-dual-stack-dns-name | public-ipv4-dns-name | public-ipv6-dns-name`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecondaryPrivateIpAddressCount`

The number of secondary private IPv4 addresses to assign to a network interface. When
you specify a number of secondary IPv4 addresses, Amazon EC2 selects these IP addresses
within the subnet's IPv4 CIDR range. You can't specify this option and specify more than
one private IP address using `privateIpAddresses`.

When creating a Network Interface, you can't specify a count of private IPv4 addresses if you've specified one of the following: specific private
IPv4 addresses, specific IPv4 prefixes, or a count of IPv4 prefixes.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceDestCheck`

Enable or disable source/destination checks, which ensure that the instance is either
the source or the destination of any traffic that it receives. If the value is
`true`, source/destination checks are enabled; otherwise, they are
disabled. The default value is `true`. You must disable source/destination
checks if the instance runs services such as network address translation, routing, or
firewalls.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubnetId`

The ID of the subnet to associate with the network interface.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags to apply to the network interface.

_Required_: No

_Type_: Array of [Tag](aws-properties-ec2-networkinterface-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the network interface.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Id`

The ID of the network interface.

`PrimaryIpv6Address`

The primary IPv6 address of the network interface.

`PrimaryPrivateIpAddress`

The primary private IP address of the network interface. For example,
`10.0.0.192`.

`SecondaryPrivateIpAddresses`

The secondary private IP addresses of the network interface. For example,
`["10.0.0.161", "10.0.0.162", "10.0.0.163"]`.

`VpcId`

The ID of the VPC.

## Examples

- [Basic network interface](#aws-resource-ec2-networkinterface--examples--Basic_network_interface)

- [Attach a network interface to an EC2 instance at launch](#aws-resource-ec2-networkinterface--examples--Attach_a_network_interface_to_an_EC2_instance_at_launch)

### Basic network interface

This example creates a standalone elastic network interface (ENI). To learn how to
attach this network interface to an instance at launch, see the next example on this
page.

#### JSON

```json

"myENI" : {
   "Type" : "AWS::EC2::NetworkInterface",
   "Properties" : {
      "Tags": [{"Key":"stack","Value":"production"}],
      "Description": "A nice description.",
      "SourceDestCheck": "false",
      "GroupSet": ["sg-75zzz219"],
      "SubnetId": "subnet-3z648z53",
      "PrivateIpAddress": "10.0.0.16"
   }
}
```

#### YAML

```yaml

   myENI:
      Type: AWS::EC2::NetworkInterface
      Properties:
         Tags:
         - Key: stack
           Value: production
         Description: A nice description.
         SourceDestCheck: 'false'
         GroupSet:
         - sg-75zzz219
         SubnetId: subnet-3z648z53
         PrivateIpAddress: 10.0.0.16
```

### Attach a network interface to an EC2 instance at launch

This example attaches a network interface to an EC2 instance. You can use the
NetworkInterface property to add more than one network interface. However, you can
specify multiple network interfaces if they all have only private IP addresses (no
associated public IP address). If you have a network interface with a public IP
address, specify when you launch the instance and then use
`AWS::EC2::NetworkInterfaceAttachment` to attach the additional network
interfaces.

#### JSON

```json

"Ec2Instance" : {
   "Type" : "AWS::EC2::Instance",
   "Properties" : {
      "ImageId" : { "Fn::FindInMap" : [ "RegionMap", { "Ref" : "AWS::Region" }, "AMI" ]},
      "KeyName" : { "Ref" : "KeyName" },
      "SecurityGroupIds" : [{ "Ref" : "WebSecurityGroup" }],
      "SubnetId" : { "Ref" : "SubnetId" },
      "NetworkInterfaces" : [ {
         "NetworkInterfaceId" : {"Ref" : "myENI"}, "DeviceIndex" : "1" } ],
      "Tags" : [ {"Key" : "Role", "Value" : "Test Instance"}],
      "UserData" : { "Fn::Base64" : { "Ref" : "WebServerPort" }}
   }
}
```

#### YAML

```yaml

Ec2Instance:
   Type: AWS::EC2::Instance
   Properties:
      ImageId:
         Fn::FindInMap:
         - RegionMap
         - Ref: AWS::Region
         - AMI
      KeyName:
         Ref: KeyName
      SecurityGroupIds:
      - Ref: WebSecurityGroup
      SubnetId:
         Ref: SubnetId
      NetworkInterfaces:
      - NetworkInterfaceId:
         Ref: myENI
        DeviceIndex: '1'
      Tags:
      - Key: Role
        Value: Test Instance
      UserData:
         Fn::Base64:
            Ref: WebServerPort
```

## See also

- [NetworkInterface](../../../../reference/awsec2/latest/apireference/api-networkinterface.md) in the _Amazon EC2 API_
_Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

ConnectionTrackingSpecification

All content copied from https://docs.aws.amazon.com/.
