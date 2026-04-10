This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFRegional::IPSet IPSetDescriptor

###### Note

AWS WAF Classic support will end on September 30, 2025.

This is **AWS WAF Classic** documentation. For
more information, see [AWS WAF Classic](../../../waf/latest/developerguide/classic-waf-chapter.md) in the developer guide.

**For the latest version of AWS WAF**, use the AWS WAFV2 API and see the [AWS WAF Developer Guide](../../../waf/latest/developerguide/waf-chapter.md). With the latest version, AWS WAF has a single set of endpoints for regional and global use.

Specifies the IP address type ( `IPV4` or `IPV6`) and the IP address range (in CIDR format) that web requests originate from.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : String,
  "Value" : String
}

```

### YAML

```yaml

  Type: String
  Value: String

```

## Properties

`Type`

Specify `IPV4` or `IPV6`.

_Required_: Yes

_Type_: String

_Allowed values_: `IPV4 | IPV6`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

Specify an IPv4 address by using CIDR notation. For example:

- To configure AWS WAF to allow, block, or count requests that originated from the IP address 192.0.2.44, specify `192.0.2.44/32`.

- To configure AWS WAF to allow, block, or count requests that originated from IP addresses from 192.0.2.0 to 192.0.2.255, specify
`192.0.2.0/24`.

For more information about CIDR notation, see the Wikipedia entry
[Classless Inter-Domain Routing](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing).

Specify an IPv6 address by using CIDR notation. For example:

- To configure AWS WAF to allow, block, or count requests that originated from the IP address 1111:0000:0000:0000:0000:0000:0000:0111, specify `1111:0000:0000:0000:0000:0000:0000:0111/128`.

- To configure AWS WAF to allow, block, or count requests that originated from IP addresses 1111:0000:0000:0000:0000:0000:0000:0000 to 1111:0000:0000:0000:ffff:ffff:ffff:ffff, specify `1111:0000:0000:0000:0000:0000:0000:0000/64`.

_Required_: Yes

_Type_: String

_Pattern_: `.*\S.*`

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::WAFRegional::IPSet

AWS::WAFRegional::RateBasedRule

All content copied from https://docs.aws.amazon.com/.
