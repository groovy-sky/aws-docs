This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Route53Resolver::ResolverEndpoint

Creates a Resolver endpoint. There are two types of Resolver endpoints, inbound and outbound:

- An _inbound Resolver endpoint_ forwards DNS queries to the DNS service for a VPC
from your network.

- An _outbound Resolver endpoint_ forwards DNS queries from the DNS service for a VPC
to your network.

###### Important

- You cannot update `ResolverEndpointType` and `IpAddresses` in the same request.

- When you update a dual-stack IP address, you must update both IP addresses. You can’t update only an IPv4 or IPv6 and keep an existing IP address.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Route53Resolver::ResolverEndpoint",
  "Properties" : {
      "Direction" : String,
      "IpAddresses" : [ IpAddressRequest, ... ],
      "Name" : String,
      "OutpostArn" : String,
      "PreferredInstanceType" : String,
      "Protocols" : [ String, ... ],
      "ResolverEndpointType" : String,
      "RniEnhancedMetricsEnabled" : Boolean,
      "SecurityGroupIds" : [ String, ... ],
      "Tags" : [ Tag, ... ],
      "TargetNameServerMetricsEnabled" : Boolean
    }
}

```

### YAML

```yaml

Type: AWS::Route53Resolver::ResolverEndpoint
Properties:
  Direction: String
  IpAddresses:
    - IpAddressRequest
  Name: String
  OutpostArn: String
  PreferredInstanceType: String
  Protocols:
    - String
  ResolverEndpointType: String
  RniEnhancedMetricsEnabled: Boolean
  SecurityGroupIds:
    - String
  Tags:
    - Tag
  TargetNameServerMetricsEnabled: Boolean

```

## Properties

`Direction`

Indicates whether the Resolver endpoint allows inbound or outbound DNS queries:

- `INBOUND`: allows DNS queries to your VPC from your network

- `OUTBOUND`: allows DNS queries from your VPC to your network

- `INBOUND_DELEGATION`: Resolver delegates queries to Route 53 private hosted zones from your network.

_Required_: Yes

_Type_: String

_Allowed values_: `INBOUND | OUTBOUND | INBOUND_DELEGATION`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`IpAddresses`

The subnets and IP addresses in your VPC that DNS queries originate from (for outbound endpoints) or that you forward
DNS queries to (for inbound endpoints). The subnet ID uniquely identifies a VPC.

###### Note

Even though the minimum is 1, Route 53 requires that you create at least two.

_Required_: Yes

_Type_: Array of [IpAddressRequest](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-route53resolver-resolverendpoint-ipaddressrequest.html)

_Minimum_: `2`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

A friendly name that lets you easily find a configuration in the Resolver dashboard in the Route 53 console.

_Required_: No

_Type_: String

_Pattern_: `(?!^[0-9]+$)([a-zA-Z0-9\-_' ']+)`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OutpostArn`

The ARN (Amazon Resource Name) for the Outpost.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws([a-z-]+)?:outposts:[a-z\d-]+:\d{12}:outpost/op-[a-f0-9]{17}$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PreferredInstanceType`

The Amazon EC2 instance type.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Protocols`

Protocols used for the endpoint. DoH-FIPS is applicable for a default inbound endpoints only.

For an inbound endpoint you can apply the protocols as follows:

- Do53 and DoH in combination.

- Do53 and DoH-FIPS in combination.

- Do53 alone.

- DoH alone.

- DoH-FIPS alone.

- None, which is treated as Do53.

For a delegation inbound endpoint you can use Do53 only.

For an outbound endpoint you can apply the protocols as follows:

- Do53 and DoH in combination.

- Do53 alone.

- DoH alone.

- None, which is treated as Do53.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `2`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResolverEndpointType`

The Resolver endpoint IP address type.

_Required_: No

_Type_: String

_Allowed values_: `IPV6 | IPV4 | DUALSTACK`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RniEnhancedMetricsEnabled`

Indicates whether RNI enhanced metrics are enabled for the Resolver endpoint.
When enabled, one-minute granular metrics are published in CloudWatch for each RNI associated with this endpoint.
When disabled, these metrics are not published.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecurityGroupIds`

The ID of one or more security groups that control access to this VPC. The security group must include one or more inbound rules
(for inbound endpoints) or outbound rules (for outbound endpoints). Inbound and outbound rules must allow TCP and UDP access.
For inbound access, open port 53. For outbound access, open the port that you're using for DNS queries on your network.

_Required_: Yes

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Route 53 Resolver doesn't support updating tags through CloudFormation.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-route53resolver-resolverendpoint-tag.html)

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetNameServerMetricsEnabled`

Indicates whether target name server metrics are enabled for the outbound Resolver endpoint.
When enabled, one-minute granular metrics are published in CloudWatch for each target name server associated with this endpoint.
When disabled, these metrics are not published. This feature is not supported for inbound Resolver endpoint.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the `ResolverEndpoint` object.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the resolver endpoint, such as
`arn:aws:route53resolver:us-east-1:123456789012:resolver-endpoint/resolver-endpoint-a1bzhi`.

`Direction`

Indicates whether the resolver endpoint allows inbound or outbound DNS queries.

`HostVPCId`

The ID of the VPC that you want to create the resolver endpoint in.

`IpAddressCount`

The number of IP addresses that the resolver endpoint can use for DNS queries.

`Name`

The name that you assigned to the resolver endpoint when you created the endpoint.

`ResolverEndpointId`

The ID of the resolver endpoint.

## Examples

- [Create Resolver endpoint](#aws-resource-route53resolver-resolverendpoint--examples--Create_Resolver_endpoint)

- [Associate a Resolver rule with a VPC](#aws-resource-route53resolver-resolverendpoint--examples--Associate_a_Resolver_rule_with_a_VPC)

### Create Resolver endpoint

The following example creates an Amazon Route 53 outbound resolver endpoint. The `IpAddresses` object
includes values for `SubnetId` but not for `Ip`. This causes Route 53 Resolver to automatically
choose an IP address from among the available IP addresses in the specified subnet.

#### JSON

```json

{
  "Type" : "AWS::Route53Resolver::ResolverEndpoint",
  "Properties" : {
    "Direction" : "Outbound",
    "IpAddresses" : [
      {
        "SubnetId": "subnet-0bca4d363dexample"
      },
      {
        "SubnetId": "subnet-0cdb5e474dexample"
      }
    ],
    "Name" : "MyOutboundEndpoint",
    "SecurityGroupIds" : [
      "sg-071b99f42example"
    ],
    "Tags" : [
      "Key": "LineOfBusiness",
      "Value": "Engineering"
    ]
  }
}
```

#### YAML

```yaml

Type : AWS::Route53Resolver::ResolverEndpoint
Properties :
  Direction : Outbound
  IpAddresses :
    - SubnetId: subnet-0bca4d363dexample
    - SubnetId: subnet-0cdb5e474dexample
  Name : MyOutboundEndpoint
  SecurityGroupIds :
    - sg-071b99f42example
  Tags :
    -
      Key: LineOfBusiness
      Value: Engineering
```

### Associate a Resolver rule with a VPC

The following example associates a resolver rule with a VPC. When you associate a rule with a VPC,
Resolver forwards all DNS queries for the domain name that is specified in the rule and that originate
in the VPC. The queries are forwarded to the IP addresses for the DNS resolvers that are specified in the rule.

#### JSON

```json

{
 "Type" : "AWS::Route53Resolver::ResolverRuleAssociation",
 "Properties" : {
	 "Name" : "MyResolverRuleAssociation",
	 "ResolverRuleId" : "rslvr-rr-5328a0899aexample",
	 "VPCId" : "vpc-03cf94c75cexample"
	 }
 }
```

#### YAML

```yaml

Type: "AWS::Route53Resolver::ResolverRuleAssociation"
	Properties:
	  Name: MyResolverRuleAssociation
	  ResolverRuleId: rslvr-rr-5328a0899aexample
	  VPCId: vpc-03cf94c75cexample
```

## See also

- [ResolverEndpoint](https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_ResolverEndpoint.html)
in the _Amazon Route 53 API Reference_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Route53Resolver::ResolverDNSSECConfig

IpAddressRequest
