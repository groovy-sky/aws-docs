This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::NatGateway

Specifies a network address translation (NAT) gateway in the specified subnet. You can
create either a public NAT gateway or a private NAT gateway. The default is a public NAT
gateway. If you create a public NAT gateway, you must specify an elastic IP address.

With a NAT gateway, instances in a private subnet can connect to the internet, other
AWS services, or an on-premises network using the IP address of the NAT
gateway. For more information, see [NAT gateways](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-nat-gateway.html) in the _Amazon VPC User Guide_.

If you add a default route ( `AWS::EC2::Route` resource) that points to a NAT
gateway, specify the NAT gateway ID for the route's `NatGatewayId`
property.

###### Important

When you associate an Elastic IP address or secondary Elastic IP address with a
public NAT gateway, the network border group of the Elastic IP address must match the network
border group of the Availability Zone (AZ) that the public NAT gateway is in. Otherwise, the
NAT gateway fails to launch. You can see the network border group for the AZ by viewing the
details of the subnet. Similarly, you can view the network border group for the Elastic IP
address by viewing its details. For more information, see [Allocate an Elastic IP address](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-eips.html#allocate-eip)
in the _Amazon VPC User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::NatGateway",
  "Properties" : {
      "AllocationId" : String,
      "AvailabilityMode" : String,
      "AvailabilityZoneAddresses" : [ AvailabilityZoneAddress, ... ],
      "ConnectivityType" : String,
      "MaxDrainDurationSeconds" : Integer,
      "PrivateIpAddress" : String,
      "SecondaryAllocationIds" : [ String, ... ],
      "SecondaryPrivateIpAddressCount" : Integer,
      "SecondaryPrivateIpAddresses" : [ String, ... ],
      "SubnetId" : String,
      "Tags" : [ Tag, ... ],
      "VpcId" : String
    }
}

```

### YAML

```yaml

Type: AWS::EC2::NatGateway
Properties:
  AllocationId: String
  AvailabilityMode: String
  AvailabilityZoneAddresses:
    - AvailabilityZoneAddress
  ConnectivityType: String
  MaxDrainDurationSeconds: Integer
  PrivateIpAddress: String
  SecondaryAllocationIds:
    - String
  SecondaryPrivateIpAddressCount: Integer
  SecondaryPrivateIpAddresses:
    - String
  SubnetId: String
  Tags:
    - Tag
  VpcId: String

```

## Properties

`AllocationId`

\[Public NAT gateway only\] The allocation ID of the Elastic IP address that's associated with the NAT gateway.
This property is required for a public NAT gateway and cannot be specified with a private NAT gateway.

_Required_: Conditional

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`AvailabilityMode`

Indicates whether this is a zonal (single-AZ) or regional (multi-AZ) NAT gateway.

A zonal NAT gateway is a NAT Gateway that provides redundancy and scalability within a single availability zone. A regional NAT gateway is a single NAT Gateway that works across multiple availability zones (AZs) in your VPC, providing redundancy, scalability and availability across all the AZs in a Region.

For more information, see [Regional NAT gateways for automatic multi-AZ expansion](https://docs.aws.amazon.com/vpc/latest/userguide/nat-gateways-regional.html) in the _Amazon VPC User Guide_.

_Required_: No

_Type_: String

_Allowed values_: `zonal | regional`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`AvailabilityZoneAddresses`

For regional NAT gateways only: Specifies which Availability Zones you want the NAT gateway to support and the Elastic IP addresses (EIPs) to use in each AZ. The regional NAT gateway uses these EIPs to handle outbound NAT traffic from their respective AZs. If not specified, the NAT gateway will automatically expand to new AZs and associate EIPs upon detection of an elastic network interface. If you specify this parameter, auto-expansion is disabled and you must manually manage AZ coverage.

A regional NAT gateway is a single NAT Gateway that works across multiple availability zones (AZs) in your VPC, providing redundancy, scalability and availability across all the AZs in a Region.

For more information, see [Regional NAT gateways for automatic multi-AZ expansion](https://docs.aws.amazon.com/vpc/latest/userguide/nat-gateways-regional.html) in the _Amazon VPC User Guide_.

_Required_: No

_Type_: Array of [AvailabilityZoneAddress](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ec2-natgateway-availabilityzoneaddress.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConnectivityType`

Indicates whether the NAT gateway supports public or private connectivity.
The default is public connectivity.

_Required_: No

_Type_: String

_Allowed values_: `private | public`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MaxDrainDurationSeconds`

The maximum amount of time to wait (in seconds) before forcibly releasing the IP addresses if connections are still in progress. Default value is 350 seconds.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `4000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PrivateIpAddress`

The private IPv4 address to assign to the NAT gateway. If you don't provide an address, a private IPv4 address will be automatically assigned.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SecondaryAllocationIds`

Secondary EIP allocation IDs. For more information, see [Create a NAT gateway](https://docs.aws.amazon.com/vpc/latest/userguide/nat-gateway-working-with.html)
in the _Amazon VPC User Guide_.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecondaryPrivateIpAddressCount`

\[Private NAT gateway only\] The number of secondary private IPv4 addresses you want to assign to the NAT gateway. For more information about secondary addresses, see [Create a NAT gateway](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-nat-gateway.html#nat-gateway-creating) in the _Amazon Virtual Private Cloud User Guide_.

`SecondaryPrivateIpAddressCount` and `SecondaryPrivateIpAddresses` cannot be set at the same time.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecondaryPrivateIpAddresses`

Secondary private IPv4 addresses. For more information about secondary addresses, see [Create a NAT gateway](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-nat-gateway.html#nat-gateway-creating) in the _Amazon Virtual Private Cloud User Guide_.

`SecondaryPrivateIpAddressCount` and `SecondaryPrivateIpAddresses` cannot be set at the same time.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubnetId`

The ID of the subnet in which the NAT gateway is located.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags for the NAT gateway.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ec2-natgateway-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcId`

The ID of the VPC in which the NAT gateway is located.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the NAT gateway. For example,
`nat-0a12bc456789de0fg`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AutoProvisionZones`

For regional NAT gateways only: Indicates whether AWS automatically manages AZ coverage. When enabled, the NAT gateway associates EIPs in all AZs where your VPC has subnets to handle outbound NAT traffic, expands to new AZs when you create subnets there, and retracts from AZs where you've removed all subnets. When disabled, you must manually manage which AZs the NAT gateway supports and their corresponding EIPs.

A regional NAT gateway is a single NAT Gateway that works across multiple availability zones (AZs) in your VPC, providing redundancy, scalability and availability across all the AZs in a Region.

For more information, see [Regional NAT gateways for automatic multi-AZ expansion](https://docs.aws.amazon.com/vpc/latest/userguide/nat-gateways-regional.html) in the _Amazon VPC User Guide_.

`AutoScalingIps`

For regional NAT gateways only: Indicates whether AWS automatically allocates additional Elastic IP addresses (EIPs) in an AZ when the NAT gateway needs more ports due to increased concurrent connections to a single destination from that AZ.

For more information, see [Regional NAT gateways for automatic multi-AZ expansion](https://docs.aws.amazon.com/vpc/latest/userguide/nat-gateways-regional.html) in the _Amazon VPC User Guide_.

`EniId`

The ID of the network interface.

`NatGatewayId`

The ID of the NAT gateway.

`RouteTableId`

For regional NAT gateways only, this is the ID of the NAT gateway.

## Examples

### NAT gateway

The following example creates a public NAT gateway and a route that sends all
internet-bound traffic from the private subnet with EC2 instances to the NAT gateway.
A public NAT gateway uses an elastic IP address to provide it with a public IP address
that doesn't change. Note that the route table for the public subnet with the NAT gateway
must also have a route that sends all internet-bound traffic to an internet gateway,
so that the NAT gateway can connect to the internet.

#### JSON

```json

"NATGateway" : {
   "Type" : "AWS::EC2::NatGateway",
   "Properties" : {
      "AllocationId" : {
          "Fn::GetAtt" : ["NATGatewayEIP", "AllocationId"]
      },
      "SubnetId" : {
          "Ref" : "PublicSubnet"
      },
      "Tags" : [
          {"Key" : "stack", "Value" : "production" }
      ]
     }
},
"NATGatewayEIP" : {
   "Type" : "AWS::EC2::EIP",
   "Properties" : {
      "Domain" : "vpc"
   }
},
"RouteNATGateway" : {
   "Type" : "AWS::EC2::Route",
   "Properties" : {
      "RouteTableId" : { "Ref" : "PrivateRouteTable" },
      "DestinationCidrBlock" : "0.0.0.0/0",
      "NatGatewayId" : { "Ref" : "NATGateway" }
   }
}
```

#### YAML

```yaml

NATGateway:
   Type: AWS::EC2::NatGateway
   Properties:
      AllocationId: !GetAtt NATGatewayEIP.AllocationId
      SubnetId: !Ref PublicSubnet
      Tags:
      - Key: stack
        Value: production
NATGatewayEIP:
   Type: AWS::EC2::EIP
   Properties:
      Domain: vpc
RouteNATGateway:
   Type: AWS::EC2::Route
   Properties:
      RouteTableId: !Ref PrivateRouteTable
      DestinationCidrBlock: '0.0.0.0/0'
      NatGatewayId: !Ref NATGateway
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AvailabilityZoneAddress
