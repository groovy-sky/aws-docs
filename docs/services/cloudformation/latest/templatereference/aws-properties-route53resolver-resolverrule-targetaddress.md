This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Route53Resolver::ResolverRule TargetAddress

In a
[CreateResolverRule](../../../../reference/route53/latest/apireference/api-route53resolver-createresolverrule.md)
request, an array of the IPs that you want to forward DNS queries to.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Ip" : String,
  "Ipv6" : String,
  "Port" : String,
  "Protocol" : String,
  "ServerNameIndication" : String
}

```

### YAML

```yaml

  Ip: String
  Ipv6: String
  Port: String
  Protocol: String
  ServerNameIndication: String

```

## Properties

`Ip`

One IPv4 address that you want to forward DNS queries to.

_Required_: No

_Type_: String

_Minimum_: `7`

_Maximum_: `36`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Ipv6`

One IPv6 address that you want to forward DNS queries to.

_Required_: No

_Type_: String

_Minimum_: `7`

_Maximum_: `39`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Port`

The port at `Ip` that you want to forward DNS queries to.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `65535`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Protocol`

The protocols for the target address. The protocol you choose needs to be supported by the outbound endpoint of the Resolver rule.

_Required_: No

_Type_: String

_Allowed values_: `Do53 | DoH`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServerNameIndication`

The Server Name Indication of the DoH server that you want to forward queries to.
This is only used if the Protocol of the `TargetAddress` is `DoH`.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [Return values](../userguide/aws-resource-route53resolver-resolverrule.md#aws-resource-route53resolver-resolverrule-return-values)
in the topic
[AWS::Route53Resolver::ResolverRule](../userguide/aws-resource-route53resolver-resolverrule.md)

- [TargetAddress](../../../../reference/route53/latest/apireference/api-route53resolver-targetaddress.md)
in the _Amazon Route 53 API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWS::Route53Resolver::ResolverRuleAssociation

All content copied from https://docs.aws.amazon.com/.
