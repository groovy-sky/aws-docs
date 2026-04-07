This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Route53Resolver::ResolverEndpoint IpAddressRequest

In a
[CreateResolverEndpoint](../../../../reference/route53/latest/apireference/api-route53resolver-createresolverendpoint.md)
request, the IP address that DNS queries originate from (for outbound endpoints) or that you forward DNS queries to (for inbound endpoints).
`IpAddressRequest` also includes the ID of the subnet that contains the IP address.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Ip" : String,
  "Ipv6" : String,
  "SubnetId" : String
}

```

### YAML

```yaml

  Ip: String
  Ipv6: String
  SubnetId: String

```

## Properties

`Ip`

The IPv4 address that you want to use for DNS queries.

_Required_: No

_Type_: String

_Minimum_: `7`

_Maximum_: `36`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Ipv6`

The IPv6 address that you want to use for DNS queries.

_Required_: No

_Type_: String

_Minimum_: `7`

_Maximum_: `39`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubnetId`

The ID of the subnet that contains the IP address.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `32`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [Return values](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53resolver-resolverendpoint.html#aws-resource-route53resolver-resolverendpoint-return-values)
in the topic
[AWS::Route53Resolver::ResolverEndpoint](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53resolver-resolverendpoint.html)

- [IpAddressRequest](https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_IpAddressRequest.html)
in the _Amazon Route 53 API Reference_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Route53Resolver::ResolverEndpoint

Tag
