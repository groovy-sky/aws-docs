This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFv2::WebACL ForwardedIPConfiguration

The configuration for inspecting IP addresses in an HTTP header that you specify, instead of using the IP address that's reported by the web request origin. Commonly, this is the X-Forwarded-For (XFF) header, but you can specify any header name.

###### Note

If the specified header isn't present in the request, AWS WAF doesn't apply the rule to the web request at all.

This configuration is used for [GeoMatchStatement](../userguide/aws-properties-wafv2-webacl-statement.md#cfn-wafv2-webacl-statement-geomatchstatement), [AsnMatchStatement](../userguide/aws-properties-wafv2-webacl-statement.md#cfn-wafv2-webacl-statement-asnmatchstatement), and
[RateBasedStatement](../userguide/aws-properties-wafv2-webacl-statement.md#cfn-wafv2-webacl-statement-ratebasedstatement). For [IPSetReferenceStatement](../userguide/aws-properties-wafv2-webacl-statement.md#cfn-wafv2-webacl-statement-ipsetreferencestatement), use [IPSetForwardedIPConfig](../userguide/aws-properties-wafv2-rulegroup-ipsetreferencestatement.md#cfn-wafv2-rulegroup-ipsetreferencestatement-ipsetforwardedipconfig) instead.

AWS WAF only evaluates the first IP address found in the specified HTTP header.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FallbackBehavior" : String,
  "HeaderName" : String
}

```

### YAML

```yaml

  FallbackBehavior: String
  HeaderName: String

```

## Properties

`FallbackBehavior`

The match status to assign to the web request if the request doesn't have a valid IP address in the specified position.

###### Note

If the specified header isn't present in the request, AWS WAF doesn't apply the rule to the web request at all.

You can specify the following fallback behaviors:

- `MATCH` \- Treat the web request as matching the rule statement. AWS WAF applies the rule action to the request.

- `NO_MATCH` \- Treat the web request as not matching the rule statement.

_Required_: Yes

_Type_: String

_Allowed values_: `MATCH | NO_MATCH`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HeaderName`

The name of the HTTP header to use for the IP address. For example, to use the X-Forwarded-For (XFF) header, set this to `X-Forwarded-For`.

###### Note

If the specified header isn't present in the request, AWS WAF doesn't apply the rule to the web request at all.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9-]+{1,255}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FieldToProtect

GeoMatchStatement

All content copied from https://docs.aws.amazon.com/.
