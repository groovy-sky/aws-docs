This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::Subnet

Specifies a subnet for the specified VPC.

For an IPv4 only subnet, specify an IPv4 CIDR block. If the VPC has an IPv6 CIDR block,
you can create an IPv6 only subnet or a dual stack subnet instead. For an IPv6 only subnet,
specify an IPv6 CIDR block. For a dual stack subnet, specify both an IPv4 CIDR block and
an IPv6 CIDR block.

For more information, see [Subnets for your VPC](https://docs.aws.amazon.com/vpc/latest/userguide/configure-subnets.html) in the _Amazon VPC User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::Subnet",
  "Properties" : {
      "AssignIpv6AddressOnCreation" : Boolean,
      "AvailabilityZone" : String,
      "AvailabilityZoneId" : String,
      "CidrBlock" : String,
      "EnableDns64" : Boolean,
      "EnableLniAtDeviceIndex" : Integer,
      "Ipv4IpamPoolId" : String,
      "Ipv4NetmaskLength" : Integer,
      "Ipv6CidrBlock" : String,
      "Ipv6IpamPoolId" : String,
      "Ipv6Native" : Boolean,
      "Ipv6NetmaskLength" : Integer,
      "MapPublicIpOnLaunch" : Boolean,
      "OutpostArn" : String,
      "PrivateDnsNameOptionsOnLaunch" : PrivateDnsNameOptionsOnLaunch,
      "Tags" : [ Tag, ... ],
      "VpcId" : String
    }
}

```

### YAML

```yaml

Type: AWS::EC2::Subnet
Properties:
  AssignIpv6AddressOnCreation: Boolean
  AvailabilityZone: String
  AvailabilityZoneId: String
  CidrBlock: String
  EnableDns64: Boolean
  EnableLniAtDeviceIndex: Integer
  Ipv4IpamPoolId: String
  Ipv4NetmaskLength: Integer
  Ipv6CidrBlock: String
  Ipv6IpamPoolId: String
  Ipv6Native: Boolean
  Ipv6NetmaskLength: Integer
  MapPublicIpOnLaunch: Boolean
  OutpostArn: String
  PrivateDnsNameOptionsOnLaunch:
    PrivateDnsNameOptionsOnLaunch
  Tags:
    - Tag
  VpcId: String

```

## Properties

`AssignIpv6AddressOnCreation`

Indicates whether a network interface created in this subnet receives an IPv6 address.
The default value is `false`.

If you specify `AssignIpv6AddressOnCreation`, you must also specify
an IPv6 CIDR block.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AvailabilityZone`

The Availability Zone of the subnet.

If you update this property, you must also update the `CidrBlock`
property.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`AvailabilityZoneId`

The AZ ID of the subnet.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CidrBlock`

The IPv4 CIDR block assigned to the subnet.

If you update this property, we create a new subnet, and then delete the existing
one.

_Required_: Conditional

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EnableDns64`

Indicates whether DNS queries made to the Amazon-provided DNS Resolver in this subnet
should return synthetic IPv6 addresses for IPv4-only destinations.

###### Note

You must first configure a NAT gateway in a public subnet (separate from the subnet containing the IPv6-only workloads). For example, the subnet containing the NAT gateway should have a `0.0.0.0/0` route pointing to the internet gateway. For more information, see [Configure DNS64 and NAT64](https://docs.aws.amazon.com/vpc/latest/userguide/nat-gateway-nat64-dns64.html#nat-gateway-nat64-dns64-walkthrough) in the _Amazon Virtual Private Cloud User Guide_.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnableLniAtDeviceIndex`

Indicates the device position for local network interfaces in this subnet. For example,
`1` indicates local network interfaces in this subnet are the secondary
network interface (eth1).

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Ipv4IpamPoolId`

An IPv4 IPAM pool ID for the subnet.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Ipv4NetmaskLength`

An IPv4 netmask length for the subnet.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Ipv6CidrBlock`

The IPv6 CIDR block.

If you specify `AssignIpv6AddressOnCreation`, you must also specify
an IPv6 CIDR block.

_Required_: Conditional

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`Ipv6IpamPoolId`

An IPv6 IPAM pool ID for the subnet.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Ipv6Native`

Indicates whether this is an IPv6 only subnet. For more information, see [Subnet basics](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Subnets.html#subnet-basics) in the _Amazon Virtual Private Cloud User Guide_.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Ipv6NetmaskLength`

An IPv6 netmask length for the subnet.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MapPublicIpOnLaunch`

Indicates whether instances launched in this subnet receive a public IPv4 address. The
default value is `false`.

AWS charges for all public IPv4 addresses, including public IPv4 addresses
associated with running instances and Elastic IP addresses. For more information, see the _Public IPv4 Address_ tab
on the [VPC pricing page](https://aws.amazon.com/vpc/pricing).

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OutpostArn`

The Amazon Resource Name (ARN) of the Outpost.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PrivateDnsNameOptionsOnLaunch`

The hostname type for EC2 instances launched into this subnet and how DNS A and AAAA record queries to the instances should be handled. For more information, see [Amazon EC2 instance hostname types](../../../ec2/latest/userguide/ec2-instance-naming.md) in the _Amazon Elastic Compute Cloud User Guide_.

Available options:

- EnableResourceNameDnsAAAARecord (true \| false)

- EnableResourceNameDnsARecord (true \| false)

- HostnameType (ip-name \| resource-name)

_Required_: No

_Type_: [PrivateDnsNameOptionsOnLaunch](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ec2-subnet-privatednsnameoptionsonlaunch.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Any tags assigned to the subnet.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ec2-subnet-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcId`

The ID of the VPC the subnet is in.

If you update this property, you must also update the `CidrBlock`
property.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the subnet.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AvailabilityZone`

The Availability Zone of this subnet. For example, `us-east-1a`.

`AvailabilityZoneId`

The Availability Zone ID of this subnet. For example, `use1-az1`.

`CidrBlock`

The IPv4 CIDR blocks that are associated with the subnet.

`Ipv6CidrBlocks`

The IPv6 CIDR blocks that are associated with the subnet.

`NetworkAclAssociationId`

The ID of the network ACL that is associated with the subnet's VPC, such as
`acl-5fb85d36`.

`OutpostArn`

The Amazon Resource Name (ARN) of the Outpost.

`SubnetId`

The ID of the subnet.

`VpcId`

The ID of the subnet's VPC, such as `vpc-11ad4878`.

## Examples

- [Subnet with an IPv4 CIDR](#aws-resource-ec2-subnet--examples--Subnet_with_an_IPv4_CIDR)

- [Subnet with an IPv6 CIDR](#aws-resource-ec2-subnet--examples--Subnet_with_an_IPv6_CIDR)

### Subnet with an IPv4 CIDR

The following example creates a subnet with an IPv4 CIDR in a VPC with an IPv4
CIDR of 10.0.0.0/16. The VPC is declared elsewhere in the same template.

#### JSON

```json

"mySubnet" : {
   "Type" : "AWS::EC2::Subnet",
   "Properties" : {
      "VpcId" : { "Ref" : "myVPC" },
      "CidrBlock" : "10.0.0.0/24",
      "AvailabilityZone" : "us-east-1a",
      "Tags" : [ { "Key" : "stack", "Value" : "production" } ]
   }
}
```

#### YAML

```yaml

  mySubnet:
    Type: AWS::EC2::Subnet
    Properties:
      VpcId: !Ref myVPC
      CidrBlock: 10.0.0.0/24
      AvailabilityZone: "us-east-1a"
      Tags:
      - Key: stack
        Value: production
```

### Subnet with an IPv6 CIDR

The following example creates a subnet with an IPv6 CIDR in a VPC with an IPv6
CIDR provided by Amazon. The VPC is declared elsewhere in the same template. The
example uses the [Fn:Cidr](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-cidr.html) intrinsic function to select an IPv6 range with a /64 netmask for
the subnet.

#### JSON

```json

"mySubnet": {
   "Type": "AWS::EC2::Subnet",
   "Properties": {
      "VpcId": { "Ref": "myVPC" },
      "Ipv6Native": "true",
      "Ipv6CidrBlock": {
         "Fn::Select":
            [ 0, { "Fn::Cidr": [{"Fn::Select": [0, {"Fn::GetAtt": ["myVpc", "Ipv6CidrBlocks"]}]}, 1, 64 ]}]
      },
      "AssignIpv6AddressOnCreation": "true"
   }
}
```

#### YAML

```yaml

  mySubnet:
    Type: AWS::EC2::Subnet
    Properties:
      VpcId: !Ref myVPC
      Ipv6Native: true
      Ipv6CidrBlock: !Select [ 0, !Cidr [ !Select [ 0, !GetAtt myVpc.Ipv6CidrBlocks], 1, 64 ]]
      AssignIpv6AddressOnCreation: true
```

## See also

- [CreateSubnet](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateSubnet.html) in the _Amazon EC2 API_
_Reference_

- [VPC and\
subnets](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Subnets.html) in the _Amazon VPC User Guide_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::EC2::SqlHaStandbyDetectedInstance

BlockPublicAccessStates
