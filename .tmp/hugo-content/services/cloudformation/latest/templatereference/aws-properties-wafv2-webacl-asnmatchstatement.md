This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFv2::WebACL AsnMatchStatement

A rule statement that inspects web traffic based on the Autonomous System Number (ASN) associated with the request's IP address.

For additional details, see [ASN match rule statement](../../../waf/latest/developerguide/waf-rule-statement-type-asn-match.md) in the [AWS WAF Developer Guide](../../../waf/latest/developerguide/waf-chapter.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AsnList" : [ Integer, ... ],
  "ForwardedIPConfig" : ForwardedIPConfiguration
}

```

### YAML

```yaml

  AsnList:
    - Integer
  ForwardedIPConfig:
    ForwardedIPConfiguration

```

## Properties

`AsnList`

Contains one or more Autonomous System Numbers (ASNs).
ASNs are unique identifiers assigned to large internet networks managed by organizations such as
internet service providers, enterprises, universities, or government agencies.

_Required_: No

_Type_: Array of Integer

_Minimum_: `0`

_Maximum_: `4294967295`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ForwardedIPConfig`

The configuration for inspecting IP addresses to match against an ASN in an HTTP header that you specify,
instead of using the IP address that's reported by the web request origin. Commonly, this is the X-Forwarded-For (XFF) header,
but you can specify any header name.

_Required_: No

_Type_: [ForwardedIPConfiguration](aws-properties-wafv2-webacl-forwardedipconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ApplicationConfig

AssociationConfig

All content copied from https://docs.aws.amazon.com/.
