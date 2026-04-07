This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElasticLoadBalancingV2::LoadBalancer SubnetMapping

Specifies a subnet for a load balancer.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AllocationId" : String,
  "IPv6Address" : String,
  "PrivateIPv4Address" : String,
  "SourceNatIpv6Prefix" : String,
  "SubnetId" : String
}

```

### YAML

```yaml

  AllocationId: String
  IPv6Address: String
  PrivateIPv4Address: String
  SourceNatIpv6Prefix: String
  SubnetId: String

```

## Properties

`AllocationId`

\[Network Load Balancers\] The allocation ID of the Elastic IP address for an
internet-facing load balancer.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IPv6Address`

\[Network Load Balancers\] The IPv6 address.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PrivateIPv4Address`

\[Network Load Balancers\] The private IPv4 address for an internal load balancer.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceNatIpv6Prefix`

\[Network Load Balancers with UDP listeners\] The IPv6 prefix to use for source NAT.
Specify an IPv6 prefix (/80 netmask) from the subnet CIDR block or `auto_assigned`
to use an IPv6 prefix selected at random from the subnet CIDR block.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubnetId`

The ID of the subnet.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

MinimumLoadBalancerCapacity

Tag
