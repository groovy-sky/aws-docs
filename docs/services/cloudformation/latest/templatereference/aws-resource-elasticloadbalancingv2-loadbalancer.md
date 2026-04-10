This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElasticLoadBalancingV2::LoadBalancer

Specifies an Application Load Balancer, a Network Load Balancer, or a Gateway Load
Balancer.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ElasticLoadBalancingV2::LoadBalancer",
  "Properties" : {
      "EnableCapacityReservationProvisionStabilize" : Boolean,
      "EnablePrefixForIpv6SourceNat" : String,
      "EnforceSecurityGroupInboundRulesOnPrivateLinkTraffic" : String,
      "IpAddressType" : String,
      "Ipv4IpamPoolId" : String,
      "LoadBalancerAttributes" : [ LoadBalancerAttribute, ... ],
      "MinimumLoadBalancerCapacity" : MinimumLoadBalancerCapacity,
      "Name" : String,
      "Scheme" : String,
      "SecurityGroups" : [ String, ... ],
      "SubnetMappings" : [ SubnetMapping, ... ],
      "Subnets" : [ String, ... ],
      "Tags" : [ Tag, ... ],
      "Type" : String
    }
}

```

### YAML

```yaml

Type: AWS::ElasticLoadBalancingV2::LoadBalancer
Properties:
  EnableCapacityReservationProvisionStabilize: Boolean
  EnablePrefixForIpv6SourceNat: String
  EnforceSecurityGroupInboundRulesOnPrivateLinkTraffic: String
  IpAddressType: String
  Ipv4IpamPoolId: String
  LoadBalancerAttributes:
    - LoadBalancerAttribute
  MinimumLoadBalancerCapacity:
    MinimumLoadBalancerCapacity
  Name: String
  Scheme: String
  SecurityGroups:
    - String
  SubnetMappings:
    - SubnetMapping
  Subnets:
    - String
  Tags:
    - Tag
  Type: String

```

## Properties

`EnableCapacityReservationProvisionStabilize`

Indicates whether to enable stabilization when creating or updating an LCU reservation.
This ensures that the final stack status reflects the status of the LCU reservation.
The default is `false`.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnablePrefixForIpv6SourceNat`

\[Network Load Balancers with UDP listeners\] Indicates whether to use an IPv6 prefix
from each subnet for source NAT. The IP address type must be `dualstack`.
The default value is `off`.

_Required_: No

_Type_: String

_Allowed values_: `on | off`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnforceSecurityGroupInboundRulesOnPrivateLinkTraffic`

Indicates whether to evaluate inbound security group rules for traffic sent to a
Network Load Balancer through AWS PrivateLink. The default is `on`.

You can't configure this property on a Network Load Balancer unless you associated a
security group with the load balancer when you created it.

_Required_: No

_Type_: String

_Allowed values_: `on | off`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IpAddressType`

The IP address type. Internal load balancers must use `ipv4`.

\[Application Load Balancers\] The possible values are `ipv4` (IPv4 addresses),
`dualstack` (IPv4 and IPv6 addresses), and `dualstack-without-public-ipv4`
(public IPv6 addresses and private IPv4 and IPv6 addresses).

Application Load Balancer authentication supports IPv4 addresses only when
connecting to an Identity Provider (IdP) or Amazon Cognito endpoint. Without a public
IPv4 address the load balancer can't complete the authentication process, resulting
in HTTP 500 errors.

\[Network Load Balancers and Gateway Load Balancers\] The possible values are `ipv4`
(IPv4 addresses) and `dualstack` (IPv4 and IPv6 addresses).

_Required_: No

_Type_: String

_Allowed values_: `ipv4 | dualstack | dualstack-without-public-ipv4`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Ipv4IpamPoolId`

The ID of the IPv4 IPAM pool.

_Required_: No

_Type_: String

_Pattern_: `^(ipam-pool-)[a-zA-Z0-9]+$`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LoadBalancerAttributes`

The load balancer attributes. Attributes that you do not modify retain their current values.

_Required_: No

_Type_: Array of [LoadBalancerAttribute](aws-properties-elasticloadbalancingv2-loadbalancer-loadbalancerattribute.md)

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinimumLoadBalancerCapacity`

The minimum capacity for a load balancer.

_Required_: No

_Type_: [MinimumLoadBalancerCapacity](aws-properties-elasticloadbalancingv2-loadbalancer-minimumloadbalancercapacity.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the load balancer. This name must be unique per region per account, can have
a maximum of 32 characters, must contain only alphanumeric characters or hyphens, must not
begin or end with a hyphen, and must not begin with "internal-".

If you don't specify a name, AWS CloudFormation generates a unique
physical ID for the load balancer. If you specify a name, you cannot perform updates that
require replacement of this resource, but you can perform other updates. To replace the
resource, specify a new name.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Scheme`

The nodes of an Internet-facing load balancer have public IP addresses. The DNS name of an
Internet-facing load balancer is publicly resolvable to the public IP addresses of the nodes.
Therefore, Internet-facing load balancers can route requests from clients over the
internet.

The nodes of an internal load balancer have only private IP addresses. The DNS name of an
internal load balancer is publicly resolvable to the private IP addresses of the nodes.
Therefore, internal load balancers can route requests only from clients with access to the VPC
for the load balancer.

The default is an Internet-facing load balancer.

You can't specify a scheme for a Gateway Load Balancer.

_Required_: No

_Type_: String

_Allowed values_: `internet-facing | internal`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SecurityGroups`

\[Application Load Balancers and Network Load Balancers\] The IDs of the security groups for
the load balancer.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubnetMappings`

The IDs of the subnets. You can specify only one subnet per Availability Zone. You
must specify either subnets or subnet mappings, but not both.

\[Application Load Balancers\] You must specify subnets from at least two Availability
Zones. You can't specify Elastic IP addresses for your subnets.

\[Application Load Balancers on Outposts\] You must specify one Outpost subnet.

\[Application Load Balancers on Local Zones\] You can specify subnets from one or more Local
Zones.

\[Network Load Balancers\] You can specify subnets from one or more Availability Zones. You
can specify one Elastic IP address per subnet if you need static IP addresses for your
internet-facing load balancer. For internal load balancers, you can specify one private IP
address per subnet from the IPv4 range of the subnet. For internet-facing load balancer, you
can specify one IPv6 address per subnet.

\[Gateway Load Balancers\] You can specify subnets from one or more Availability Zones. You
can't specify Elastic IP addresses for your subnets.

_Required_: Conditional

_Type_: Array of [SubnetMapping](aws-properties-elasticloadbalancingv2-loadbalancer-subnetmapping.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Subnets`

The IDs of the subnets. You can specify only one subnet per Availability Zone. You
must specify either subnets or subnet mappings, but not both. To specify an Elastic IP
address, specify subnet mappings instead of subnets.

\[Application Load Balancers\] You must specify subnets from at least two Availability
Zones.

\[Application Load Balancers on Outposts\] You must specify one Outpost subnet.

\[Application Load Balancers on Local Zones\] You can specify subnets from one or more Local
Zones.

\[Network Load Balancers and Gateway Load Balancers\] You can specify subnets from one or more
Availability Zones.

_Required_: Conditional

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags to assign to the load balancer.

_Required_: No

_Type_: Array of [Tag](aws-properties-elasticloadbalancingv2-loadbalancer-tag.md)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of load balancer. The default is `application`.

_Required_: No

_Type_: String

_Allowed values_: `application | network | gateway`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the load balancer.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CanonicalHostedZoneID`

The ID of the Amazon Route 53 hosted zone associated with the load balancer. For
example, `Z2P70J7EXAMPLE`.

`DNSName`

The DNS name for the load balancer. For example,
`my-load-balancer-424835706.us-west-2.elb.amazonaws.com`.

`LoadBalancerArn`

The Amazon Resource Name (ARN) of the load balancer.

`LoadBalancerFullName`

The full name of the load balancer. For example,
`app/my-load-balancer/50dc6c495c0c9188`.

`LoadBalancerName`

The name of the load balancer. For example, `my-load-balancer`.

`SecurityGroups`

The IDs of the security groups for the load balancer.

## Examples

To get started with Elastic Load Balancer create a load balancer.
After you create your load balancer, add a listener using [AWS::ElasticLoadBalancingV2::Listener](../userguide/aws-resource-elasticloadbalancingv2-listener.md).

- [Create an Application Load Balancer](#aws-resource-elasticloadbalancingv2-loadbalancer--examples--Create_an_Application_Load_Balancer)

- [Create a Network Load Balancer](#aws-resource-elasticloadbalancingv2-loadbalancer--examples--Create_a_Network_Load_Balancer)

### Create an Application Load Balancer

The following example creates an internal Application Load Balancer with an associated
security group and a load balancer attribute.

#### YAML

```yaml

myLoadBalancer:
    Type: 'AWS::ElasticLoadBalancingV2::LoadBalancer'
    Properties:
      Name: my-alb
      Type: application
      Scheme: internal
      Subnets:
        - !Ref subnet-AZ1
        - !Ref subnet-AZ2
      SecurityGroups:
        - !Ref mySecurityGroup
      LoadBalancerAttributes:
        - Key: "deletion_protection.enabled"
          Value: "true"
```

#### JSON

```json

{
    "myLoadBalancer": {
        "Type": "AWS::ElasticLoadBalancingV2::LoadBalancer",
        "Properties": {
            "Name": "my-alb",
            "Type": "application",
            "Scheme": "internal",
            "Subnets": [
                {
                    "Ref": "subnet-AZ1"
                },
                {
                    "Ref": "subnet-AZ2"
                }
            ],
            "SecurityGroups": [
                {
                    "Ref": "mySecurityGroup"
                }
            ],
            "LoadBalancerAttributes": [
                {
                    "Key": "deletion_protection.enabled",
                    "Value": true
                }
            ]
        }
    }
}
```

### Create a Network Load Balancer

The following example creates an internal Network Load Balancer with an associated
security group and a load balancer attribute.

#### YAML

```yaml

myLoadBalancer:
    Type: 'AWS::ElasticLoadBalancingV2::LoadBalancer'
    Properties:
      Name: my-nlb
      Type: network
      Scheme: internal
      Subnets:
        - !Ref subnet-AZ1
        - !Ref subnet-AZ2
      SecurityGroups:
        - !Ref mySecurityGroup
      LoadBalancerAttributes:
        - Key: "deletion_protection.enabled"
          Value: "true"
```

#### JSON

```json

{
    "myLoadBalancer": {
        "Type": "AWS::ElasticLoadBalancingV2::LoadBalancer",
        "Properties": {
            "Name": "my-alb",
            "Type": "network",
            "Scheme": "internal",
            "Subnets": [
                {
                    "Ref": "subnet-AZ1"
                },
                {
                    "Ref": "subnet-AZ2"
                }
            ],
            "SecurityGroups": [
                {
                    "Ref": "mySecurityGroup"
                }
            ],
            "LoadBalancerAttributes": [
                {
                    "Key": "deletion_protection.enabled",
                    "Value": true
                }
            ]
        }
    }
}
```

## See also

- [CreateLoadBalancer](../../../../reference/elasticloadbalancing/latest/apireference/api-createloadbalancer.md) in the _Elastic Load Balancing API_
_Reference (version 2015-12-01)_

- [User\
Guide for Application Load Balancers](../../../elasticloadbalancing/latest/application.md)

- [User Guide\
for Network Load Balancers](../../../elasticloadbalancing/latest/network.md)

- [User Guide\
for Gateway Load Balancers](../../../elasticloadbalancing/latest/gateway.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Transform

LoadBalancerAttribute

All content copied from https://docs.aws.amazon.com/.
