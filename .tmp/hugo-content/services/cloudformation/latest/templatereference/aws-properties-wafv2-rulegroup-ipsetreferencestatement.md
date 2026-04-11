This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFv2::RuleGroup IPSetReferenceStatement

A rule statement used to detect web requests coming from particular IP addresses or
address ranges. To use this, create an [AWS::WAFv2::IPSet](aws-resource-wafv2-ipset.md) that specifies the addresses
you want to detect, then use the ARN of that set in this statement.

Each IP set rule statement references an IP set. You create and maintain the set
independent of your rules. This allows you to use the single set in multiple rules. When
you update the referenced set, AWS WAF automatically updates all rules that
reference it.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Arn" : String,
  "IPSetForwardedIPConfig" : IPSetForwardedIPConfiguration
}

```

### YAML

```yaml

  Arn: String
  IPSetForwardedIPConfig:
    IPSetForwardedIPConfiguration

```

## Properties

`Arn`

The Amazon Resource Name (ARN) of the [AWS::WAFv2::IPSet](aws-resource-wafv2-ipset.md) that this statement
references.

_Required_: Yes

_Type_: String

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IPSetForwardedIPConfig`

The configuration for inspecting IP addresses in an HTTP header that you specify, instead of using the IP address that's reported by the web request origin. Commonly, this is the X-Forwarded-For (XFF) header, but you can specify any header name.

###### Note

If the specified header isn't present in the request, AWS WAF doesn't apply the rule to the web request at all.

_Required_: No

_Type_: [IPSetForwardedIPConfiguration](aws-properties-wafv2-rulegroup-ipsetforwardedipconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IPSetForwardedIPConfiguration

JA3Fingerprint

All content copied from https://docs.aws.amazon.com/.
